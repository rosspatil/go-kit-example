package storage

import (
	"context"
	"sync"
)

// MyDB - my in memory database
type MyDB interface {
	Set(ctx context.Context, key, value interface{}) error
	Get(key interface{}) (interface{}, error)
	Delete(key interface{}) error
}

// DB -
type DB struct {
	MyDB
	m *sync.Map
}

// NewClient -
func NewClient() *DB {
	return &DB{
		m: new(sync.Map),
	}
}

// Set -
func (db *DB) Set(ctx context.Context, key, value interface{}) error {
	err := ctx.Err()
	if err != nil {
		return err
	}
	db.m.Store(key, value)
	return nil
}

// Get -
func (db *DB) Get(ctx context.Context, key interface{}) (interface{}, error) {
	err := ctx.Err()
	if err != nil {
		return nil, err
	}
	value, ok := db.m.Load(key)
	if !ok {
		return nil, nil
	}
	return value, nil
}

// Delete -
func (db *DB) Delete(ctx context.Context, key interface{}) error {
	err := ctx.Err()
	if err != nil {
		return err
	}
	db.m.Delete(key)
	return nil
}
