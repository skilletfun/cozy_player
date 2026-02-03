import { API } from "$lib/services/api.js";
import { ENV, APP_DATA, STORE } from "$lib/shared.svelte.js";
import { acts } from "@tadashi/svelte-notification";

export function setupPlayer() {
  try {
    if (!("mediaSession" in navigator)) return;

    const player = document.getElementById("player");

    player.addEventListener("error", () => {
      acts.add({ mode: "danger", message: `Cannot play track`, lifetime: 5 });
    });
    player.addEventListener("ended", nextTrack);
    player.addEventListener("play", () => {
      APP_DATA.isPlayingNow = true;
    });
    player.addEventListener("pause", () => {
      APP_DATA.isPlayingNow = false;
    });

    navigator.mediaSession.setActionHandler("play", resumePlaying);
    navigator.mediaSession.setActionHandler("pause", pausePlaying);
    navigator.mediaSession.setActionHandler("previoustrack", previousTrack);
    navigator.mediaSession.setActionHandler("nexttrack", nextTrack);
    navigator.mediaSession.setActionHandler("seekto", (details) => {
      seekTo(details.seekTime);
    });
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in setupPlayer: ${e.message}`,
      lifetime: 5,
    });
  }
}

export function removeHandlers() {
  if (!("mediaSession" in navigator)) return;

  navigator.mediaSession.setActionHandler("play", null);
  navigator.mediaSession.setActionHandler("pause", null);
  navigator.mediaSession.setActionHandler("previoustrack", null);
  navigator.mediaSession.setActionHandler("nexttrack", null);
  navigator.mediaSession.setActionHandler("seekto", null);
}

export function play() {
  try {
    let player = document.getElementById("player");

    player.src = `${ENV.API_URL}/track/${APP_DATA.currentTrack.id}`;
    player.load();
    player.play();

    updateMediaSessionInfo();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in play: ${e.message}`,
      lifetime: 5,
    });
  }
}

export function updateMediaSessionInfo() {
  try {
    const artistId = APP_DATA.currentTrack.artistId;
    const artistName = STORE.Artists.find((e) => e.id == artistId).name;
    document.title = `${APP_DATA.currentTrack.title} - ${artistName}`;

    navigator.mediaSession.metadata = new MediaMetadata({
      title: APP_DATA.currentTrack.title,
      artist: artistName,
      artwork: [{ src: API.Tracks.GetCoverURL(APP_DATA.currentTrack.id) }],
    });

    navigator.mediaSession.setPositionState({
      duration: APP_DATA.currentTrack.duration,
      position: 0,
    });
    navigator.mediaSession.playbackState = "playing";
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in updateMediaSessionInfo: ${e.message}`,
      lifetime: 5,
    });
  }
}

export async function nextTrack() {
  try {
    APP_DATA.currentTrack = await (await API.Queue.Next()).json();
    play();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in nextTrack: ${e.message}`,
      lifetime: 5,
    });
  }
}

export async function previousTrack() {
  try {
    APP_DATA.currentTrack = await (await API.Queue.Prev()).json();
    play();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in previousTrack: ${e.message}`,
      lifetime: 5,
    });
  }
}

export function pausePlaying() {
  try {
    document.getElementById("player").pause();
    navigator.mediaSession.playbackState = "paused";
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in pausePlaying: ${e.message}`,
      lifetime: 5,
    });
  }
}

export function resumePlaying() {
  try {
    document.getElementById("player").play();
    navigator.mediaSession.playbackState = "playing";
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in resumePlaying: ${e.message}`,
      lifetime: 5,
    });
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
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in seekTo: ${e.message}`,
      lifetime: 5,
    });
  }
}

export async function playTrack(track) {
  try {
    await loadArtists();
    await API.Queue.GenerateByTrack(track.id);
    await nextTrack();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in playTrack: ${e.message}`,
      lifetime: 5,
    });
  }
}

export async function playArtist(artist) {
  try {
    await loadArtists();
    await API.Queue.GenerateByArtist(artist.id);
    await nextTrack();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in playArtist: ${e.message}`,
      lifetime: 5,
    });
  }
}

export async function playMainQueue() {
  try {
    await loadArtists();
    await API.Queue.Generate();
    await nextTrack();
  } catch (e) {
    acts.add({
      mode: "danger",
      message: `Error in playMainQueue: ${e.message}`,
      lifetime: 5,
    });
  }
}

async function loadArtists() {
  if (Object.values(STORE.Artists).length == 0) {
    const response = await API.Artists.GetList();
    const data = await response.json();
    STORE.Artists = data;
  }
}
