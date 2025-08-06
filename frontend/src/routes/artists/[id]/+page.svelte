<script>
    import { secondsToHumanString } from "$lib/utils.js";
    import { API } from "$lib/api.js";
    import TrackList from "$lib/components/TrackList.svelte";
    import { playArtist } from "$lib/player.js";

    let { data } = $props();
    let artistInfoHeader = `${data.artist.tracksCount} Tracks • ${secondsToHumanString(data.artist.duration)}`;
</script>

<div class="artist-view">
    <div class="artist-header">
        <img src={API.Artists.GetCoverURL(data.artist.id)} alt="cover" />
        <div>
            <h1>{data.artist.name}</h1>
            <h2>{artistInfoHeader}</h2>
            <button class="play-btn" onclick={() => playArtist(data.artist)}>
                <i class="fas fa-play"></i> Play
            </button>
        </div>
    </div>
    <TrackList data={data.tracks} />
</div>

<style>
    .artist-view {
        align-self: center;
        margin-top: 50px;
    }
    .artist-header {
        display: flex;
        margin-bottom: 40px;
    }
    img {
        width: 200px;
        height: 200px;
        object-fit: cover;
        border-radius: 25px;
        margin-right: 30px;
        box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
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
</style>
