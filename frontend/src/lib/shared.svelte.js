import { PUBLIC_API_URL } from "$env/static/public";

export let ENV = $state({
  API_URL: PUBLIC_API_URL,
});

export let APP_DATA = $state({
  currentTrack: undefined,
  isPlayingNow: false,
});

export let STORE = $state({
  Artists: [],
});
