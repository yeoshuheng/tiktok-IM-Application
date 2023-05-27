package main

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type dbClient struct {
	cli *redis.Client
}

func (c *dbClient) InitDatabase(ctx context.Context) error {

	db := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := db.Ping(ctx).Result()
	if err != nil {
		return err
	}

	c.cli = db
	return nil
}

func (c *dbClient) WriteToDatabase(ctx context.Context, id string, input *Input) error {
	jsonRep, err := json.Marshal(input)
	if err != nil {
		return err
	}
	setMember := &redis.Z{
		Score:  float64(input.Timestamp),
		Member: jsonRep,
	}

	_, err = c.cli.ZAdd(ctx, id, *setMember).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *dbClient) ReadFromDatabase(ctx context.Context, id string, start int64, end int64, reverse bool) ([]*Input, error) {
	var rawOutputs []string
	var processedOutputs []*Input
	var err error
	if reverse {
		rawOutputs, err = c.cli.ZRevRange(ctx, id, start, end).Result()
	} else {
		rawOutputs, err = c.cli.ZRange(ctx, id, start, end).Result()
	}
	if err != nil {
		return nil, err
	}
	for _, input := range rawOutputs {
		temp := &Input{}
		err := json.Unmarshal([]byte(input), temp)
		if err != nil {
			return nil, err
		}
		processedOutputs = append(processedOutputs, temp)
	}
	return processedOutputs, nil
}
