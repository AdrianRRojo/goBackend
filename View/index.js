var btn = document.getElementById('btn')
console.log(btn)

$(document).ready( function(){
    $('#btn').click(function(e){
        e.preventDefault();

        $('#btn').prop('disabled', true);

        submit()

        setTimeout(function () { $('#btn').prop('disabled', false); }, 500);
    })
});

function submit(){
    console.log("Submit");
}