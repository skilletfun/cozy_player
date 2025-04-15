<script>
    import NavigationButton from "$lib/components/NavigationButton.svelte";
    import { onMount } from "svelte";
    import { setupPlayer } from "$lib/player.js";
    let { children } = $props();

    onMount(() => {
        setupPlayer();

        return () => {
          // Cleanup
          if ('mediaSession' in navigator) {
            navigator.mediaSession.metadata = null;
          }
        };
    });
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