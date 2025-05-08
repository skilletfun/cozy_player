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
        width: 70%;
        margin-top: 20px;
        margin-left: 125px;
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    }
</style>