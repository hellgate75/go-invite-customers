/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package invite

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/hellgate75/go-invite-customers/geo"
	"github.com/hellgate75/go-invite-customers/io"
	"github.com/hellgate75/go-invite-customers/model"
	io2 "io"
	"io/ioutil"
	"time"
)

type OutputData struct {
	Simple     *model.InviteList
	Complete   *model.CompleteInviteList
	IsComplete bool
	IsDone     bool
}

type InputData struct {
	FileOrStream      string
	HomeLatitude      float64
	HomeLongitude     float64
	Distance          float64
	MeasureUnit       string
	InputEncoding     io.Encoding
	UsePerLineInput   bool
	UseDetailedOutput bool
	SilentOutput      bool
	OutputEncoding    io.Encoding
}

func readLineByLine(r io2.Reader, inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
	br := bufio.NewReader(r)
	buff := bytes.NewBuffer([]byte{})
	line, isPref, err := br.ReadLine()
	for err == nil {
		if isPref {
			buff.Write(line)
		} else {
			if buff.Len() > 0 {
				buff.Write(line)
				line = buff.Bytes()
				buff.Reset()
			}
			customer, errP := io.ReadCustomerOffice(line, inputData.InputEncoding)
			if errP != nil {
				errCh <- errP
			} else {
				ch <- customer
			}
		}
		line, isPref, err = br.ReadLine()
	}
}
func parseAndServerList(r io2.Reader, inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		errCh <- err
		return
	}
	list, err := io.ReadCustomerOfficeList(data, inputData.InputEncoding)
	if err != nil {
		errCh <- err
		return
	}
	for _, customer := range list.List {
		ch <- customer
	}
}

func createChannelWriterFunc(url string) (function func(InputData, chan model.CustomerOffice, chan error), err error) {
	if url[:6] == "udp://" {
		// Udp protocol
		c, r, err := OpenUdpStream(url)
		if err != nil {
			return function, err
		}
		function = func(inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
			defer func() {
				// Close the udp connection
				if c != nil {
					_ = c.Close()
				}
			}()
			if inputData.UsePerLineInput {
				readLineByLine(r, inputData, ch, errCh)
			} else {
				parseAndServerList(r, inputData, ch, errCh)
			}
		}
	} else if url[:6] == "tcp://" {
		// Tcp protocol
		c, r, err := OpenTcpStream(url)
		if err != nil {
			return function, err
		}
		function = func(inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
			defer func() {
				// Close the udp connection
				if c != nil {
					_ = c.Close()
				}
			}()
			if inputData.UsePerLineInput {
				readLineByLine(r, inputData, ch, errCh)
			} else {
				parseAndServerList(r, inputData, ch, errCh)
			}
		}
	} else if url[:7] == "http://" || url[:8] == "https://" || url[:6] == "ftp://" || url[:7] == "sftp://" {
		// Http / Ftp protocol
		re, r, err := OpenUrlStream(url)
		if err != nil {
			return function, err
		}
		function = func(inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
			defer func() {
				// Close the udp connection
				if re != nil {
					_ = re.Body.Close()
				}
			}()
			if inputData.UsePerLineInput {
				readLineByLine(r, inputData, ch, errCh)
			} else {
				parseAndServerList(r, inputData, ch, errCh)
			}
		}
	} else {
		// file protocol
		f, err := OpenFileStream(url)
		if err != nil {
			return function, err
		}
		function = func(inputData InputData, ch chan model.CustomerOffice, errCh chan error) {
			defer func() {
				// Close the udp connection
				if f != nil {
					_ = f.Close()
				}
			}()
			if inputData.UsePerLineInput {
				readLineByLine(f, inputData, ch, errCh)
			} else {
				parseAndServerList(f, inputData, ch, errCh)
			}
		}
	}
	return function, err
}

func ExecuteInviteScan(input InputData) (out OutputData, errs []error) {
	out = OutputData{
		Simple:     model.NewInviteList(),
		Complete:   model.NewCompleteInviteList(),
		IsComplete: input.UseDetailedOutput,
	}
	errs = make([]error, 0)
	var ch = make(chan model.CustomerOffice)
	var errCh = make(chan error)
	fn, err := createChannelWriterFunc(input.FileOrStream)
	if err != nil {
		errs = append(errs, err)
		return out, errs
	}
	go fn(input, ch, errCh)
	var running = true
	go func(errChannel chan error) {
		// Collecting errors
		for running {
			select {
			case errX := <-errChannel:
				errs = append(errs, errX)
			case <-time.After(2 * time.Second):
				continue
			}
		}
	}(errCh)
	// Collecting customers
computeCycle:
	for true {
		select {
		case customer := <-ch:
			go func(inputData InputData, customerOffice model.CustomerOffice, out *OutputData) {
				// Verifies if customer has correct coordinates
				if !customerOffice.IsValid() {
					errs = append(errs, errors.New(fmt.Sprintf("Invalid coordinates data for customer [%v] %s", customer.UserId, customer.Name)))
				}
				// Recovers customer office latitude and longitude
				lat, _ := customerOffice.GetLatitude()
				long, _ := customerOffice.GetLongitude()
				// Calculates distance
				dist := geo.Distance(inputData.HomeLatitude, inputData.HomeLongitude, lat, long, inputData.MeasureUnit)
				if out.IsComplete {
					// If is detailed output collects invited and excluded  customers
					if dist <= inputData.Distance {
						out.Complete.AddInvited(model.ToInviteData(&customerOffice))
					} else {
						out.Complete.AddExcluded(model.ToInviteData(&customerOffice))
					}
				} else {
					// If is simple output collects only invited customers
					if dist <= inputData.Distance {
						out.Simple.Add(model.ToInviteData(&customerOffice))
					}
				}
			}(input, customer, &out)
		case <-time.After(10 * time.Second):
			running = false
			break computeCycle
		}
	}
	out.IsDone = len(errs) == 0
	return out, errs
}
