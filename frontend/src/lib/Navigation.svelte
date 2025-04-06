<script>
  import NavButton from "./NavButton.svelte";
</script>

<div>
  <h2>My Library</h2>
  <NavButton title="Home" icon="home"/>
  <NavButton title="Artists" icon="user-music"/>
  <NavButton title="Playlists" icon="list-music"/>
  <NavButton title="Tracks" icon="list-music"/>
  <NavButton title="Statistic" icon="analytics"/>
</div>

<style>
  div {
    margin-top: 50px;
    margin-left: 150px;
  }
</style>