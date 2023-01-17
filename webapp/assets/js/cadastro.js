$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault();

  if ($('#senha').val() != $('#confirmar-senha').val()) {
    Swal.fire("Ops...", "As senhas não coincidem!", "error");
    return;
  }

  $.ajax({
    url: '/usuarios',
    method: 'POST',
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val(),
    },
  })
    .done(function () {
      Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!!", "success");
    })
    .fail(function (erro) {
      console.log(erro);
      Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });
}
