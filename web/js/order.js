import {createAndAttachMap, createMarker} from "./map.js";
import {centerCoords} from "./center.js";

let map = createAndAttachMap(centerCoords.longitude, centerCoords.latitude);

// form fields
const addressField  = document.querySelector("#address");
const parkingIDField = document.querySelector("#parking-id")
const startHourField  = document.querySelector("#start-hour");
const endHourField  = document.querySelector("#end-hour");


// render all the parkings that are already avaiable
const fetchPromise = fetch("/parkings/all");
fetchPromise.then(response => {
    return response.json();
}).then(parkings => {
    for (let p of parkings) {
        const marker = createMarker(p, true, "This parking has been chosen",
            () => {
                parkingIDField.value = p.ID;
                addressField.value = p.Address;
            },
            () => {}
        );
        marker.addTo(map);
    }
});


// submit event
const form = document.querySelector("#make-order-form");
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    if (Number(startHourField.value) > Number(endHourField.value)) {
        alert("The start time can not be after the end time")
        return;
    }

    form.submit();
});