import { ENV } from "./shared.svelte";

async function get(url, params) {
  const urlParams = new URLSearchParams(params);
  const strUrlParams = urlParams.toString() ? "?" + urlParams.toString() : "";
  return await fetch(`${ENV.API_URL}${url}${strUrlParams}`);
}

export const API = {
  Artists: {
    GetList: (filters) => get(`/artists`, filters),
    GetById: (id) => get(`/artist/${id}`),
    GetCover: (id) => get(`/artist/cover/${id}`),
    GetCoverURL: (id) => `${ENV.API_URL}/artist/cover/${id}`,
  },
  Library: {
    Rescan: () => get(`/library/rescan`),
  },
  Tracks: {
    GetList: (filters) => get(`/tracks`, filters),
    GetById: (id) => get(`/track/info${id}`),
    GetCoverURL: (id) => `${ENV.API_URL}/track/cover/${id}`,
  },
  Queue: {
    Generate: () => get(`/queue`),
    GenerateByArtist: (artistId) => get(`/queue?artistId=${artistId}`),
    GenerateByTrack: (trackId) => get(`/queue?trackId=${trackId}`),
    Next: () => get(`/queue/next`),
    Prev: () => get(`/queue/prev`),
  },
};
