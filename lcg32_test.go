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

import "testing"

func TestLCG(t *testing.T) {
	lcg32 := LCG32(0)
	for _, expect := range []uint32{1178599519, 564134195, 3263168954, 2665480396, 2227175438, 4196401256, 486539424, 56623112, 2604662946, 178258093} {
		if got := lcg32(); got != expect {
			t.Errorf("wanted %d: got %d\n", expect, got)
		}
	}

	lcg32 = LCG32(1)
	for _, expect := range []uint32{2573205670, 155189689, 2963276144, 1300234382, 367002109, 2545943034, 421527894, 2013266181, 2321547461, 4265607635} {
		if got := lcg32(); got != expect {
			t.Errorf("wanted %d: got %d\n", expect, got)
		}
	}
}
