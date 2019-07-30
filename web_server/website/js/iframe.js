var global_current_lidar = ""

function lidarLoad(item){
    $('#full-screen-btn').attr('disabled', false)
    $('#switcher-btn').attr('disabled', false)
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

function fullscreenviewer(b) {
    var win = window.open(`http://localhost:8085/data/${global_current_lidar}.html`, '_blank')
    win.focus();
}
  
function switcher(b){
    console.log($(b).text())
    var this_text = $(b).text()
    if(this_text == 'Lidar'){
        $('#map').hide()
        $('#viewFrame').show()
        $(b).text(function(){
            return 'Map'
        })
    } else if (this_text == 'Map'){
        $('#viewFrame').hide()
        $('#map').show()
        $(b).text(function(){
            return 'Lidar'
        })
    }
}

$(window).on("unload", function() {
    return $('#viewFrame').removeData('src');
  });