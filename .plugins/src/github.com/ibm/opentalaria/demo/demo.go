package demo

import "log"

type Demo struct {
	A string
}

func New() (*Demo, error) {
	return &Demo{
		A: "doh",
	}, nil
}

func (d *Demo) Call() error {
	log.Println("hello from plugin")

	return nil
}
