$('#login-form').on('submit', userLogin);

function userLogin(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            name: $('#email').val(),
            email: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(error) {
        alert("erro!")
    });    
}