<script>
    import NavigationButton from "$lib/components/NavigationButton.svelte";
    import { onMount } from "svelte";
    import { pausePlaying, resumePlaying, nextTrack, previousTrack, seekTo } from "$lib/player.js";
    let { children } = $props();

    onMount(() => {
        if (!('mediaSession' in navigator)) return;

        navigator.mediaSession.setActionHandler('play', resumePlaying);
        navigator.mediaSession.setActionHandler('pause', pausePlaying);
        navigator.mediaSession.setActionHandler('previoustrack', previousTrack);
        navigator.mediaSession.setActionHandler('nexttrack', nextTrack);
        navigator.mediaSession.setActionHandler('seekto', (details) => {
            seekTo(details.seekTime);
        });

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