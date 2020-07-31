/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package main

import (
	"flag"
	"fmt"
	"github.com/hellgate75/go-invite-customers/invite"
	"github.com/hellgate75/go-invite-customers/io"
	"os"
	"strings"
)

var flagSet *flag.FlagSet

var fileOrStream string
var homeLatitude float64 = 53.339428
var homeLongitude float64 = -6.257664
var distance float64 = 100.0
var measureUnit string = "K"
var inputEncoding string = "json"
var usePerLineInput bool = true
var silentOutput bool = false
var outputEncoding string = "text"
var useDetailedOutput bool = false

func printUsage(message string, exitCode int) {
	fmt.Println("go-invite-customers -[param0]=value0 ...  -[paramN]=valueN")
	if len(message) == 0 {
		fmt.Printf("Error: %s", message)
	}
	fmt.Println("Parameters:")
	flagSet.PrintDefaults()
	if exitCode >= 0 {
		os.Exit(exitCode)
	}
}

func init() {
	flagSet = flag.NewFlagSet("go-invite-customers", flag.ContinueOnError)
	flagSet.StringVar(&fileOrStream, "input", "", "Given file, url or pipe that contains data")
	flagSet.Float64Var(&homeLatitude, "latitude", homeLatitude, "Base latitude degrees in float number [W is negative]")
	flagSet.Float64Var(&homeLongitude, "longitude", homeLongitude, "Base longitude degrees in float number [S is negative]")
	flagSet.Float64Var(&distance, "distance", distance, "Max distance from base coordinate")
	flagSet.StringVar(&measureUnit, "unit", "K", "Measure Unit for distance [K is for Kilometers, M is for Miles and N is for Nautical Miles]")
	flagSet.StringVar(&inputEncoding, "in-enc", "json", fmt.Sprintf("Input encoding format: %v", io.InputEncoding))
	flagSet.StringVar(&outputEncoding, "out-enc", "text", fmt.Sprintf("Output encoding format: %v", io.OutputEncoding))
	flagSet.BoolVar(&usePerLineInput, "per-line-input", true, "Use one read line in input for parsing the data, instead of reading the list")
	flagSet.BoolVar(&silentOutput, "silent", false, "Execute silent output")
	flagSet.BoolVar(&useDetailedOutput, "detailed", false, "Create Output for invited and excluded, instead of only invited customers")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		printUsage(err.Error(), 1)
	}
}

func main() {
	if distance <= 0 || measureUnit == "" {
		printUsage("Distance cannot be zero or less and unit cannot be empty", 2)
	}
	measureUnit = strings.ToUpper(measureUnit)
	if measureUnit != "K" && measureUnit != "M" && measureUnit != "N" {
		printUsage("Distance Measure Unit can have only on of 'K', 'M' or 'N' values", 2)
	}
	fileOrStream = strings.TrimSpace(fileOrStream)
	if "" == fileOrStream {
		printUsage("File, stream or pipe reference cannot be empty", 2)
	}
	var inEnc, outEnc io.Encoding
	var err error
	if inEnc, err = io.ToEncoding(inputEncoding); err != nil {
		printUsage(fmt.Sprintf("Error converting input encoding from string: %s", inputEncoding), 2)

	}
	if outEnc, err = io.ToEncoding(outputEncoding); err != nil {
		printUsage(fmt.Sprintf("Error converting output encoding from string: %s", outputEncoding), 2)

	}
	if !silentOutput {
		fmt.Println("Calculating customers within given distance from the base coordinates....")
	}
	out, errs := invite.ExecuteInviteScan(invite.InputData{
		fileOrStream,
		homeLatitude,
		homeLongitude,
		distance,
		measureUnit,
		inEnc,
		usePerLineInput,
		useDetailedOutput,
		silentOutput,
		outEnc,
	})
	if len(errs) > 0 {
		if silentOutput {
			fmt.Println("Processing errors occurred in evaluation, please run without silent option for details")
		} else {
			fmt.Println("Processing errors:")
			for _, err := range errs {
				fmt.Println(err.Error())
			}
		}
	}
	var data []byte
	if out.IsComplete {
		data, err = io.EncodeCustomerDetailedInvite(*out.Complete, outEnc)
	} else {
		data, err = io.EncodeCustomerInvite(*out.Simple, outEnc)
	}
	if err != nil {
		if silentOutput {
			fmt.Println("Error converting output, please run without silent option for details")
		} else {
			fmt.Printf("Error converting output: %v\n", err)
		}

	} else {
		fmt.Println(string(data))
	}
}
