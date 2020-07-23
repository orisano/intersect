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