package numsteps

func Calculate(n int, steps []int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	res := 0
	for _, i := range steps {
		res += Calculate(n-i, steps)
	}

	return res
}
