<script>
    import StatisticCard from "$lib/components/StatisticCard.svelte";
    import {secondsToHumanString} from "$lib/utils.js";
    let { data } = $props();
</script>

<div class="container">
    <StatisticCard title="Artists" text={data.artists_total}/>
    <StatisticCard title="Tracks" text={data.tracks_total}/>
    <StatisticCard title="Total duration" text={secondsToHumanString(data.tracks_duration, true)}/>
    <StatisticCard title="Total played" text={secondsToHumanString(data.tracks_played, true)}/>
</div>

<style>
    .container {
        margin-top: 100px;
        margin-left: 250px;
        flex-wrap: wrap;
        display: flex;
        flex-direction: row;
    }
</style>