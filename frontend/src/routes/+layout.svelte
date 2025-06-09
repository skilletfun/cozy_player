<script>
    import {Notifications, acts} from '@tadashi/svelte-notification'
    import NavigationButton from "$lib/components/NavigationButton.svelte";
    import { onMount } from "svelte";
    import { playMainQueue, setupPlayer } from "$lib/player.js";
    import { API } from "$lib/api.js";
    import { rescan } from "$lib/library.js";
    let { children } = $props();

    onMount(() => {
        setupPlayer();
        return () => {
            if ("mediaSession" in navigator) {
                navigator.mediaSession.metadata = null;
            }
        };
    });
</script>

<main>
    <div class="navigation">
        <NavigationButton title="Home" icon="home" href="/" />
        <NavigationButton title="Artists" icon="user-music" href="/artists" />
        <div class="pipe">|</div>
        <button class="btn play" onclick={playMainQueue}>Play</button>
        <button class="btn rescan" onclick={rescan}>Rescan</button>
    </div>
    <div class="content">
        {@render children()}
    </div>
    
    <Notifications />
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
    .pipe {
        align-self: center;
        scale: 1.5;
        margin-left: 15px;
        margin-right: 15px;
    }
    .btn {
        height: 40px;
        font-size: 16px;
        border-radius: 10px;
        background-color: transparent;
        color: #E08B5D;
        align-self: center;
        align-items: center;
        justify-content: center;
        border-radius: 150px;
        display: flex;
        padding-right: 15px;
        padding-left: 15px;
    }
    .btn:hover {
        background-color: #3c3c3c;
    }
    .btn.play::before {
        margin-right: 10px;
        scale: 1.25;
        content: "▶";
    }
    .btn.rescan::before {
        margin-right: 10px;
        scale: 1.25;
        content: "⟳"
    }
</style>
