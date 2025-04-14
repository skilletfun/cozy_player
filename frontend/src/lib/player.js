import {API} from "$lib/api.js";

export let playQueue = [];
export let currentTrackIndex = 0;

export function play() {
    if (playQueue.length === 0) {
        return;
    }

    let player = document.getElementById("player");
    const track = playQueue[currentTrackIndex];

    player.src = `${API.Tracks}${track.id}/`;
    player.load();
    player.play();

    document.title = `${track.title} - ${track.artist_name}`;
    navigator.mediaSession.metadata = new MediaMetadata({
      title: track.title,
      artist: track.artist_name,
      artwork: [{src: `${API.Tracks}cover/${track.id}/`}]
    });
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
