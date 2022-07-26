
/*Добавление Школы*/
async function SchoolAdd() {
    const {value: School} = await Swal.fire({
        title: 'Добавить школу',
        input: 'text',
        inputPlaceholder: 'Название школы',
        showCancelButton: true,
        inputName: 'school'
    })

    if (School) {
        $.ajax({
            type: 'POST',
            url: 'ajax/schoolAdd',
            data: {name: School},
            success: function () {
                Swal.fire(`Добалена`)
                $("#school").empty();
                $("#school").append('<option>Выберите школу</option>');
                SchoolGet()
            }
        });
    }
}

/*Get schools*/
async function SchoolGet() {
    $.ajax({
        url: 'ajax/schoolGet',
        type: "GET",
        success: function (results) {
            let Sc = results.schools
            for(let i = 0; i < Sc.length; i++){
                $("#school").append('<option>' + Sc[i]["Name"] + '</option>');
            }
        }
    });
}


/*Регистрация user*/
$(function () {
    $("#btn_sign_up").click(function (e) {
        var FirstName = $("#name").val(),
            Surname = $("#surname").val(),
            Email = $("#email").val(),
            Patronymic = $("#patronymic").val(),
            School = $("#school").val(),
            Class = $("#class").val(),
            Password = $("#password").val();
        $.ajax({
            type: 'POST',
            url: 'ajax/createUser',
            data: {
                school: School, name: FirstName, surname: Surname,
                email: Email, class: Class, patronymic: Patronymic, password: Password
            },
            success: function () {
            }
        });
        e.preventDefault();
    });
});





$(document).ready(function(){
    SchoolGet();
});