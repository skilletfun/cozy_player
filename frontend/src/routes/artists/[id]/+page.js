import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
	const responseArtist = await API.Artists.getById(params.id);
    const artistData = await responseArtist.json();

    const responseTracks = await API.Tracks.getList({artist: params.id});
    const tracksData = await responseTracks.json();

    return {artist: artistData, tracks: tracksData};
}