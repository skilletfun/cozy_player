<script>
    import { playTrack } from "$lib/player.js";
    import { current } from "$lib/shared.svelte.js";

    let { index, track } = $props();

    function isCurrentTrack() {
        return current.track && current.track.id === track.id && current.track.artist === track.artist;
    }
</script>

<div>
    <p>{`${index + 1}`.padStart(5, ' ')}</p>
    <button onclick={() => playTrack(track)} class={isCurrentTrack() ? 'current' : '' }>{track.title}</button>
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

    button {
        border-width: 0;
        font-family: monospace;
        font-size: 16px;
        background-color: transparent;
    }
    button.current {
        color: orangered;
    }
</style>