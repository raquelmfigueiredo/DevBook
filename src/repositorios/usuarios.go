package repositorios

import (
	"api/API/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB

}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuaios (nome, nick, email, senha) valeus(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}