package main

import (
	"fmt"
	"strings"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

type Input struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func validateRequest(req *rpc.SendRequest) error {
	splitChat := strings.Split(req.Message.GetChat(), ":")
	err := fmt.Errorf("Invalid Request Error")
	if len(splitChat) != 2 {
		return err
	}
	found := false
	for _, party := range splitChat {
		if party == req.Message.GetSender() {
			found = true
		}
	}
	if !found {
		return err
	}
	return nil
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
