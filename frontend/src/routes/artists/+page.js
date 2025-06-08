export const ssr = false;
import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
    const response = await API.Artists.getList();
    const data = await response.json();
    return {artists: data};
}
