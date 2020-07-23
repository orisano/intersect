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
