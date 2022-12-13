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
    }).done(function() {
        $.ajax({
            url: "/register/credentials",
            method: "POST",
        })
        window.location = `/register/credentials`;
    }).fail(function(error) {
        alert("erro!")
    });    
}