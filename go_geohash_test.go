package go_geohash

import "math"

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

var longitude = 112.5584
var latitude = 37.8324
var geostr = "ww8p1r4t8"

// Helpers
func GeoHashSpec(c gospec.Context) {

	c.Specify("encodes latitude & longitude as string", func() {
		var actual = Encode(latitude, longitude, 9)
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

}
