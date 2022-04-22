package mocking

import "time"

type implementation struct {
	store Store
}

func NewImplementation(store Store) *implementation {
	return &implementation{store: store}
}

func (i *implementation) MethodA(at time.Time) error {
	err := i.store.Open(at)
	if err != nil {
		return err
	}

	err = i.store.Sell("a", 2)
	if err != nil {
		return err
	}

	err = i.store.Close(time.Now().UTC())
	if err != nil {
		return err
	}

	return nil
}
