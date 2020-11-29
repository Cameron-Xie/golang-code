package merge

func Sort(s []int) []int {
	l := len(s)
	if l == 1 {
		return s
	}

	m := l / 2
	return merge(Sort(s[:m]), Sort(s[m:]))
}

func merge(l, r []int) []int {
	res := make([]int, len(l)+len(r))
	i := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] > r[0] {
			res[i] = r[0]
			r = r[1:]
			i++
			continue
		}

		res[i] = l[0]
		l = l[1:]
		i++
	}

	for _, j := range r {
		res[i] = j
		i++
	}

	for _, j := range l {
		res[i] = j
		i++
	}

	return res
}
