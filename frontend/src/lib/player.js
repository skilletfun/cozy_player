import {API} from "$lib/api.js";

export const player = typeof document !== 'undefined' ? document.getElementById("player") : undefined;

export let playQueue = [];

function fetchPlayQueue() {

}

export function play() {
    if (playQueue.length === 0) {
        fetchPlayQueue();
    }

    const track = playQueue.shift();

    player.src = `${API.Tracks}${track.id}/`;
    player.load();
    player.play();

    document.title = `${track.title} - ${track.artist_name}`;
}

export function playTrack(track) {
    playQueue = [track];
    play();
}

export async function playArtist(artist) {
    const responseTracks = await fetch(API.Tracks + `?artist=${artist.id}`);
    playQueue = await responseTracks.json();
    play();
}

export function playPlaylist(playlist) {
    // playQueue = [track];
    // play();
}