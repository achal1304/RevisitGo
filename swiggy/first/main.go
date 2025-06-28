package first

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	numr := len(chargeTimes)
	l, r := 0, 0

	result := 0

	q := make([]int, 0)

	// if over budget-> reduce one from left , increse from right
	// increase the size of the window when in budget
	// note the window size
	curcost := 0
	for r < numr {
		for len(q) > 0 && chargeTimes[q[len(q)-1]] <= chargeTimes[r] {
			q = q[:len(q)-1]
		}
		curcost += runningCosts[r]
		q = append(q, r)
		c := int64(chargeTimes[q[0]] + ((r - l + 1) * curcost))
		if c > budget {
			curcost -= runningCosts[l]
			if q[0] == l {
				q = q[1:]
			}
			l++
		}

		if r-l+1 > result {
			result = r - l + 1
		}

		r++
	}

	return result
}
