// Copyright ©2017 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package itf8

import "testing"

func TestUint32RoundTrip(t *testing.T) {
	b := make([]byte, 5)
	for i := uint(0); i < 32; i++ {
		for off := -1; off <= 1; off++ {
			in := uint32(1<<i + off)
			inn := EncodeUint32(b, in)
			out, outn, ok := DecodeUint32(b)
			if !ok {
				t.Error("failed to decode ITF-8 bytes: %08b", b[:inn])
			}
			if inn != outn {
				t.Errorf("disagreement in number of encoded bytes: in=%d out=%d", inn, outn)
			}
			if in != out {
				t.Errorf("disagreement in encoded value: in=%d (0x%[1]x) out=%d (0x%[2]x)\nencoding=%08b", in, out, b[:inn])
			}
		}
	}
}

func TestInt32RoundTrip(t *testing.T) {
	b := make([]byte, 5)
	for i := uint(0); i < 32; i++ {
		for off := -1; off <= 1; off++ {
			in := int32(1<<i + off)
			inn := EncodeInt32(b, in)
			out, outn, ok := DecodeInt32(b)
			if !ok {
				t.Error("failed to decode ITF-8 bytes: %08b", b[:inn])
			}
			if inn != outn {
				t.Errorf("disagreement in number of encoded bytes: in=%d out=%d", inn, outn)
			}
			if in != out {
				t.Errorf("disagreement in encoded value: in=%d (0x%[1]x) out=%d (0x%[2]x)\nencoding=%08b", in, out, b[:inn])
			}
		}
	}
}
