<script>
    import {loadTracksByArtist, tracksStore} from "../stores/tracks.js";
    import {artistsURL, currentArtist, currentArtistInfo, getArtistInfo} from "../stores/artists.js";
    import {onMount} from "svelte";

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
                <h2>Tracks: {$currentArtistInfo.tracks_count}  •  Duration: {$currentArtistInfo.duration}</h2>
            </div>
        </div>
    {/if}
</div>

<!--{#each $tracksStore as track}-->
<!--    <p>{track.title}</p>-->
<!--{/each}-->

<style>
    .artist-view {
        margin-top: 70px;
        margin-left: 200px;
    }
    .artist-header {
        display: flex;
    }
    img {
        width: 200px;
        border-radius: 25px;
        margin-right: 30px;
    }
</style>