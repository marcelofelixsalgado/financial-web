$('#register-form').on('submit', userRegister);
$('#update-user').on('submit', updateUser);
$('#delete-user').on('click', deleteUser);

function userRegister(event) {
    event.preventDefault();

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            phone: $('#phone').val(),
            email: $('#email').val(),
        }
    }).done(function(data) {
        window.location = `/register/credentials?user_id=`+data.id+`&email=`+data.email;
    }).fail(function(error) {
        Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });    
}

function updateUser(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/users",
        method: "PUT",
        data: {
            name: $('#name').val(),
            phone: $('#phone').val(),
            email: $('#email').val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
            .then(function() {
                window.location = "/users/profile";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    });
}

function deleteUser() {
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmation) {
        if (confirmation.value) {
            $.ajax({
                url: "/users",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Sucesso!", "Seu usuário foi excluído com sucesso!", "success")
                    .then(function() {
                        window.location = "/logout";
                    })
            }).fail(function() {
                Swal.fire("Ops...", "Ocorreu um erro ao excluir o seu usuário!", "error");
            });
        }
    })
}