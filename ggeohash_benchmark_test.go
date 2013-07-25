package ggeohash

import "testing"

//
// Benchmark the tests
//
func Benchmark_Encode(b *testing.B) {
	var precision uint8 = 12
	var longitude = 112.5584
	var latitude = 37.8324
	for i := 0; i < b.N; i++ {
		Encode(latitude, longitude, precision)
	}
}

func Benchmark_DecodeBoundBox(b *testing.B) {
	var geostr = "ww8p1r4t8"
	for i := 0; i < b.N; i++ {
		DecodeBoundBox(geostr)
	}
}

func Benchmark_Decode(b *testing.B) {
	var geostr = "ww8p1r4t8"
	for i := 0; i < b.N; i++ {
		Decode(geostr)
	}
}

func Benchmark_NeighborNorth(b *testing.B) {
	var neighbor = "dqcjq"
	var directions = [2]int{1, 0}
	for i := 0; i < b.N; i++ {
		Neighbor(neighbor, directions)
	}
}

func Benchmark_NeighborSouthWest(b *testing.B) {
	var neighbor = "dqcjq"
	var directions = [2]int{-1, -1}
	for i := 0; i < b.N; i++ {
		Neighbor(neighbor, directions)
	}
}
