package main

import (
	"github.com/wolfedale/metal666/pkg/storages/types"
)

type exec struct {
	cert    types.Certificate
	storage types.Storager
}

//storage = secret.New()

func New(storage types.Storager) *exec {
	return &exec{
		cert: types.Certificate{
			Name: "blah",
			Data: []byte("blah"),
		},
		storage: storage,
	}
}

func (e *exec) Create() {
	e.storage.Create(e.cert)
}