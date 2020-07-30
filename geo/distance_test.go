/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package geo

import "testing"

func Test_Distance(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
		unit []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Calculate coordinates distance in Miles",
			args{
				32.9697,
				-96.80322,
				29.46786,
				-98.53506,
				[]string{"M"},
			},
			262.677793805435,
		},
		{
			"Calculate coordinates distance in Kilometers",
			args{
				32.9697,
				-96.80322,
				29.46786,
				-98.53506,
				[]string{"K"},
			},
			422.738931394014,
		},
		{
			"Calculate coordinates distance in Nautical Miles",
			args{
				32.9697,
				-96.80322,
				29.46786,
				-98.53506,
				[]string{"N"},
			},
			228.10939614063972,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2, tt.args.unit...); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
