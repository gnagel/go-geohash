package main

import "fmt"
import "os"
import "strings"
import "./ggeohash"

import flags "github.com/jessevdk/go-flags"

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	//
	// Operation To Perform
	//
	Encode bool `long:"encode" description:"Encode a geohash"`

	Decode bool `long:"decode" description:"Decode a geohash"`

	Neighbor bool `long:"neighbor" description:"Find the neighbor of a geohash in a given direction: North, South, East, West, North+East, South+West, etc"`

	//
	// Input Parameters for above
	//

	Latitude float64 `long:"latitude" description:"Latitude [-90.0 ... +90.0]"`

	Longitude float64 `long:"longitude" description:"Longitude [-180.0 ... +180.0]"`

	Precision int `long:"precision" description:"Precision [1 ... 12]"`

	GeoHash string `long:"geohash" description:"GeoHash string"`

	North bool `long:"north" description:"Neighbor to the North"`
	South bool `long:"south" description:"Neighbor to the South"`
	East  bool `long:"east" description:"Neighbor to the East"`
	West  bool `long:"west" description:"Neighbor to the West"`
}

func init() {
}

func main() {
	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	args, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(0)
		return
	}

	if opts.Verbose {
		fmt.Printf("Verbosity: %v\n", opts.Verbose)
		fmt.Printf("Encode: %v\n", opts.Encode)
		fmt.Printf("Decode: %v\n", opts.Decode)
		fmt.Printf("Neighbor: %v\n", opts.Neighbor)
		fmt.Printf("Latitude: %v\n", opts.Latitude)
		fmt.Printf("Longitude: %v\n", opts.Longitude)
		fmt.Printf("Precision: %v\n", opts.Precision)
		fmt.Printf("GeoHash: '%v'\n", opts.GeoHash)
		fmt.Printf("Remaining args: [%s]\n", strings.Join(args, " "))
	}

	// Shorthand to make the code below readable
	lat := opts.Latitude
	lon := opts.Longitude
	precision := uint8(opts.Precision)
	geo := opts.GeoHash

	if opts.Encode {
		output := ggeohash.Encode(lat, lon, precision)
		fmt.Printf("latitude = %2.10f, longitude = %3.10f, precision = %d, geohash = %s\n", lat, lon, precision, output)
	}

	if opts.Decode {
		output := ggeohash.Decode(geo)
		fmt.Printf("geohash = %s, latitude = %2.10f, longitude = %3.10f, latitude.err = %2.10f, longitude.err = %3.10f\n", geo, output.Latitude, output.Longitude, output.LatitudeError, output.LongitudeError)
	}

	if opts.Neighbor {
		directions := [2]int{0, 0}
		if opts.North {
			directions[0] = 1
		}
		if opts.South {
			directions[0] = -1
		}
		if opts.East {
			directions[1] = 1
		}
		if opts.West {
			directions[1] = -1
		}

		output := ggeohash.Neighbor(geo, directions)
		fmt.Printf("geohash = %s, directions[0] = %d, directions[1] = %d, neighbor = %s\n", geo, directions[0], directions[1], output)
	}
}
