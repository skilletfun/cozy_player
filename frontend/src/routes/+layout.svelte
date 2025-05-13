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
            if ("mediaSession" in navigator) {
                navigator.mediaSession.metadata = null;
            }
        };
    });

    async function rescanLibrary() {
        try {
            await API.Library.rescan();
        } catch (e) {
            console.log(e);
            throw e;
        }
    }
</script>

<main>
    <div class="navigation">
        <button class="btn" onclick={playMainQueue}>Play</button>
        <div class="spacer"></div>
        <!-- <button class="btn" onclick={rescanLibrary}>Rescan</button> -->
        <NavigationButton title="Home" icon="home" href="/" />
        <NavigationButton title="Artists" icon="user-music" href="/artists" />
        <NavigationButton
            title="Playlists"
            icon="list-music"
            href="/playlists"
        />
        <NavigationButton title="Tracks" icon="list-music" href="/tracks" />
    </div>
    <div class="content">
        {@render children()}
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
    }
    div.navigation {
        margin-top: 10px;
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
    div.content {
        margin-left: 100px;
        margin-right: 100px;
    }
    .spacer {
        width: 100px;
    }
    .btn {
        width: 120px;
        height: 40px;
        font-size: 16px;
        border-radius: 10px;
        background-color: transparent;
        align-self: center;
        align-items: center;
        justify-content: center;
        border-radius: 150px;
        display: flex;
        margin-left: -220px;
    }
    .btn:hover {
        background-color: #3c3c3c;
    }
    .btn::before {
        margin-right: 10px;
        scale: 1.25;
        content: "▶";
    }
</style>
