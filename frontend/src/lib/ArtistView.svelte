<script>
    import {loadTracksByArtist, tracksStore} from "../stores/tracks.js";
    import {artistsURL, currentArtist, currentArtistInfo, getArtistInfo} from "../stores/artists.js";
    import {onMount} from "svelte";
    import {secondsToHumanString} from "../utils.js";
    import TrackItem from "./TrackItem.svelte";

    onMount(async function() {
        currentArtistInfo.set(await getArtistInfo($currentArtist));
        tracksStore.set(await loadTracksByArtist($currentArtist));
    });
</script>

<div class="artist-view">
    {#if $currentArtistInfo}
        <div class="artist-header">
            <img src="{artistsURL}cover/{$currentArtist}/" alt="cover">
            <div>
                <h1>{$currentArtistInfo.name}</h1>
                <h2>{$currentArtistInfo.tracks_count} Tracks • {secondsToHumanString($currentArtistInfo.duration)}</h2>
                <button class="play-btn"><i class="fas fa-play"></i> Play</button>
                <button class="more"><i class="fas fa-ellipsis-h"></i></button>
            </div>
        </div>
        <div style="overflow-y: scroll; height:550px;width: 100%">
            {#each $tracksStore as track, index}
                <TrackItem index="{index}" track="{track}"/>
            {/each}
        </div>
    {/if}
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