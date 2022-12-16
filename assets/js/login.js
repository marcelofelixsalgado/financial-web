$('#login-form').on('submit', userLogin);

function userLogin(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(error) {
        alert("erro!")
    });    
}