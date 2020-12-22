package types

type Certificate struct {
	Name string
	Data []byte
	Storage Storages
}

type Storages interface {
	Create(cert Certificate) (Certificate, error)
	Update(cert Certificate) (Certificate, error)
	Delete(cert Certificate) (Certificate, error)
	List() ([]Certificate, error)
}
