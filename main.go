/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package main

import (
	"fmt"
	"github.com/hellgate75/go-invite-customers/geo"
)

func init() {

}

func main() {
	fmt.Printf("%f Miles\n", geo.Distance(32.9697, -96.80322, 29.46786, -98.53506, "M"))
	fmt.Printf("%f Kilometers\n", geo.Distance(32.9697, -96.80322, 29.46786, -98.53506, "K"))
	fmt.Printf("%f Nautical Miles\n", geo.Distance(32.9697, -96.80322, 29.46786, -98.53506, "N"))

}
