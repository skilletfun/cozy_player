<div class="mt-2">
  {#each data.artists as artist}
    <ArtistCard {artist} />
  {/each}
</div>

<script>
  import ArtistCard from "$lib/components/ArtistCard.svelte";
  
  let { data } = $props();
</script>

<style>
  div {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    overflow-y: auto;
    scrollbar-width: none;
    grid-auto-rows: min-content;
    height: 90vh;
  }
</style>
