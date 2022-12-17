package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
	"log"
)

// Usuarios representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statemtn, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statemtn.Close()

	resultado, erro := statemtn.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar tras todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	log.Print("Função Buscar das respostas em execução")
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	log.Print("Retorno da função Buscar")
	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		log.Print("Erro ao tentar atualizar o usuário - 116")
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		log.Print("Erro ao executar a atualização do usuário - 117")
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		log.Printf("Erro ao preparar o insert dos dados - 131")
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		log.Printf("Erro executar o insert no banco de dados - 132")
		return erro
	}

	log.Printf("repositorios.Seguir | Usuário %d está seguindo usuário %d", seguidorID, usuarioID)
	return nil
}

// PararDeSeguir permite que um usuário pare de seguir outro.
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		log.Printf("Erro ao preparar o delete dos dados - 133")
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		log.Printf("Erro executar o insert no banco de dados - 134")
		return erro
	}

	log.Printf("repositorios.PararDeSeguir | Usuário %d está deixando de seguir o usuário %d", seguidorID, usuarioID)
	return nil
}

// BuscarSeguidores trás todos os seguidores d eum usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id 
		where s.usuario_id = ?`,
		usuarioID)
		if erro != nil {
			return nil, erro
		}
		defer linhas.Close()

		var usuarios []modelos.Usuario
		for linhas.Next() {
			var usuario modelos.Usuario

			if erro = linhas.Scan(
				&usuario.ID,
				&usuario.Nome,
				&usuario.Nick,
				&usuario.Email,
				&usuario.CriadoEm,); erro != nil {
					return nil, erro
				}

				usuarios = append(usuarios, usuario)
		}

		return usuarios, nil
}

// BuscarSeguindo trás todos os usuários que um determinado usuário está seguindo.
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u inner join seguidores s on u.id = s.usuario_id 
	where s.seguidor_id = ?`, usuarioID,)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			); erro != nil {
				return nil, erro
			}
			usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		log.Printf("repositorios.AtualizarSenha | Erro ao tentar preparar o update para alterar a senha do usuário")
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		log.Printf("repositorios.AtualizarSenha | Erro ao tentar executar o update para alterar a senha do usuário")
		return erro
	}

	return nil
}