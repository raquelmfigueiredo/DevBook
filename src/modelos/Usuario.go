package modelos

import "time"

// Usuario representa a urilizacao do usuario.
type Usuario struct {
	ID       uint64    `json:"id, omitempy"`
	Nome     string    `json:"nome, omitempty"`
	Nick     string    `json:"nick, omitempty"`
	Email    string    `json:"email, omitempty"`
	Senha    string    `json:"senha, omitempty"`
	CriadoEm time.Time `json:"criadoEm, omitempty"`
}
