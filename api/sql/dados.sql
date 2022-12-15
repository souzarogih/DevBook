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