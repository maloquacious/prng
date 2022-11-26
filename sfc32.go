/*
 * prng - an assortment of psuedo-random number generators
 * Copyright (c) 2022 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package prng

// SFC32 is from https://simblob.blogspot.com/2022/05/upgrading-prng.html#more
//
// Example
//
//	sfc32 := prng.SFC32(0, 12345, 0, 1)
//	for i := 0; i < 10; i++ {
//		log.Println(sfc32())
//	}
func SFC32(a, b, c, d uint32) PRNG {
	fn := func() uint32 {
		t := (a + b) + d
		d = d + 1
		a = b ^ b>>9
		b = c + (c << 3)
		c = c<<21 | c>>11
		c = c + t
		return t
	}

	// source recommends tossing first 12 outputs
	for i := 0; i < 12; i++ {
		fn()
	}

	return fn
}
