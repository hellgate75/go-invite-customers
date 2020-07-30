/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package geo

import "math"

//  This routine calculates the distance between two points (given the
//  latitude/longitude of those points). It is being used to calculate
//  the distance between two locations using GeoDataSource (TM) products
//
//  Definitions/
//    South latitudes are negative, east longitudes are positive
//
//  Passed to function/
//    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)
//    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)
//    unit = the unit you desire for results
//           where/ 'M' is statute miles (default)
//                  'K' is kilometers
//                  'N' is nautical miles
//
func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
