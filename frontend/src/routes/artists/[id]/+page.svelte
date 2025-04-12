<script>
    import { secondsToHumanString } from "$lib/utils.js";
    import TrackItem from "$lib/TrackItem.svelte";
    import {API} from "$lib/api.js";
    let { data } = $props();
</script>

<div class="artist-view">
    <div class="artist-header">
        <img src="{API.Artists}cover/{data.artist.id}/" alt="cover">
        <div>
            <h1>{data.artist.name}</h1>
            <h2>{data.artist.tracks_count} Tracks • {secondsToHumanString(data.artist.duration)}</h2>
            <button class="play-btn"><i class="fas fa-play"></i> Play</button>
            <button class="more"><i class="fas fa-ellipsis-h"></i></button>
        </div>
    </div>
    <div style="overflow-y: scroll; height:550px;width: 100%">
        {#each data.tracks as track, index}
            <TrackItem index={index} track={track}/>
        {/each}
    </div>
</div>

<style>
    .artist-view {
        margin-top: 70px;
        margin-left: 200px;
        width: 100%;
    }
    .artist-header {
        display: flex;
        margin-bottom: 40px;
    }
    img {
        width: 200px;
        border-radius: 25px;
        margin-right: 30px;
    }
    .play-btn {
        height: 35px;
        width: 90px;
        border-radius: 18px;
        border-width: 0;
        cursor: pointer;
    }
    .play-btn:active {
        background-color: #373737;
    }
    .more {
        margin-left: 15px;
        height: 35px;
        width: 35px;
        border-radius: 18px;
        border-width: 0;
        cursor: pointer;
        background-color: transparent;
    }
</style>