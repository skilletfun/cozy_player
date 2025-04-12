<script>
  import { API } from "$lib/api.js";
  let { artist } = $props();
</script>

<div>
  <img src="{API.ArtistCover}/{artist.id}/" alt="cover">
  <a href="/artists/{artist.id}">{artist.name}</a>
</div>

<style>
  div {
    width: 150px;
    height: 200px;
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
    max-width: 75%;
    margin-top: 15px;
  }

  a {
    align-self: center;
    text-align: center;
    margin-inline: 5px;
    font-size: 11pt;
  }
</style>