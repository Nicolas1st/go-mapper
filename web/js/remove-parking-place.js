import {centerCoords} from "./center.js";
import {createAndAttachMap, displayMarker} from "./map.js";

let map = createAndAttachMap(centerCoords.longitude, centerCoords.latitude);

// render all the parkings that are already avaiable
const idInput = document.querySelector("#id");
const fetchPromise = fetch("/parkings/all");
const markers = {};
fetchPromise.then(response => {
    return response.json();
}).then(parkings => {
    for (let p of parkings) {
        const marker = displayMarker(p, true, "This parking place will be removed", p.ID, idInput);
        marker.addTo(map);
        markers[p.ID] = marker;
    }
})

const form = document.querySelector(".remove-parking-place-form");
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    fetch(`/parkings/`, {
        method: "DELETE",
        header: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            ID: Number(idInput.value),
        }),
    })
    .catch(() => alert("failed to add the parking"))
    .then((result) => {
        return result.json();
    })
    .then(parking => {
        idInput.value = "";
        markers[parking.ID].remove();
        delete markers[parking.ID]
    });
});