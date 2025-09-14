<div class="artist-view">
  <div class="artist-header flex">
    <img src={API.Artists.GetCoverURL(data.artist.id)} alt="cover" class="cover"/>
    <div class="column">
      <h1>{data.artist.name}</h1>
      <h2>{artistInfoHeader}</h2>
      <button class="play-btn" onclick={() => playArtist(data.artist)}>
        <i class="fas fa-play"></i> Play
      </button>
    </div>
  </div>
  <TrackList data={data.tracks} />
</div>

<script>
  import { secondsToHumanString } from "$lib/services/utils.js";
  import { API } from "$lib/services/api.js";
  import TrackList from "$lib/components/TrackList.svelte";
  import { playArtist } from "$lib/services/player.js";

  let { data } = $props();
  let artistInfoHeader = `${data.artist.tracksCount} Tracks • ${secondsToHumanString(data.artist.duration)}`;
</script>

<style>
  .artist-view {
    align-self: center;
    margin-top: 50px;
  }
  .artist-header {
    margin-bottom: 40px;
  }
  .column {
    justify-content: center;
  }
  h1 {
    margin-top: 0;
    margin-bottom: 0;
  }
  img {
    width: 200px;
    height: 200px;
    border-radius: 25px;
    margin-right: 30px;
  }
  .play-btn {
    height: 35px;
    width: 90px;
    border-radius: 18px;
  }
</style>
