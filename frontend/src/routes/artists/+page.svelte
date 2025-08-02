<script>
    import ArtistCard from "$lib/components/ArtistCard.svelte";
    let { data } = $props();
</script>

<div>
    {#each data.artists as artist}
        <ArtistCard {artist} />
    {/each}
</div>

<style>
    div {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    }
</style>
