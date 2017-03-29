// Copyright ©2017 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// gen_nlo_const outputs candidate constants and a deBruijn look up table
// for implementing a fast bit fiddling number of leading ones function.
package main

import "fmt"

func nlo(x, add, mult, shift byte, andnot bool) int {
	x = ^x
	x |= x >> 1
	x |= x >> 2
	if andnot {
		x &^= x >> 4
	} else {
		x |= x >> 4
	}
	x += add
	x *= mult
	x >>= shift
	return int(x)
}

func loopNLO(x byte) int {
	var n int
	for b := 0x80; b > 0; b >>= 1 {
		if x&byte(b) == 0 {
			break
		}
		n++
	}
	return n
}

func main() {
	for add := byte(0); add < 2; add++ {
		for mult := byte(1); mult < 64; mult++ {
			for shift := byte(0); shift < 8; shift++ {
				for _, andnot := range []bool{false, true} {
					hits := make(map[int]int)
					for x := 0; x < 256; x++ {
						i := nlo(byte(x), add, mult, shift, andnot)
						prev, ok := hits[i]
						n := loopNLO(byte(x))
						if ok && n != prev {
							goto fail
						}
						hits[i] = n
					}
					if len(hits) == 9 {
						var max int
						for k := range hits {
							if k > max {
								max = k
							}
						}
						if max < 16 {
							fmt.Printf("add=%d mult=%d shift=%d andnot=%t table=%#v\n", add, mult, shift, andnot, hits)
						}
					}
				fail:
				}
			}
		}
	}
}
