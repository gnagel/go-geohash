// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package go_geohash

import "testing"

const (
	latitude  = -33.863574
	longitude = 150.915070
	hash      = "r3gr65m42h22"
)

func Test_Encode_1(t *testing.T) {
	result := Encode(latitude, longitude, 12)
	if result != hash {
		t.Error(result, hash)
	}
}
