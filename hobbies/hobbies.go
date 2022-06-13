package main

import (
	"fmt"
)

func isContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	var hobbyists []string
	for key, v := range hobbies {
		if isContains(v, hobby) {
			hobbyists = append(hobbyists, key)
		}
	}
	return hobbyists
}
func main() {
	hobbies := map[string][]string{
		"Steve": {"Fashion", "Piano", "Reading"},
		"Patty": {"Drama", "Magic", "Pets"},
		"Chad":  {"Puzzles", "Pets", "Yoga"},
	}
	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
