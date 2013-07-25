//
// Unit Test Suite
//

package ggeohash

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(GeoHashSpec)

	// Execute the suite
	gospec.MainGoTest(r, t)
}
