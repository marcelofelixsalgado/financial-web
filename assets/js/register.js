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
        
        // $.ajax({
        //     url: "/register/credentials",
        //     method: "GET",
        //     data: {
        //         user_id: data.id,
        //     }            
        // }).done(function(error) {
        //     alert("OK!")
        // }).fail(function(error) {
        //     alert("erro!")
        // })

        window.location = `/register/credentials?user_id=`+data.id;
    }).fail(function(error) {
        alert("erro!")
    });    
}