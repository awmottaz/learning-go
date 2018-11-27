// Benchmark results:
//	solution	10000000	199 ns/op	240 B/op	4 allocs/op
//	solutionB	10000000	262 ns/op	128 B/op	1 allocs/op

package dcp002

// This solution runs faster (24%)
func solution(list []int) (result []int) {
	if len(list) < 2 {
		return
	}
	result = []int{list[1], list[0]}
	nextEntry := list[0] * list[1]
	for i := 2; i < len(list); i++ {
		entry := list[i]
		for j := range result {
			result[j] *= entry
		}
		result = append(result, nextEntry)
		nextEntry *= entry
	}
	return
}

// This solution uses less memory (46.7%)
func solutionB(list []int) (result []int) {
	if len(list) < 2 {
		return
	}
	result = make([]int, len(list))
	for i := range result {
		result[i] = 1
	}
	for i, v := range list {
		for j := range result {
			if i == j {
				continue
			}
			result[j] *= v
		}
	}
	return
}
