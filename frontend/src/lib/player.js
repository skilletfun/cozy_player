import { API } from "$lib/api.js";
import { current } from "$lib/shared.svelte.js";

export let playQueue = [];
export let currentTrackIndex = 0;

export function setupPlayer() {
    if (!('mediaSession' in navigator)) return;

    document.getElementById("player").addEventListener("ended", onTrackEnds);

    navigator.mediaSession.setActionHandler('play', resumePlaying);
    navigator.mediaSession.setActionHandler('pause', pausePlaying);
    navigator.mediaSession.setActionHandler('previoustrack', previousTrack);
    navigator.mediaSession.setActionHandler('nexttrack', nextTrack);
    navigator.mediaSession.setActionHandler('seekto', (details) => {
        seekTo(details.seekTime);
    });
}

export function play() {
    if (playQueue.length === 0) return;

    let player = document.getElementById("player");
    const track = playQueue[currentTrackIndex];
    current.track = track;

    player.src = `${API.Tracks}${track.id}/`;
    player.load();
    player.play();

    updateMediaSessionInfo(track);
}

export function updateMediaSessionInfo(track) {
    document.title = `${track.title} - ${track.artist_name}`;
    navigator.mediaSession.metadata = new MediaMetadata({
      title: track.title,
      artist: track.artist_name,
      artwork: [{src: `${API.Tracks}cover/${track.id}/`}]
    });
    navigator.mediaSession.setPositionState({
        duration: track.duration,
        position: 0,
    });
    navigator.mediaSession.playbackState = "playing";
}

function onTrackEnds() {
    const track = playQueue[currentTrackIndex];
    fetch(
    `${API.Tracks}${track.id}/`,
    {
            method: 'PATCH',
            body: JSON.stringify({play_count: track.play_count + 1}),
            headers: {"Content-Type": "application/json"},
        },
    );
    nextTrack();
}

export function nextTrack() {
    currentTrackIndex = currentTrackIndex + 1 === playQueue.length ? 0 : currentTrackIndex + 1;
    play();
}

export function previousTrack() {
    currentTrackIndex = Math.max(currentTrackIndex - 1, 0);
    play();
}

export function pausePlaying() {
    document.getElementById("player").pause();
    navigator.mediaSession.playbackState = "paused";
}

export function resumePlaying() {
    document.getElementById("player").play();
    navigator.mediaSession.playbackState = "playing";
}

export function seekTo(seekTime) {
    let player = document.getElementById("player");
    seekTime = Math.min(Math.max(seekTime, 0), player.duration);
    player.currentTime = seekTime;
    navigator.mediaSession.setPositionState({
        duration: player.duration,
        position: seekTime,
    });
}

export function playTrack(track) {
    playQueue = [track];
    currentTrackIndex = 0;
    play();
}

export async function playArtist(artist) {
    const responseTracks = await fetch(API.Tracks + `?artist=${artist.id}`);
    playQueue = await responseTracks.json();
    currentTrackIndex = 0;
    play();
}
