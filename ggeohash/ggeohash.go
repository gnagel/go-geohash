package ggeohash

import "strings"
import "bytes"

// Static array of 0-9, a-z
var base32_codes [32]byte = [32]byte{}
var base32_map map[byte]uint8 = map[byte]uint8{}

func init() {
	characters := "0123456789bcdefghjkmnpqrstuvwxyz"
	
	// Map the runes to bytes & index positions
	for index, rune := range characters {
		byte_at := byte(rune)
		base32_codes[index] = byte_at
		base32_map[byte_at] = uint8(index)
	}
}

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
	// var output string = strings.Repeat(" ", int(precision))

	// DecodedBoundBox for the lat/lon + errors
	var bbox DecodedBoundBox = DecodedBoundBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}

	var min_max_avg float64 = 0
	var islon bool = true
	var num_bits uint = 0
	var hash_index int = 0

	var buffer bytes.Buffer
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
			buffer.WriteByte(base32_codes[hash_index])

			output_length++
			num_bits = 0
			hash_index = 0
		}
	}

	return buffer.String()
}

// Clone the Map instance
// This prevents multiple threads from accessing the same map instance con-currently
func base32_map_factory() map[byte]uint8 {
	return base32_map
}

func DecodeBoundBox(hash_string string) *DecodedBoundBox {
	// Downcase the input string
	hash_string = strings.ToLower(hash_string)

	output := &DecodedBoundBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}

	var islon bool = true

	// We can't do this as a static map in GO
	// Since the MAP type is not thread safe
	lookup := base32_map_factory()

	for _, c := range hash_string {
		byte_at := (byte(c))
		var index_of uint8 = lookup[byte_at]

		for bits := 4; bits >= 0; bits-- {
			bit := (index_of >> uint8(bits)) & 1
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

func Decode(hash_string string) *DecodedPosition {
	bbox := DecodeBoundBox(hash_string)
	output := &DecodedPosition{}
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
	position := Decode(hash_string)
	precision := uint8(len(hash_string))
	lat := position.Latitude + float64(direction[0])*position.LatitudeError*2
	lon := position.Longitude + float64(direction[1])*position.LongitudeError*2

	return Encode(lat, lon, precision)
}
