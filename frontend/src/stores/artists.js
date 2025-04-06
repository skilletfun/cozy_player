import { writable } from "svelte/store";


export const artistsURL = "http://127.0.0.1:8000/api/artists/";

export const artistsStore = writable([]);

export const currentArtist = writable(0);
export const currentArtistInfo = writable();

export async function loadArtists() {
    const response = await fetch(artistsURL);
    const artists = await response.json()
    artistsStore.set(artists);
}

export async function getArtistInfo(id) {
    const response = await fetch(artistsURL + `${id}/`);
    return await response.json();
}
