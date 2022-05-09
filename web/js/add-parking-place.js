import {centerCoords} from "./center.js";
import {createAndAttachMap, displayMarker} from "./map.js";

// init map
let map = createAndAttachMap(centerCoords.longitude, centerCoords.latitude);

// render all the parkings that are already avaiable
const fetchPromise = fetch("/parkings/all");
fetchPromise.then(response => {
    return response.json();
}).then(parkings => {
    for (let p of parkings) {
        const marker = displayMarker(p);
        marker.addTo(map);
    }
})

// fields in the form
const latitudeField = document.querySelector("#latitude");
const longitudeField = document.querySelector("#longitude");
const addressField = document.querySelector("#address");
const numberOfSlotsFields = document.querySelector("#numberOfSlots");

// adding new marker on the map
let marker = undefined; // to remove the previous marker
map.on("click", function (e) {
    console.log(e.target);
    latitudeField.value = e.lngLat.lat;
    longitudeField.value = e.lngLat.lng;

    try {
        marker.remove();
    } catch {
    }

    marker = displayMarker({Latitude: e.lngLat.lat, Longitude: e.lngLat.lng, NumberOfSlots: numberOfSlotsFields.value});
    marker.addTo(map);
});

// submit event
const form = document.querySelector(".add-parking-place");
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    fetch(`/parkings/`, {
        method: "POST",
        header: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Address: addressField.value,
            Latitude: Number(latitudeField.value),
            Longitude: Number(longitudeField.value),
            NumberOfSlots: Number(numberOfSlotsFields.value),
        }),
    })
    .catch(() => alert("failed to add the parking"))
    .then((result) => {
        addressField.value = "";
        latitudeField.value = "";
        longitudeField.value = "";
        numberOfSlotsFields.value = "";

        return result.json();
    })
    .then(parking => {
        displayMarker(parking).addTo(map);
    });
});