<div class="artist-view">
  <div class="artist-header flex mb-4">
    <img src={API.Artists.GetCoverURL(data.artist.id)} alt="cover" class="cover"/>
    <div class="artist-info">
      <h1 class="mt-0 mb-0">{data.artist.name}</h1>
      <h2>{artistInfoHeader}</h2>
      <ActionButton title="Play" icon="play" onclick={() => playArtist(data.artist)} class="play"/>
    </div>
  </div>
  <TrackList data={data.tracks} />
</div>

<script>
  import { secondsToHumanString } from "$lib/services/utils.js";
  import { API } from "$lib/services/api.js";
  import TrackList from "$lib/components/TrackList.svelte";
  import { playArtist } from "$lib/services/player.js";
  import ActionButton from "$lib/components/ActionButton.svelte";

  let { data } = $props();
  let artistInfoHeader = `${data.artist.tracksCount} Tracks • ${secondsToHumanString(data.artist.duration)}`;
</script>

<style>
  .artist-view {
    align-self: center;
    margin-top: 50px;
  }

  .artist-info {
    justify-content: center;
    display: flex;
    flex-direction: column;
    align-items: baseline;
  }

  img {
    width: 200px;
    height: 200px;
    border-radius: 25px;
    margin-right: 30px;
  }
</style>
