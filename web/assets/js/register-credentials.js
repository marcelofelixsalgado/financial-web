$('#register-credentials-form').on('submit', userRegisterCredentials);

function userRegisterCredentials(event) {
    event.preventDefault();
    
    if ($('#password').val() != $('#password-confirm').val()) {
        alert("The password don't match");
    } else {
        $.ajax({
            url: "/register/credentials",
            method: "POST",
            data: {
                user_id: $('#user_id').val(),
                password: $('#password').val(),
            }
        }).done(function(error) {
            window.location = "/home"
        }).fail(function(error) {
            alert("erro!")
        })
    }
}