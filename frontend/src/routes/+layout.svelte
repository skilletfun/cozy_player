<script>
    import NavigationButton from "$lib/components/NavigationButton.svelte";
    import { onMount } from "svelte";
    import { playMainQueue, setupPlayer } from "$lib/player.js";
    import { API } from "$lib/api.js";
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

    async function rescanLibrary() {
        try {
          await API.Library.rescan();
        }
        catch (e) {
          console.log(e);
          throw e;
        }

    }

</script>

<main>
    <div class="navigation">
        <div style="display: flex;flex-direction: column">
            <button class="btn" onclick={playMainQueue}>Play</button>
            <button class="btn" onclick={rescanLibrary}>Rescan</button>
        </div>
        <h2>My Library</h2>
        <NavigationButton title="Home" icon="home" href="/"/>
        <NavigationButton title="Artists" icon="user-music" href="/artists"/>
<!--        <NavigationButton title="Playlists" icon="list-music" href="/playlists"/>-->
<!--        <NavigationButton title="Tracks" icon="list-music"/>-->
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
  .btn {
      width: 125px;
      height: 40px;
      font-size: 16px;
      border-radius: 10px;
      background-color: transparent;
      border: 1px solid grey;
      margin-bottom: 15px;
  }
</style>