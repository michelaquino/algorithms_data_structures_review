package main

func matchingStrings(strings []string, queries []string) []int32 {
	queriesMap := map[string]int32{}

	for _, query := range queries {
		queriesMap[query] = 0
	}

	for _, str := range strings {
		_, contain := queriesMap[str]

		if contain {
			queriesMap[str] += 1
		}
	}

	result := []int32{}
	for _, query := range queries {
		result = append(result, queriesMap[query])
	}

	return result
}
