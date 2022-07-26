

function GetPath(){
    var path = document.getElementById("path").innerText;
    return path;
}


$(document).ready(function(){
    var Dir = "New folder";
    var path;
    if(GetPath === "/"){
        path = "/" + Dir;
    }
    else{ path =  GetPath() + "/" + Dir;}
  console.log(path + "/SomePath")
});


//Просмотр пути
$('#addDir').click(async function () {
    const {value: Dir} = await Swal.fire({
        title: 'Название папки',
        input: 'text',
        inputPlaceholder: 'Название папки',
        showCancelButton: true,
        inputName: 'dir'
    })
if(Dir){
    var NewDir;
    if(GetPath === "/"){
        NewDir = "/" + Dir;
    }
   else {NewDir = GetPath() + "/" + Dir}
    $.ajax({
        type: 'Post',
        async: true,
        url: '/ajax/makeDir',
        processData: false,
        contentType: false,
        data: {dir: NewDir},
        success: function (){
            alert("Папка создана")
        }
        });
}
});

// Заполняем переменную данными, при изменении значения поля file
$('#file_upload').on('change', function (){

   var file = this.files[0];
/*    event.stopPropagation();
    event.preventDefault();*/


    //ничего не делаем если файл пустой
    //if( typeof files == 'undefined') return;

    // создадим объект данных формы
    var data = new FormData();
    var path = GetPath();
    data.append('file', file);
    data.append('path', path);

    //return;

    // заполняем объект данных файлами в подходящем для отправки формате
/*    $.each( files, function( key, value ){
        data.append( key, value );
    });*/

    // добавим переменную для идентификации запроса

    $.ajax({
        type: 'POST',
        async: true,
        url: '/ajax/addFile',
        processData: false,
        contentType: false,
        data: data,
        success: function (){
            alert("Ура!!")
        }
    });
})
