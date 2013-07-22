package go_geohash

import "strings"

// Default precision to use
const default_precision = 9

// Static array of 0-9, a-z
const base32_codes = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "b", "c", "d", "e", "f", "g", "h", "j", "k", "m", "n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

const base32_map = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "b": 10, "c": 11, "d": 12, "e": 13, "f": 14, "g": 15, "h": 16, "j": 17, "k": 18, "m": 19, "n": 20, "p": 21, "q": 22, "r": 23, "s": 24, "t": 25, "u": 26, "v": 27, "w": 28, "x": 29, "y": 30, "z": 31}

//
// Bounded box for parsing latitude and longitude
//
type DecodedBoundBox struct {
	MinLatitude  float64
	MinLongitude float64

	MaxLatitude  float64
	MaxLongitude float64
}

//
// Decoded Latitude,Longitude + error position
//
type DecodedPosition struct {
	Latitude  float64
	Longitude float64

	LatitudeError  float64
	LongitudeError float64
}

//
// Encode a Latitude and Longitude as a string
//
// Arguments:
//  latitude  float64
//  longitude float64
//  precision uint8  (ie how long should the hash string be?)
//
func Encode(latitude float64, longitude float64, precision uint8) string {
	// Pre-Allocate the hash string
	var output string = strings.Repeat(' ', precision)

	// DecodedBoundBox for the lat/lon + errors
	var bbox DecodedBoundBox = DecodedBoundBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}

	var min_max_avg float64 = 0
	var islon bool = true
	var num_bits uint = 0
	var hash_index int = 0

	var output_length uint8 = 0
	for output_length < precision {
		if islon {
			min_max_avg = (bbox.MaxLongitude + bbox.MinLongitude) / 2
			if longitude > min_max_avg {
				hash_index = (hash_index << 1) + 1
				bbox.MinLongitude = min_max_avg
			} else {
				hash_index = (hash_index << 1) + 0
				bbox.MaxLongitude = min_max_avg
			}
		} else {
			min_max_avg = (bbox.MaxLatitude + bbox.MinLatitude) / 2
			if latitude > min_max_avg {
				hash_index = (hash_index << 1) + 1
				bbox.MinLatitude = min_max_avg
			} else {
				hash_index = (hash_index << 1) + 0
				bbox.MaxLatitude = min_max_avg
			}
		}
		islon = !islon

		num_bits++
		if 5 == num_bits {
			output[output_length] = base32_codes[hash_index]

			output_length++
			num_bits = 0
			hash_index = 0
		}
	}

	return output
}

func DecodeBoundBox(hash_string string) DecodedBoundBox {
	// Downcase the input string
	hash_string = strings.ToLower(hash_string)

	var output DecodedBoundBox = DecodedBoundBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}

	var islon bool = true
	for i := 0; i < len(hash_string); i++ {
		var char_index = base32_map[hash_string[i]]

		for bits := 4; bits >= 0; bits-- {
			bit := (char_index >> bits) & 1
			if islon {
				mid := (output.MaxLongitude + output.MinLongitude) / 2
				if bit == 1 {
					output.MinLongitude = mid
				} else {
					output.MaxLongitude = mid
				}
			} else {
				mid := (output.MaxLatitude + output.MinLatitude) / 2
				if bit == 1 {
					output.MinLatitude = mid
				} else {
					output.MaxLatitude = mid
				}
			}
			islon = !islon
		}
	}

	return output
}

func Decode(hash_string string) DecodedPosition {
	var bbox DecodeBoundBox = DecodeBoundBox(hash_string)
	var output DecodedPosition
	// Mid point of box
	output.Latitude = (bbox.MinLatitude + bbox.MaxLatitude) / 2
	output.Longitude = (bbox.MinLongitude + bbox.MaxLongitude) / 2
	// Mid Point -  Min/Max ==> Error
	output.LatitudeError = bbox.MaxLatitude - output.Latitude
	output.LongitudeError = bbox.MaxLongitude - output.Longitude

	return output
}

func Neighbor(hash_string string, direction [2]int) string {
	// Adjust the DecodedPosition for the direction of the neighbors
	var position DecodedPosition = Decode(hash_string)
	position.Latitude += direction[0] * position.LatitudeError * 2
	position.Longitude += direction[1] * position.LongitudeError * 2

	return Encode(position.latitude, position.longitude, hash_string.length())
}
