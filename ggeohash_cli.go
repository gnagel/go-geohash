package main

import "fmt"
import "os"
import "strings"

import flags "github.com/jessevdk/go-flags"

// import "./ggeohash"

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	Encode bool `long:"encode" description:"Encode a geohash"`

	Decode bool `long:"decode" description:"Decode a geohash"`

	Neighbor bool `long:"neighbor" description:"Find the neighbor of a geohash in a given direction: North, South, East, West, North+East, South+West, etc"`

	Latitude  float64 `long:"latitude" description:"Latitude [-90.0 ... +90.0]"`
	Longitude float64 `long:"longitude" description:"Longitude [-180.0 ... +180.0]"`
	Precision uint8 `long:"precision" description:"Precision [1 ... 12]"`
	GeoHash string `long:"geohash" description:"GeoHash string"`
}

func init() {
}

func main() {
	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	args, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

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
