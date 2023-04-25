insert into usuarios (nome, nick, email, senha)
values
("Developer", "developer", "usuario_teste_dev@gmail.com", "$2a$10$IASye6wKEDeJ9qTTz36mjugueB9NjLfcbY8YRJoiyEDLEpS4AocyG"),
("UsuarioPadrao", "usuario_padrao", "usuario_padrao@gmail.com", "$2a$10$IASye6wKEDeJ9qTTz36mjugueB9NjLfcbY8YRJoiyEDLEpS4AocyG"),
("Usuario", "usuario", "usuario@gmail.com", "$2a$10$IASye6wKEDeJ9qTTz36mjugueB9NjLfcbY8YRJoiyEDLEpS4AocyG");
("RenatoBorges", "idreeh", "rdriano03@gmail.com", "$2a$10$IASye6wKEDeJ9qTTz36mjugueB9NjLfcbY8YRJoiyEDLEpS4AocyG");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(1, 3),
(2, 1),
(2, 3),
(3, 1),
(3, 2);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do usuário 1", "Essa é a publicação do usuário 1", 1),
("Publicação do usuário 2", "Essa é a publicação do usuário 2", 2),
("Publicação do usuário 3", "Essa é a publicação do usuário 3", 3);