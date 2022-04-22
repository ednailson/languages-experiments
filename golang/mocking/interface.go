package mocking

import "time"

type Store interface {
	Open(at time.Time) error
	Sell(id string, quantity int) error
	Get(name string) (string, error)
	Close(at time.Time) error
}
