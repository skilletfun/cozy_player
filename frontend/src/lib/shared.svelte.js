import { PUBLIC_API_URL } from '$env/static/public';
export let current = $state({track: undefined, isPlayingNow: false, API_URL: PUBLIC_API_URL});
