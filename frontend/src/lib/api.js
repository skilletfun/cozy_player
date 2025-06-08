import { current } from "./shared.svelte";

const API_URL = current.API_URL;

async function get(url, params) {
    const urlParams = new URLSearchParams(params);
    const strUrlParams = urlParams.toString() ? "?" + urlParams.toString() : "";
    return await fetch(`${url}${strUrlParams}`);
}

async function patch(url, id, data) {
    return await fetch(
        `${url}${id}/`,
        {
            method: 'PATCH',
            body: JSON.stringify(data),
            headers: {"Content-Type": "application/json"},
        },
    );
}

export const API = {
    Artists: {
        baseURL: `${API_URL}/artists/`,
        getList: (filters) => get(API.Artists.baseURL, filters),
        getById: (id) => get(`${API.Artists.baseURL}${id}/`),
        getCover: (id) => get(`${API.Artists.baseURL}cover/${id}/`),
        getCoverURL: (id) => `${API.Artists.baseURL}cover/${id}/`,
    },
    Library: {
        baseURL: `${API_URL}/library/`,
        rescan: () => get(`${API.Library.baseURL}rescan/`),
        getStatistic: () => get(`${API.Library.baseURL}stats/`),
    },
    Tracks: {
        baseURL: `${API_URL}/tracks/`,
        getList: (filters) => get(API.Tracks.baseURL, filters),
        getById: (id) => get(`${API.Tracks.baseURL}${id}/`),
        getQueue: () => get(`${API.Tracks.baseURL}queue/`),
        incrementPlayCount: (track) => patch(API.Tracks.baseURL, track.id, {play_count: track.play_count + 1}),
        getCoverURL: (id) => `${API.Tracks.baseURL}cover/${id}/`,
    }
}