package certificate

type Repository interface {
	RegisterCertificate(c *Certificate) (*Certificate, error)
}
