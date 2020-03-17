package storage

import (
	"context"
	"errors"
	"sync"
)

// MyDB - my in memory database
type MyDB interface {
	Set(ctx context.Context, key, value interface{}) error
	Get(key interface{}) (interface{}, error)
	Delete(key interface{}) error
}

type DB struct {
	MyDB
	m *sync.Map
}

func NewClient() *DB {
	return &DB{
		m: new(sync.Map),
	}
}

func (db *DB) Set(ctx context.Context, key, value interface{}) error {
	err := ctx.Err()
	if err != nil {
		return errors.New("Context Cancelled")
	}
	db.m.Store(key, value)
	return nil
}

func (db *DB) Get(ctx context.Context, key interface{}) (interface{}, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.New("Context Cancelled")
	}
	value, ok := db.m.Load(key)
	if !ok {
		return nil, errors.New("Key not found")
	}
	return value, nil
}

func (db *DB) Delete(ctx context.Context, key interface{}) error {
	err := ctx.Err()
	if err != nil {
		return errors.New("Context Cancelled")
	}
	db.m.Delete(key)
	return nil
}
