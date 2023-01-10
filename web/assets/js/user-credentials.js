$('#register-credentials-form').on('submit', userRegisterCredentials);
$('#update-credentials').on('submit', updatePassword);

function userRegisterCredentials(event) {
    event.preventDefault();
    
    if ($('#password').val() != $('#passwordConfirm').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return
    }
    
    $.ajax({
        url: "/register/credentials",
        method: "POST",
        data: {
            user_id: $('#user_id').val(),
            password: $('#password').val(),
        }
    }).done(function(error) {
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                })
            })
    }).fail(function(error) {
        Swal.fire("Ops...", "Erro ao cadastrar as credenciais do usuário!", "error");
    })    
}

function updatePassword(event) {
    event.preventDefault();

    if ($('#newPassword').val() != $('#newPasswordConfirm').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/register/credentials",
        method: "PUT",
        data: {
            current_password: $('#currentPassword').val(),
            new_password: $('#newPassword').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso!", "success")
            .then(function() {
                window.location = "/users/profile";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
    });    
}