<script>
    import TrackItem from "$lib/components/TrackItem.svelte";
    let { data } = $props();
</script>

<div>
    {#each data as track, index}
        <TrackItem {index} {track} />
    {/each}
</div>

<style>
    div {
        overflow-y: auto;
        height: 525px;
        scrollbar-width: thin;
        scrollbar-color: grey transparent;
    }
</style>
