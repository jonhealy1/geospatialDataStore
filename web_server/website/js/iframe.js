var global_current_lidar = ""

function lidarLoad(item){
    $('#full-screen-btn').attr('disabled', false)
    $('#switcher-btn').attr('disabled', false)
    global_current_lidar = $(item).find("p").text()
    // set
    if(global_current_lidar == 'hanover'){
        $('#viewFrame').attr('src', `http://localhost:8085/data/hanover.html`)
    } else if(global_current_lidar == 'beer-caves') {
        $('#viewFrame').attr('src', `http://localhost:8085/data/beer-caves.html`)
    }else{
        $('#viewFrame').attr('src', `http://localhost:8085/data/view.html?r="http://localhost:8111/${global_current_lidar}"`)
    }
    //load
    $( '#viewFrame' ).attr( 'src', function ( i, val ) { return val; });
    //caption
    $('#viewer-title').text(function(){
        return `Viewing → ${global_current_lidar}` 
    })
}

function fullscreenviewer(b) {
    if(global_current_lidar == 'hanover'){
        var win = window.open(`http://localhost:8085/data/hanover.html`, '_blank')
    } else if(global_current_lidar == 'beer-caves'){
        var win = window.open(`http://localhost:8085/data/beer-caves.html`, '_blank')
    }else{
        var win = window.open(`http://localhost:8085/data/view.html?r="http://localhost:8111/${global_current_lidar}"`, '_blank')
    }
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