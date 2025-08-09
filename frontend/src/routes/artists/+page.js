export const ssr = false;
import { API } from "$lib/api.js";
import { APP_DATA } from "$lib/shared.svelte.js";

export async function load({ fetch, params }) {
  const response = await API.Artists.GetList();
  const data = await response.json();
  APP_DATA.artists = data;
  return { artists: data };
}
