package client

type Credential struct {
	AccessID  string
	SecretKey string
}

func NewCredential(accessID, secretKey string) *Credential {
	return &Credential{
		AccessID:  accessID,
		SecretKey: secretKey,
	}
}
