export const ssr = false;
import { API } from "$lib/api.js";

export async function load({ fetch, params }) {
  const responseArtist = await API.Artists.GetById(params.id);
  const artistData = await responseArtist.json();

  const responseTracks = await API.Tracks.GetList({ artistId: params.id });
  const tracksData = await responseTracks.json();

  return { artist: artistData, tracks: tracksData };
}
