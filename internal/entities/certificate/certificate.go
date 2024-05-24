package certificate

import "time"

type Certificate struct {
	Certificate []byte    `json:"certificate"`
	Password    string    `json:"password"`
	Cnpj        string    `json:"cnpj"`
	UploadedAt  time.Time `json:"uploaded_at"`
	Email       string    `json:"email"`
}
