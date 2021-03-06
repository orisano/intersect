// Copyright 2020-2021 Nao Yonashiro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package intersect

func Naive(dest, small, large []int) []int {
	var s, l int
	for s < len(small) && l < len(large) {
		x := small[s]
		y := large[l]
		if x < y {
			s++
		} else if x > y {
			l++
		} else {
			dest = append(dest, x)
			s++
			l++
		}
	}
	return dest
}
