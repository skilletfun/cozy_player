<script>
    import TrackItem from "$lib/components/TrackItem.svelte";
    let { data } = $props();
</script>

<div>
    {#each data as track, index}
        <TrackItem index={index} track={track}/>
    {/each}
</div>

<style>
    div {
        overflow-y: scroll;
        height: 525px;
        scrollbar-width: thin;
        scrollbar-color: grey transparent;
    }
</style>