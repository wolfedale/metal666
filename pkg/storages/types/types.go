package types

type Certificate struct {
	Name string
	Data []byte
}

type Storager interface {
	Create(cert Certificate) (Certificate, error)
	Update(cert Certificate) (Certificate, error)
	Delete(cert Certificate) (Certificate, error)
	List() ([]Certificate, error)
}
