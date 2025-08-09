<script>
    import { playTrack } from "$lib/player.js";
    import { APP_DATA } from "$lib/shared.svelte.js";

    let { index, track } = $props();

    function isCurrentTrack() {
        return (
            APP_DATA.currentTrack &&
            APP_DATA.currentTrack.id === track.id &&
            APP_DATA.currentTrack.artist === track.artist
        );
    }
</script>

<div>
    <p>{`${index + 1}`.padStart(5, " ")}</p>
    <button
        onclick={() => playTrack(track)}
        class={isCurrentTrack() ? "current" : ""}>{track.title}</button
    >
    <p class="album">{track.album}</p>
</div>

<style>
    div {
        display: flex;
        margin-left: -10px;
        align-items: center;
    }
    p {
        margin-right: 20px;
        color: #777777;
        white-space: break-spaces;
    }
    p.album {
        font-size: 14px;
        margin-left: 20px;
        line-height: 1;
    }

    button {
        border-width: 0;
        font-family: monospace;
        font-size: 16px;
        background-color: transparent;
        width: 400px;
        text-align: left;
    }
    button.current {
        color: orangered;
    }
</style>
