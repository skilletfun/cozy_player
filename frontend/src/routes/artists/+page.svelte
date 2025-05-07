<script>
    import ArtistCard from "$lib/components/ArtistCard.svelte";
    let { data } = $props();
</script>

<div>
    {#each data.artists as artist }
        <ArtistCard artist={artist} />
    {/each}
</div>

<style>
    div {
        display: grid;
        margin-top: 20px;
        margin-left: 150px;
        grid-template-columns: auto auto auto auto auto auto auto auto;
    }
</style>