package ggeohash

import "os"
import "math"

// import "io/ioutil"
import "encoding/csv"
import "strconv"

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

// Helpers
func GeoHashSpec(c gospec.Context) {
	//
	// Verify the conversion methods return the correct input/output
	//
	c.Specify("Converts ints to bytes and back again", func() {
		// This is important to get right ...
		// The GeoHash algorithm depends on this being fast and consistent
		src := "0123456789bcdefghjkmnpqrstuvwxyz"
		for i := 0; i < 32; i++ {
			// What is the character at "i"
			src_at := byte(src[i])

			// Map the position to the byte
			byte_at := ConvertIndexToByte(i)
			c.Expect(byte_at, Equals, src_at)

			// Map the byte back to the position
			index_at := ConvertByteToIndex(byte_at)
			c.Expect(index_at, Equals, i)
		}
	})

	//
	// Verify the GeoHash methods encode/decode the input correctly
	//
	var longitude = 112.5584
	var latitude = 37.8324
	var precision = uint8(9)
	var geostr = "ww8p1r4t8"

	c.Specify("encodes latitude & longitude as string", func() {
		var actual = Encode(latitude, longitude, precision)
		var expected = geostr
		c.Expect(actual, Equals, expected)
	})

	c.Specify("decodes string to bounded box", func() {
		var actual = DecodeBoundBox(geostr)
		var expected = DecodedBoundBox{MinLatitude: 37.83236503601074, MinLongitude: 112.55836486816406, MaxLatitude: 37.83240795135498, MaxLongitude: 112.5584077835083}

		c.Expect(actual.MinLatitude, Equals, expected.MinLatitude)
		c.Expect(actual.MinLongitude, Equals, expected.MinLongitude)
		c.Expect(actual.MaxLatitude, Equals, expected.MaxLatitude)
		c.Expect(actual.MaxLongitude, Equals, expected.MaxLongitude)
	})

	c.Specify("decodes string to latitude", func() {
		var actual = Decode(geostr)
		var diff = math.Abs(latitude-actual.Latitude) < 0.0001
		c.Expect(diff, IsTrue)
	})

	c.Specify("decodes string to longitude", func() {
		var actual = Decode(geostr)
		var diff = math.Abs(longitude-actual.Longitude) < 0.0001
		c.Expect(diff, IsTrue)
	})

	c.Specify("finds neighbor to the north", func() {
		var directions = [2]int{1, 0}
		var actual = Neighbor("dqcjq", directions)
		var expected = "dqcjw"
		c.Expect(actual, Equals, expected)
	})

	c.Specify("finds neighbor to the south-west", func() {
		var directions = [2]int{-1, -1}
		var actual = Neighbor("dqcjq", directions)
		var expected = "dqcjj"
		c.Expect(actual, Equals, expected)
	})

	//
	// Batch-test the GeoHash Encode against the Node.js's output.
	//
	// It is critical this is correct and consistent across platforms.
	//
	// There are multiple imlementations of this in several languages:
	// - GO
	// - Node.JS
	// - Ruby
	// - C++
	// - etc
	//
	c.Specify("CSV of encoded latitude, longitude, and precision matches encode", func() {
		file, err := os.Open("./encode_the_world.csv")
		if nil != err {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		lines, err := reader.ReadAll()
		if nil != err {
			panic(err)
		}

		for index, line := range lines {
			if index == 0 {
				continue
			}
			i := 0
			latitude, _ := strconv.ParseFloat(line[i], 64)
			i++
			longitude, _ := strconv.ParseFloat(line[i], 64)
			i++
			precision, _ := strconv.Atoi(line[i])
			i++
			expected := line[i]
			i++

			var actual = Encode(latitude, longitude, uint8(precision))
			c.Expect(actual, Equals, expected)
		}
	})
}
