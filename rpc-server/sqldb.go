package main

import (
	"context"
	"database/sql"
	"fmt"
)

type sqlDB struct {
	cli *sql.DB
}

func (s *sqlDB) initSQLDB() error {
	db, err := sql.Open("mysql", "username:password@tcp(8080:8080)/messageDB")
	if err != nil {
		return err
	}
	s.cli = db
	defer db.Close()
	return nil
}

func (s *sqlDB) WriteToSQLDB(ctx context.Context, id string, input *Input) error {
	sender, time, msg := input.Sender, input.Timestamp, input.Message
	cmd := fmt.Sprintf("INSERT INTO messageDB VALUES(%s, %d, %s, %s)", id, time, sender, msg)
	_, err := s.cli.Query(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlDB) ReadFromSQLDB(ctx context.Context, id string, start int64, end int64, reverse bool) ([]*Input, error) {
	var processedOutputs []*Input
	var direction string
	if reverse {
		direction = "ASC"
	} else {
		direction = "DESC"
	}
	cmd := fmt.Sprintf("SELECT * FROM messageDB WHERE id == %s ORDER BY timestamp %s", id, direction)
	results, err := s.cli.Query(cmd)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var temp Input
		err := results.Scan(temp.Timestamp, temp.Sender, temp.Message)
		if err != nil {
			return nil, err
		}
		processedOutputs = append(processedOutputs, &temp)
	}
	return processedOutputs, nil
}
