<div class="column center">
  {#if APP_DATA.currentTrack && APP_DATA.currentTrack.id}
    <div class="row center">
      <img
        class="cover"
        src={API.Tracks.GetCoverURL(APP_DATA.currentTrack.id)}
        alt="Cover"
      />
      <div class="player-info column center">
        <p class="track-title">{APP_DATA.currentTrack.title}</p>
        <a
          class="track-group"
          href="/artists/{APP_DATA.currentTrack.artistId}"
        >
          {getArtistName(APP_DATA.currentTrack.artistId)}
        </a>
        <div class="player-controls flex">
          <PlayerButton icon="backward" onclick={previousTrack} />
          <PlayerButton
            icon={APP_DATA.isPlayingNow ? "pause" : "play"}
            onclick={getFunction}
          />
          <PlayerButton icon="forward" onclick={nextTrack} />
        </div>
      </div>
    </div>
  {:else}
    <div style="scale: 2;">
      <ActionButton title="Play" icon="play" onclick={playMainQueue} class="play"/>
    </div>
  {/if}
</div>

<script>
  import { APP_DATA, STORE } from "$lib/shared.svelte.js";
  import { API } from "$lib/services/api";
  import {
    playMainQueue,
    previousTrack,
    nextTrack,
    pausePlaying,
    resumePlaying,
  } from "$lib/services/player";
  import PlayerButton from "./PlayerButton.svelte";
  import ActionButton from "./ActionButton.svelte";

  function getFunction() {
    return APP_DATA.isPlayingNow ? pausePlaying() : resumePlaying();
  }

  function getArtistName(artistId) {
    return STORE.Artists.find((e) => e.id == artistId).name;
  }
</script>

<style>
  div.column {
    height: 85vh;
  }

  .player-info {
    width: 600px;
  }

  .player-controls {
    flex-direction: row;
    justify-content: space-between;
    width: 50%;
    margin-top: 25px;
  }

  p.track-title {
    font-size: 30px;
    text-align: center;
    width: 500px;
  }

  a.track-group {
    font-size: 20px;
    margin-top: -10px;
  }

  img {
    border-radius: 35px;
    height: 300px;
    width: 300px;
    margin-top: 15px;
  }
</style>
