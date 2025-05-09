import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
	const response = await API.Library.getStatistic();
    return await response.json();
}
