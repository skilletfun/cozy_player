import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
	const responseArtist = await fetch(API.Artists + `${params.id}/`);
    const artistData = await responseArtist.json();

    const responseTracks = await fetch(API.Tracks + `?artist=${params.id}`);
    const tracksData = await responseTracks.json();

    return {artist: artistData, tracks: tracksData};
}