package rotas

import (
	"api/API/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/usuarioId",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuarios,
		RequerAutenticacao: false,
	},
}

