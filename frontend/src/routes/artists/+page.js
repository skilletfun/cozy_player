import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
	const response = await fetch(API.Artists);
    const data = await response.json();
    return {artists: data};
}
