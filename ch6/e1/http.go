package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []User{}, err
	}

	var user []User
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return []User{}, err
	}
	return user, nil
}
