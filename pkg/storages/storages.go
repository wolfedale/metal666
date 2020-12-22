package storages

import (
	"fmt"
	"github.com/wolfedale/metal666/pkg/storages/types"
)

func ListAllCertificates(storage types.Storages) ([]types.Certificate, error) {
	fmt.Println(storage.List())
	return nil, nil
}
