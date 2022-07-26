async function UserTableGet(){
    $.ajax({
    url: '/ajax/userTable',
    type: "GET",
    success: function (results) {
        let User = results.u
        let Category = results.c
        let Link = results.l
        
        User.forEach(function(user){
            let catout = "";
            let check = 'checked';
            Link.forEach(function(link){
                Category.forEach(function(cat){
                    if(user.Id == link.User_id){
                        if(link.Category_id == cat.ID){
                            catout += cat.Name + ';';
                        }
                    }
                })
            })


            if(user.Status == "Y"){
                var status = '<td className="cell"><input className="form-check-input" name="" type="checkbox" id="flexSwitchCheckChecked" checked></td>'
            }else {
               status = '<td className="cell"><input className="form-check-input" name="" type="checkbox" id="flexSwitchCheckChecked"></td>'
            }
            $("#main_info").append('<tr>' +
                '<td class="cell">' + user.Id + '</td>' +
                '<td class="cell">'  + user.Name + user.Surname +'</td>' +
                '<td class="cell">' + user.Email  + '</td>' +
                  status   +
                '<td class="cell">'+ catout +'</td>' +
                '<td class="cell" id="list">' +
                '<input class="form-check-input" type="checkbox" value="" checked>' +
                '<label class="form-check-label"> Судья |</label>' +
                '<input class="form-check-input" type="checkbox" value="">' +
                '<label class="form-check-label"> Ведущий |</label>' +
                '<input class="form-check-input" type="checkbox" value="">' +
                '<label class="form-check-label"> Creator</label>' +
                '</td> </tr>');

        })
    }
    });
}

$(document).ready(function(){
    UserTableGet();
});


function update (id){

}