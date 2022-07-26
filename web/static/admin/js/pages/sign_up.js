$(function () {
    /*Регистрация админ*/
    $("#admin_sign_up").click(function (e) {
        var FirstName = $("#signup-name").val(),
            Surname = $("#signup-surname").val(),
            Email = $("#signup-email").val(),
            Login = $("#signup-login").val(),
            Password = $("#signup-password").val();
        $.ajax({
            type: 'POST',
            url: '/ajax/createAdmin',
            data:  {
                name: FirstName, surname: Surname,
                email : Email, login: Login , password: Password
            },
            success: function () {

            }
        });
        e.preventDefault();
    });
});