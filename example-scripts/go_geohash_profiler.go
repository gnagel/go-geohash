package main

import "time"
import "fmt"
import "./go_geohash"

var longitude = 112.5584
var latitude = 37.8324
var geostr = "ww8p1r4t8"
var neighbor = "dqcjq"
var directions = [2]int{1, 0}
var num_loops int = 1000 * 1000
var format_tag = "[%-30s][%-5s]\t MS Per/Call %.4f (ms)\t Total for %d Calls = %5f (ms)\n"

func main() {
	var start time.Time
	var ms_total float64
	var ms_per_call float64

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Encode(latitude, longitude, 9)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "encode", "GO", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.DecodeBoundBox(geostr)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "decode_bbox", "GO", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Decode(geostr)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "decode", "GO", ms_per_call, num_loops, ms_total)

	start = time.Now()
	for i := 0; i < num_loops; i++ {
		go_geohash.Neighbor(neighbor, directions)
	}
	ms_total = float64(time.Now().Sub(start).Nanoseconds() / int64(1000000))
	ms_per_call = ms_total / float64(num_loops)
	fmt.Printf(format_tag, "neighbor", "GO", ms_per_call, num_loops, ms_total)
}
