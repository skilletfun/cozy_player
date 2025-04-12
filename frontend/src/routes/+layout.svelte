<script>
    import NavigationButton from "$lib/NavigationButton.svelte";
    let { children } = $props();
</script>

<main>
    <div class="navigation">
        <h2>My Library</h2>
        <NavigationButton title="Home" icon="home" href="/"/>
        <NavigationButton title="Artists" icon="user-music" href="/artists"/>
<!--        <NavigationButton title="Playlists" icon="list-music" href="/playlists"/>-->
<!--        <NavigationButton title="Tracks" icon="list-music"/>-->
<!--        <NavigationButton title="Statistic" icon="analytics"/>-->
    </div>
    {@render children()}
</main>

<style>
  main {
    display: flex;
  }
  div.navigation {
    margin-top: 50px;
    margin-left: 150px;
  }
</style>