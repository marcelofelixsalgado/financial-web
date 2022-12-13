$('#register-credentials-form').on('submit', userRegisterCredentials);

function userRegisterCredentials(event) {
    event.preventDefault();
    
    if ($('password').val() != $('password-confirm').val()) {
        alert("The password don't match");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            password: $('password').val()
        }
    })
}