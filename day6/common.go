package day6

func countFish(generations *map[int]int) (count int) {
	for _, partialCount := range *generations {
		count += partialCount
	}

	return
}

func simulate(generations *map[int]int, days int) int {
	for days >= 0 {
		days--


		nextGenerations := map[int]int{}
		for timer, count := range *generations {
			newTimer := timer - 1
			if newTimer < 0 {
				updateGenerations(&nextGenerations, 6, count)
				updateGenerations(&nextGenerations, 8, count)
			} else {
				updateGenerations(&nextGenerations, newTimer, count)
			}
		}

		*generations = nextGenerations
	}

	return countFish(generations)
}

func updateGenerations(generations *map[int]int, timer int, members int) {
	if _, ok := (*generations)[timer]; ok {
		(*generations)[timer] += members
	} else {
		(*generations)[timer] = members
	}
}
