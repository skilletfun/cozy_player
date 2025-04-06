<script>
    import {artistsStore, loadArtists} from "../stores/artists.js";
    import ArtistCard from "./ArtistCard.svelte";

    if (!$artistsStore.length) loadArtists();
</script>

<div>
    {#each $artistsStore as artist }
        <ArtistCard artist={artist} />
    {/each}
</div>

<style>
    div {
        display: grid;
        margin-top: 50px;
        margin-left: 200px;
        grid-template-columns: auto auto auto auto auto auto;
    }
</style>