package tokenrepository

import (
	"time"

	companyentitie "github.com/aurindo10/invoice_issuer/internal/entities/CompanyEntitie"
	tokenentitie "github.com/aurindo10/invoice_issuer/internal/entities/tokenEntitie"
	"github.com/aurindo10/invoice_issuer/pkg/utils"
	"github.com/golang-jwt/jwt"
)

type TokenRepository struct {
}

func (c *TokenRepository) CreteToken(p *companyentitie.Company) (*tokenentitie.Token, error) {
	claims := jwt.MapClaims{
		"CNPJ":           p.CNPJ,
		"RAZAO_SOCIAL":   p.RAZAO_SOCIAL,
		"Owner":          p.Owner,
		"FoundationDate": p.FoundationDate.Unix(),
		"exp":            time.Now().Add(time.Hour * 200).Unix(), // Expira em 72 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := utils.GetEnv("CLERK_SECRET_KEY", "KJKJSDNSAJHLKASLKDJASHDKJ")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &tokenentitie.Token{
		Token: tokenString,
	}, nil
}
func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}
