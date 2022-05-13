import {createAndAttachMap, displayMarker} from "./map.js";
import {centerCoords} from "./center.js";

let map = createAndAttachMap(centerCoords.longitude, centerCoords.latitude);

// render all the parkings that are already avaiable
const inputAddress = document.querySelector("#id");
const fetchPromise = fetch("/parkings/all");
fetchPromise.then(response => {
    return response.json();
}).then(parkings => {
    for (let p of parkings) {
        const marker = displayMarker(p, true, "This library has been chosen", p.ID, inputAddress);
        marker.addTo(map);
    }
});