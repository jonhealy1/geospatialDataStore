var catalogList = [];

$(document).ready(function() {
    $.post({
        url: "/catalogs",
        success: initiateCatalogs,
        complete: function() {
            loadComplete.push("catalogs");
        }
    })
})


function initiateCatalogs(catalogs){
    catalogList = catalogs;
}