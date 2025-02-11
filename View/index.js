var btn = document.getElementById('btn')
console.log(btn)

$(document).ready( function(){
    $('#btn').click(function(e){
        e.preventDefault();

        $('#btn').prop('disabled', true);

        submit();

        setTimeout(function () { $('#btn').prop('disabled', false); }, 500);
    })
    
});

function submit(){
    var Fname = $('#Fname').val();
    var Lname = $('#Lname').val();
    var Dob = $('#Dob').val();

    var jsonData = {
        Fname: Fname,
        Lname: Lname,
        Dob: Dob
    }
    
    
    try{
        fetch('http://localhost:8200/submit',{
            method: 'POST',
            body: JSON.stringify(jsonData)
        })
        .then(response => response.text())
        .then(data => alert(data));
    } catch {
        console.log("Error: ", error);
    }
}