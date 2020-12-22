package secret

import (
	"fmt"
	"github.com/wolfedale/metal666/pkg/storages/types"
)

type Secret struct {
	Name       string
	Data       map[string][]byte
	SecretType string
}

func New() Secret {
	return Secret{}
}

func (s *Secret) Create(cert types.Certificate) (types.Certificate, error) {
	fmt.Println("Create() method from Secret pkg")
	return types.Certificate{}, nil
}

func (s *Secret) Update(cert types.Certificate) (types.Certificate, error) {
	fmt.Println("Update() method from Secret pkg")
	return types.Certificate{}, nil
}

func (s *Secret) Delete(cert types.Certificate) (types.Certificate, error) {
	fmt.Println("Delete() method from Secret pkg")
	return types.Certificate{}, nil
}

func (s *Secret) List() ([]types.Certificate, error) {
	fmt.Println("List() method from Secret pkg")
	return []types.Certificate{}, nil
}

func (s *Secret) Get(name string) (*types.Certificate, error) {
	return &types.Certificate{}, nil
}
