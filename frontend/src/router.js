import { writable, readable } from "svelte/store";
import ArtistList from "./lib/ArtistList.svelte";
import PlaylistList from "./lib/PlaylistList.svelte";
import HomeView from "./lib/HomeView.svelte";
import TrackList from "./lib/TrackList.svelte";
import ArtistView from "./lib/ArtistView.svelte";


export const router = readable({
    Artists: ArtistList,
    Home: HomeView,
    Playlists: PlaylistList,
    Tracks: TrackList,
    ArtistView: ArtistView,
});

export const currentRoute = writable("Home");