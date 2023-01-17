$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
  evento.preventDefault();

  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      senha: $('#senha').val(),
    },
  })
    .done(function () {
      window.location = '/home';
    })
    .fail(function (erro) {
      console.log('erro--->>>', erro);
      Swal.fire("Ei!!!!", "Usuário ou senha invalidos!!", "error");
    });
}
