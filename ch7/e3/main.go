package main

import "fmt"

func fetchTasks(baseURL, availability string) []Issue {
	availabilityIssues := 0
	// check availablity to add to limit query parameter
	switch availability {
	case "Low":
		availabilityIssues = 1
	case "Medium":
		availabilityIssues = 3
	case "High":
		availabilityIssues = 5
	}

	fullURL := fmt.Sprintf("%s?sort=estimate&limit=%v", baseURL, availabilityIssues)
	return getIssues(fullURL)
}
