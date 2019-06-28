var mymap = L.map('mapid').setView([51.505, -0.09], 3);

L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 18,
    minZoom: 3,
    id: 'mapbox.streets',
    accessToken: 'pk.eyJ1Ijoib290aGViYWxsc29vIiwiYSI6ImNqeGZpbGRheDA1c3Uzc284Y20xMGh3eWoifQ.5BZj7OvzM6cn6whqRl37XQ'
}).addTo(mymap);



// Test area plot
var testArea = L.polygon([
    [47.780557,-20.126939],
    [48.930517,-20.301523],
    [49.177531,-19.240716],
    [48.035511,-19.067345],
    [47.780557,-20.126939]
]).addTo(mymap);
testArea.setStyle({color: 'red'});
testArea.setStyle({fillColor: '#f03'});
testArea.setStyle({fillOpacity: '0.5'});