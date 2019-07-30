var global_current_lidar = ""

function lidarLoad(item){
    global_current_lidar = $(item).find("p").text()
    // set
    $('#viewFrame').attr('src', `http://localhost:8085/data/${global_current_lidar}.html`)
    //load
    $( '#viewFrame' ).attr( 'src', function ( i, val ) { return val; });
    //caption
    $('#viewer-title').text(function(){
        return `Viewing → ${global_current_lidar}` 
    })
}

function fullscreenviewer() {
    var win = window.open(`http://localhost:8085/data/${global_current_lidar}.html`, '_blank')
    win.focus();
}
    

$( window ).unload(function() {
    return $('#viewFrame').removeData('src');
  });