<div class="hover">
  <img src={API.Artists.GetCoverURL(artist.id)} alt="cover" class="cover"/>
  <button aria-label="play" class="play-btn" onclick={() => playArtist(artist)} >
    <i class="fas fa-play"></i>
  </button>
  <a href="/artists/{artist.id}">{artist.name}</a>
</div>

<script>
  import { API } from "$lib/api.js";
  import { playArtist } from "$lib/player.js";
  
  let { artist } = $props();
</script>

<style>
  div {
    width: 150px;
    height: 195px;
    background-color: transparent;
    border-radius: 10px;
    max-width: 550px;

    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    align-content: space-around;
  }

  img {
    border-radius: 10px;
    width: 115px;
    height: 115px;
    margin-top: 15px;
  }

  a {
    text-align: center;
    margin-inline: 5px;
    margin-top: 10px;
    font-size: 11pt;
  }
  .play-btn {
    height: 35px;
    width: 60px;
    border-radius: 18px;
    margin-top: -45px;
    margin-bottom: 10px;
    align-self: center;
    opacity: 0;
  }
  div:hover > button.play-btn {
    opacity: 1;
  }
  .play-btn:active {
    background-color: #373737;
  }
</style>
