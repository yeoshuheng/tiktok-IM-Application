package main

import (
	"context"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	err := validateRequest(req)
	if err != nil {
		return nil, err
	}

	id := getID(req.Message.GetChat())

	toSend := &Input{
		Message:   req.Message.GetText(),
		Sender:    req.Message.GetSender(),
		Timestamp: time.Now().Unix(),
	}

	// Writing to Redis Database.
	err = db.WriteToDatabase(ctx, id, toSend)
	if err != nil {
		return nil, err
	}

	// Writing to SQL DB.
	//err = db.WriteToSQLDB(ctx, id, toSend)
	//if err != nil {
	//	return nil, err
	//}

	resp := rpc.NewSendResponse()
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {

	id := getID(req.GetChat())
	strt := req.GetCursor()
	limit := req.GetLimit()
	end := strt + int64(limit)

	// Reading from Redis DB.
	rawMsg, err := db.ReadFromDatabase(ctx, id, strt, end, req.GetReverse())
	if err != nil {
		return nil, err
	}

	// Reading from SQL DB.
	//rawMsg, err := db.ReadFromSQLDB(ctx, id, strt, end, req.GetReverse())
	//if err != nil {
	//	return nil, err
	//}

	ret := make([]*rpc.Message, 0)
	var c int32 = 0
	var n int64 = 0
	hasmore := false
	for _, curr := range rawMsg {
		if c+1 > limit {
			hasmore = true
			n = end
			break
		}
		ret = append(ret, &rpc.Message{
			Chat:     req.GetChat(),
			Text:     curr.Message,
			Sender:   curr.Sender,
			SendTime: curr.Timestamp,
		})
		c++
	}

	resp := rpc.NewPullResponse()
	resp.Messages, resp.HasMore, resp.NextCursor = ret, &hasmore, &n
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}
