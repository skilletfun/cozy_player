import { ENV } from "$lib/shared.svelte";
import { get, post } from "$lib/services/client";

export const API = {
  Artists: {
    GetList: (filters) => get(`/artists`, filters),
    GetById: (id) => get(`/artist/${id}`),
    GetCover: (id) => get(`/artist/cover/${id}`),
    GetCoverURL: (id) => `${ENV.API_URL}/artist/cover/${id}`,
  },
  Library: {
    Rescan: () => post(`/library/rescan`),
  },
  Tracks: {
    GetList: (filters) => get(`/tracks`, filters),
    GetById: (id) => get(`/track/info${id}`),
    GetCoverURL: (id) => `${ENV.API_URL}/track/cover/${id}`,
  },
  Queue: {
    Generate: () => post(`/queue`, {}),
    GenerateByArtist: (artistId) => post("/queue", {artistId: artistId}),
    GenerateByTrack: (trackId) => post("/queue", {trackId: trackId}),
    Next: () => get(`/queue/next`),
    Prev: () => get(`/queue/prev`),
  },
};
