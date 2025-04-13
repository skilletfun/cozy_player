<script>
    import {playTrack} from "$lib/player.js";

    let { index, track } = $props();
</script>

<div>
    <p>{`${index + 1}`.padStart(5, ' ')}</p>
    <button onclick={() => playTrack(track)}>{track.title}</button>
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
</style>