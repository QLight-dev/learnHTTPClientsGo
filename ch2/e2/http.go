package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return []Issue{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	issues := []Issue{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil, errors.New("error decoding response body")
	}
	for i, issue := range issues {
		issues[i] = issue
	}
	return issues, nil
}
