package tools

import (
	mapset "github.com/deckarep/golang-set"
)

func DiffString(newSlice, oldSlice []string) (addSlice []string, deleteSlice []string) {
	newSet := mapset.NewThreadUnsafeSet()
	oldSet := mapset.NewThreadUnsafeSet()

	for i, _ := range newSlice {
		newSet.Add(newSlice[i])
	}

	for i, _ := range oldSlice {
		oldSet.Add(oldSlice[i])
	}

	for _, s := range newSet.Difference(oldSet).ToSlice() {
		addSlice = append(addSlice, s.(string))
	}

	for _, s := range oldSet.Difference(newSet).ToSlice() {
		deleteSlice = append(deleteSlice, s.(string))
	}

	return
}

func DiffInt(newSlice, oldSlice []int) (addSlice []int, deleteSlice []int) {
	newSet := mapset.NewThreadUnsafeSet()
	oldSet := mapset.NewThreadUnsafeSet()

	for i, _ := range newSlice {
		newSet.Add(newSlice[i])
	}

	for i, _ := range oldSlice {
		oldSet.Add(oldSlice[i])
	}

	for _, s := range newSet.Difference(oldSet).ToSlice() {
		addSlice = append(addSlice, s.(int))
	}

	for _, s := range oldSet.Difference(newSet).ToSlice() {
		deleteSlice = append(deleteSlice, s.(int))
	}

	return
}
