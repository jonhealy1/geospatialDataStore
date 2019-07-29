function lidarLoad(item){
    var lid = $(item).find("p").text()
    console.log($(item).find("p").text())
    $('#viewFrame').attr('src', `http://localhost:8085/data/${lid}.html`)
    $( '#viewFrame' ).attr( 'src', function ( i, val ) { return val; });

}
    