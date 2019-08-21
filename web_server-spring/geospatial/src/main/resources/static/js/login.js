$(document).ready(function(){

    // $(document).ready(function(){
//     $.post({
//         url: "/deleteClients"
//     })
// })


    $('#submitForm').submit(function(e){

        e.preventDefault();

        var formData = $(this).serialize();

        // $.post("/signupClient", formData).complete(function(){
        //     console.log("Success")
        // })

        $.post({
            url: "/signupClient",
            data: formData,
            complete: function(){
                console.log("success")
            } 
        })
    })
})
