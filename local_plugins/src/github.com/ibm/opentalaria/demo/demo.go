package demo

import "log"

type Demo struct {
}

func New() (*Demo, error) {
	return &Demo{}, nil
}

func (d *Demo) Call() error {
	log.Println("hello from plugin")

	return nil
}
