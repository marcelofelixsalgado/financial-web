$('#register-form').on('submit', userRegister);

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