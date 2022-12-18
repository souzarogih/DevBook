insert into usuarios (nome, nick, email, senha)
values
("Higor Souza","higor.souza","higor@hotmail.com","$2a$10$RnnQ8msf6oZQm1DaJs7XOu66BJUIdN/qy2O8GtiF34ctYNE5jXlmq"),
("Rogih Souza","rogih.souza","rogih@hotmail.com","$2a$10$RnnQ8msf6oZQm1DaJs7XOu66BJUIdN/qy2O8GtiF34ctYNE5jXlmq"),
("São João","saint.john","sain_john@hotmail.com","$2a$10$RnnQ8msf6oZQm1DaJs7XOu66BJUIdN/qy2O8GtiF34ctYNE5jXlmq");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa e a publicação do usuário 1! Oba", 1),
("Publicação do usuário 2", "Essa e a publicação do usuário 2! Oba", 2),
("Publicação do usuário 3", "Essa e a publicação do usuário 3! Oba", 3),
("Publicação do usuário 5", "Essa e a publicação do usuário 5! Oba", 5)