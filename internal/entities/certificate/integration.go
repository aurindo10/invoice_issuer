package certificate

type Integration interface {
	RegisterCertificate(c *Certificate) error
}
