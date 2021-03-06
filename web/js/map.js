export function createAndAttachMap(longitude, latitude) {
    var map = new maplibregl.Map({
        container: "map",
        style: "https://api.maptiler.com/maps/streets/style.json?key=get_your_own_OpIi9ZULNHzrESv6T2vL",
        center: [longitude, latitude],
        zoom: 11,
    });
    map.scrollZoom.disable();

    return map;
}

export function createMarker(
    parking,
    addPopup,
    popupMessage,
    callbackOnOpen,
    callbackOnClose,
) {
    // create DOM element for the marker
    let el = document.createElement("div");
    if (parking.NumberOfSlots !== undefined) {
        el.innerText = parking.NumberOfSlots;
    }
    el.dataset.id = parking.ID;
    el.classList.add("marker");

    // create the marker
    let marker = new maplibregl.Marker(el)
        .setLngLat([parking.Longitude, parking.Latitude])

    // add popup
    if (addPopup) {
        var popup = new maplibregl.Popup({ offset: 25 }).setText(popupMessage);
        popup.on("open", () => {
            callbackOnOpen();
        });

        popup.on("close", () => {
            callbackOnClose();
        });

        marker.setPopup(popup)
    }

    return marker;
}