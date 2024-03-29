package util

func RemoveItem[T comparable](items []T, indexes []int) []T {
	indiciesMap := make(map[int]bool)
	for _, i := range indexes {
		indiciesMap[i] = true
	}
	var result []T
	for i, item := range items {
		if !indiciesMap[i] {
			result = append(result, item)
		}
	}
	return result
}

func RemoveValue[T comparable](items []T, value T) []T {
	var result []T
	for _, item := range items {
		if item != value {
			result = append(result, item)
		}
	}
	return result
}

func RemoveValues[T comparable](items []T, values []T) []T {
	valuesMap := make(map[T]bool)
	for _, v := range values {
		valuesMap[v] = true
	}
	var result []T
	for _, item := range items {
		if !valuesMap[item] {
			result = append(result, item)
		}
	}
	return result
}
