// test_responder_test.go
//
// Copyright (C) 2022-2024 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package nsdp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartStop(t *testing.T) {
	responder, err := NewTestResponder("127.0.0.1:63322")
	require.Nil(t, err)
	responder.AddResponses(
		"0102000000000000bcd07432b8dc6cb0ce1c8394000099d14e534450000000000001000847533130384576330003000773776974636831000400066cb0ce1c839400050000000600040a01000300070004ffff0000000800040a010001000b000100000d0007322e30362e3137000e0000000f0001010c0000030105000c0000030200000c0000030304000c0000030400000c0000030504000c0000030600000c0000030700000c0000030800001000003101000000011b86e2c2000000000d159e3800000000000000000000000000000000000000000000000000000000000000001000003102000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000310300000000039bd6ce000000000874984f000000000000000000000000000000000000000000000000000000000000000010000031040000000000133f340000000000cf6d03000000000000000000000000000000000000000000000000000000000000000010000031050000000009668768000000010afa8d1d0000000000000000000000000000000000000000000000000000000000000000100000310600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000031070000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000003108000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000")
	err = responder.Start()
	require.Nil(t, err)
	err = responder.Stop()
	require.Nil(t, err)
}
