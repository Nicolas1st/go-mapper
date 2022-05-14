import {centerCoords} from "./center.js";
import {createAndAttachMap, createMarker} from "./map.js";

let map = createAndAttachMap(centerCoords.longitude, centerCoords.latitude);

// render all the parkings that are already avaiable
const idInput = document.querySelector("#id");
const fetchPromise = fetch("/parkings/all");
const markers = {};
fetchPromise.then(response => {
    return response.json();
}).then(parkings => {
    for (let p of parkings) {
        const marker = createMarker(
            p,
            true,
            "This parking place will be removed",
            () => {
                idInput.value = p.ID;
            },
            () => {
                idInput.value = "";
            },
        );

        marker.addTo(map);
        markers[p.ID] = marker;
    }
})

// handle form submission
const form = document.querySelector("#remove-parking-place-form");
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
    .catch(() => alert("failed to remove the parking"))
    .then((result) => {
        return result.json();
    })
    .then(parking => {
        try {
            markers[parking.ID].remove();
            delete markers[parking.ID]
        } catch {}
    });
});