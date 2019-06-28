
var mymap = L.map('mapid').setView([51.505, -0.09], 3);

L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 18,
    id: 'mapbox.streets',
    accessToken: 'pk.eyJ1Ijoib290aGViYWxsc29vIiwiYSI6ImNqeGZpbGRheDA1c3Uzc284Y20xMGh3eWoifQ.5BZj7OvzM6cn6whqRl37XQ'
}).addTo(mymap);