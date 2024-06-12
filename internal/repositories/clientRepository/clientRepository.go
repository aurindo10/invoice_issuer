package clientrepository

import (
	"database/sql"

	nfeentitie "github.com/aurindo10/invoice_issuer/internal/entities/nfeEntitie"
	"github.com/google/uuid"
)

type OpticaRepository struct {
	db *sql.DB
}

func (c *OpticaRepository) GetClientInfo(id uuid.UUID) (*nfeentitie.Dest, error)
func (c *OpticaRepository) GetCompanyInfo(id uuid.UUID) (*nfeentitie.Emit, error)
func NewOpticaRepository(db *sql.DB) *OpticaRepository {
	return &OpticaRepository{
		db: db,
	}
}
