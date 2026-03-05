package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/jsonData")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	var user User
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
