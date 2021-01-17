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

func Galloping(dest, small, large []int) []int {
	if len(small) == 0 {
		return dest
	}

	var s, l int
	for s < len(small) && l < len(large) {
		target1 := small[s]
		base1 := l
		n := len(large) - l

		bound := 1
		for bound < n && large[base1+bound] < target1 {
			bound <<= 1
		}
		ok := base1 + (bound >> 1)
		ng := base1 + bound
		if ng > len(large) {
			ng = len(large)
		}
		for ng-ok > 1 {
			mid := (ok + ng) >> 1
			if large[mid] < target1 {
				ok = mid
			} else {
				ng = mid
			}
		}
		var index1 int
		if large[ok] < target1 {
			index1 = ok + 1
		} else {
			index1 = ok
		}
		if index1 < len(large) && large[index1] == target1 {
			dest = append(dest, target1)
			l = index1 + 1
		} else {
			l = index1
		}
		s += 1
	}
	return dest
}
