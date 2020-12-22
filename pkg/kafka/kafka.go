package kafka

import (
	"github.com/wolfedale/metal666/pkg/storages/secret"
	"github.com/wolfedale/metal666/pkg/storages/types"
)

func Run() {
	storage := secret.New()

	cert := types.Certificate{}
	cert.Name = "blah"
	cert.Data = []byte("blah")
	cert.Storage = storage

	cert.Storage.Create()
}