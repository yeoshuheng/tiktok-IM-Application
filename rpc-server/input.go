package main

import (
	"fmt"
	"strings"
)

type Input struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func getID(chat string) string {
	cleaned := strings.Split(strings.ToLower(chat), ":")
	p1, p2 := cleaned[0], cleaned[1]
	var id string
	if p1 > p2 {
		id = fmt.Sprintf("%s:%s", p1, p2)
	} else {
		id = fmt.Sprintf("%s:%s", p2, p1)
	}
	return id
}
