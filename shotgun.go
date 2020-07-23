package intersect

func Shotgun1(dest, small, large []int) []int {
	if len(small) == 0 {
		return dest
	}

	var s, l int
	for s < len(small) && l < len(large) {
		target1 := small[s]
		base1 := l
		for n := len(large) - l; n > 1; {
			half := n >> 1
			if large[base1+half] < target1 {
				base1 += half
			}
			n -= half
		}
		var index1 int
		if large[base1] < target1 {
			index1 = base1 + 1
		} else {
			index1 = base1
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

func Shotgun4(dest, small, large []int) []int {
	if len(small) == 0 {
		return dest
	}

	var s, l int
	for s+4 <= len(small) && l < len(large) {
		x := small[s : s+4 : s+4]
		target1 := x[0]
		target2 := x[1]
		target3 := x[2]
		target4 := x[3]
		base1 := l
		base2 := l
		base3 := l
		base4 := l
		for n := len(large) - l; n > 1; {
			half := n >> 1
			if large[base1+half] < target1 {
				base1 += half
			}
			if large[base2+half] < target2 {
				base2 += half
			}
			if large[base3+half] < target3 {
				base3 += half
			}
			if large[base4+half] < target4 {
				base4 += half
			}
			n -= half
		}

		var index1, index2, index3, index4 int
		if large[base1] < target1 {
			index1 = base1 + 1
		} else {
			index1 = base1
		}
		if large[base2] < target2 {
			index2 = base2 + 1
		} else {
			index2 = base2
		}
		if large[base3] < target3 {
			index3 = base3 + 1
		} else {
			index3 = base3
		}
		if large[base4] < target4 {
			index4 = base4 + 1
		} else {
			index4 = base4
		}
		if index1 < len(large) && large[index1] == target1 {
			dest = append(dest, target1)
		}
		if index2 < len(large) && large[index2] == target2 {
			dest = append(dest, target2)
		}
		if index3 < len(large) && large[index3] == target3 {
			dest = append(dest, target3)
		}
		if index4 < len(large) && large[index4] == target4 {
			dest = append(dest, target4)
		}
		s += 4
		l = index4
	}
	if s < len(small) && l < len(large) {
		dest = Shotgun1(dest, small[s:], large[l:])
	}
	return dest
}
