import { API } from "$lib/api.js";
import { current } from "$lib/shared.svelte.js";
import { acts } from '@tadashi/svelte-notification';

export let playQueue = [];
export let currentTrackIndex = 0;

export function setupPlayer() {
  try {
    if (!('mediaSession' in navigator)) return;

    document.getElementById("player").addEventListener("ended", onTrackEnds);
    document.getElementById("player").addEventListener("play", () => { current.isPlayingNow = true; });
    document.getElementById("player").addEventListener("pause", () => { current.isPlayingNow = false; });

    navigator.mediaSession.setActionHandler('play', resumePlaying);
    navigator.mediaSession.setActionHandler('pause', pausePlaying);
    navigator.mediaSession.setActionHandler('previoustrack', previousTrack);
    navigator.mediaSession.setActionHandler('nexttrack', nextTrack);
    navigator.mediaSession.setActionHandler('seekto', (details) => {
        seekTo(details.seekTime);
    });
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in setupPlayer: ${e.message}`, lifetime: 5});
  }
}

export function play() {
  try {
    if (playQueue.length === 0) return;

    let player = document.getElementById("player");
    const track = playQueue[currentTrackIndex];
    current.track = track;

    player.src = `${API.Tracks.baseURL}${track.id}/`;
    player.load();
    player.play();

    updateMediaSessionInfo(track);
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in play: ${e.message}`, lifetime: 5});
  }
}

export function updateMediaSessionInfo(track) {
  try {
    document.title = `${track.title} - ${track.artist_name}`;
    navigator.mediaSession.metadata = new MediaMetadata({
      title: track.title,
      artist: track.artist_name,
      artwork: [{src: API.Tracks.getCoverURL(track.id)}]
    });
    navigator.mediaSession.setPositionState({
        duration: track.duration,
        position: 0,
    });
    navigator.mediaSession.playbackState = "playing";
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in updateMediaSessionInfo: ${e.message}`, lifetime: 5});
  }
}

async function onTrackEnds() {
    const track = playQueue[currentTrackIndex];
    await API.Tracks.incrementPlayCount(track);
    nextTrack();
}

export function nextTrack() {
  try {
    currentTrackIndex = currentTrackIndex + 1 === playQueue.length ? 0 : currentTrackIndex + 1;
    play();
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in nextTrack: ${e.message}`, lifetime: 5});
  }
}

export function previousTrack() {
  try {
    currentTrackIndex = Math.max(currentTrackIndex - 1, 0);
    play();
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in previousTrack: ${e.message}`, lifetime: 5});
  }
}

export function pausePlaying() {
  try {
    document.getElementById("player").pause();
    navigator.mediaSession.playbackState = "paused";
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in pausePlaying: ${e.message}`, lifetime: 5});
  }
}

export function resumePlaying() {
  try {
    document.getElementById("player").play();
    navigator.mediaSession.playbackState = "playing";
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in resumePlaying: ${e.message}`, lifetime: 5});
  }
}

export function seekTo(seekTime) {
  try {
    let player = document.getElementById("player");
    seekTime = Math.min(Math.max(seekTime, 0), player.duration);
    player.currentTime = seekTime;
    navigator.mediaSession.setPositionState({
        duration: player.duration,
        position: seekTime,
    });
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in seekTo: ${e.message}`, lifetime: 5});
  }
}

export function playTrack(track) {
  try {
    playQueue = [track];
    currentTrackIndex = 0;
    play();
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in playTrack: ${e.message}`, lifetime: 5});
  }
}

export async function playArtist(artist) {
  try {
    const responseTracks = await API.Tracks.getList({artist: artist.id});
    playQueue = await responseTracks.json();
    currentTrackIndex = 0;
    play();
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in playArtist: ${e.message}`, lifetime: 5});
  }
}

export async function playMainQueue() {
  try {
    const responseTracks = await API.Tracks.getQueue();
    playQueue = await responseTracks.json();
    currentTrackIndex = 0;
    play();
  }
  catch (e) {
    acts.add({mode: 'danger', message: `Error in playMainQueue: ${e.message}`, lifetime: 5});
  }
}