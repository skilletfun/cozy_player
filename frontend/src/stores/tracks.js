import {writable} from "svelte/store";


export const tracksURL = "http://127.0.0.1:8000/api/tracks/";

export const tracksStore = writable([]);
export const tracksFilters = writable({});

export async function loadTracks() {
    var url = new URL(tracksURL);
    url.search = new URLSearchParams(tracksFilters);

    const response = await fetch(url.toString());
}

export async function loadTracksByArtist(id) {
    const response = await fetch(tracksURL + `?artist=${id}`);
    const tracks = await response.json();
    return tracks;
}
