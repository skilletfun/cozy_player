import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
	const response = await fetch(API.LibraryStats);
    return await response.json();
}
