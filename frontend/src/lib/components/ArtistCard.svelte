<a class="hover" href="/artists/{artist.id}">
  <img src={API.Artists.GetCoverURL(artist.id)} alt="cover" class="cover"/>
  <button aria-label="play" class="play-btn row center" onclick={(e) => {e.preventDefault(); playArtist(artist);}}>
    <i class="fas fa-play flex center"></i>
  </button>
  <p>{artist.name}</p>
</a>

<script>
  import { API } from "$lib/services/api.js";
  import { playArtist } from "$lib/services/player.js";
  
  let { artist } = $props();
</script>

<style>
  a {
    width: 150px;
    height: 195px;
    background-color: transparent;
    border-radius: 10px;
    max-width: 550px;

    display: flex;
    flex-direction: column;
    align-content: space-around;
  }

  img {
    border-radius: 10px;
    width: 115px;
    height: 115px;
    margin-top: 15px;
  }

  p {
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
  a:hover > button.play-btn {
    opacity: 1;
  }
  .play-btn:active {
    background-color: #373737;
  }
</style>
