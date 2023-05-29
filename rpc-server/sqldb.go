package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type sqlDB struct {
	cli *sql.DB
}

func (s *sqlDB) initSQLDB(ctx context.Context) error {
	db, err := sql.Open("mysql", "root:password@tcp(sql:3306)/")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS messageDB")
	if err != nil {
		return err
	}
	_, err = db.Exec("USE messageDB")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id VARCHAR(255), timestamp BIGINT, sender VARCHAR(255), message VARCHAR(255))")
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	s.cli = db
	return nil
}

func (s *sqlDB) WriteToSQLDB(ctx context.Context, id string, input *Input) error {
	sender, time, msg := input.Sender, input.Timestamp, input.Message
	cmd := fmt.Sprintf("INSERT INTO messages VALUES(%s, %d, %s, %s)", id, time, sender, msg)
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
	cmd := fmt.Sprintf("SELECT timestamp, sender, message FROM messages WHERE id == %s ORDER BY timestamp %s", id, direction)
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
