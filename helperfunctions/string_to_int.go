package helperfunctions

import "strconv"

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return result
}

func StringsToInts(s []string) []int {
	results := []int{}
	for _, item := range s {
		result, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		results = append(results, result)

	}
	return results
}
