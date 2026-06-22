package entity

type CredentialRefType string

const (
	CredentialRefTypeEnv CredentialRefType = "env"
)

type CredentialRef struct {
	Type CredentialRefType
	Name string
}
