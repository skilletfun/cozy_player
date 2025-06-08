<script>
  import { API } from "$lib/api.js";
  import { playArtist } from "$lib/player.js";
  let { artist } = $props();
</script>

<div>
  <img src="{API.Artists.getCoverURL(artist.id)}" alt="cover">
  <button aria-label="play" class="play-btn" onclick={() => playArtist(artist)}><i class="fas fa-play"></i></button>
  <a href="/artists/{artist.id}">{artist.name}</a>
</div>

<style>
  div {
    width: 150px;
    height: 195px;
    background-color: transparent;
    border-radius: 10px;
    margin-top: 15px;
    margin-right: 15px;
    max-width: 550px;

    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    align-content: space-around;
  }
  div:hover {
    background-color: #3c3c3c;
    cursor: pointer;
  }

  img {
    align-self: center;
    border-radius: 10px;
    width: 115px;
    height: 115px;
    object-fit: cover;
    margin-top: 15px;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
  }

  a {
    text-align: center;
    margin-inline: 5px;
    font-size: 11pt;
  }
  .play-btn {
    height: 35px;
    width: 60px;
    border-radius: 18px;
    border-width: 0;
    cursor: pointer;
    margin-top: -45px;
    margin-bottom: 10px;
    align-self: center;
    opacity: 0;
    transition: 0.3s;
  }
  div:hover > button.play-btn {
    opacity: 1;
  }
  .play-btn:active {
    background-color: #373737;
  }
</style>