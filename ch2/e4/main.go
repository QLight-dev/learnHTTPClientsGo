package main

import "encoding/json"

func marshalAll[T any](items []T) ([][]byte, error) {
	var result [][]byte
	for _, v := range items {
		marshaledData, err := json.Marshal(v)
		if err != nil {
			return [][]byte{}, err
		}
		result = append(result, marshaledData)
	}
	return result, nil
}
