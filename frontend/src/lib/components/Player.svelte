<script>
    import { APP_DATA } from "$lib/shared.svelte.js";
    import { API } from "$lib/api";
    import {
        playMainQueue,
        previousTrack,
        nextTrack,
        pausePlaying,
        resumePlaying,
    } from "$lib/player";
    import PlayerButton from "./PlayerButton.svelte";

    function getFunction() {
        return APP_DATA.isPlayingNow ? pausePlaying() : resumePlaying();
    }

    function getArtistName(artistId) {
        return APP_DATA.artists.find((e) => e.id == artistId).name;
    }
</script>

<div class="container">
    <div style="display:flex;justify-content:center;">
        {#if APP_DATA.currentTrack && APP_DATA.currentTrack.id}
            <img
                class="cover"
                src={API.Tracks.GetCoverURL(APP_DATA.currentTrack.id)}
                alt="Cover"
            />
            <div class="player-info">
                <p class="track-title">{APP_DATA.currentTrack.title}</p>
                <a
                    class="track-group"
                    href="/artists/{APP_DATA.currentTrack.artistId}"
                >
                    {getArtistName(APP_DATA.currentTrack.artistId)}
                </a>
                <div class="player-controls">
                    <PlayerButton icon="backward" onclick={previousTrack} />
                    <PlayerButton
                        icon={APP_DATA.isPlayingNow ? "pause" : "play"}
                        onclick={getFunction}
                    />
                    <PlayerButton icon="forward" onclick={nextTrack} />
                </div>
            </div>
        {:else}
            <button class="btn play" onclick={playMainQueue}>Play</button>
        {/if}
    </div>
</div>

<style>
    .container {
        display: flex;
        justify-self: center;
        height: 85vh;
        flex-direction: column;
        justify-content: center;
    }

    .player-info {
        display: flex;
        width: 600px;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    .player-controls {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 50%;
        margin-top: 25px;
    }

    p.track-title {
        font-size: 30px;
        text-align: center;
        width: 500px;
    }

    a.track-group {
        font-size: 20px;
        margin-top: -10px;
    }

    img {
        align-self: center;
        border-radius: 35px;
        height: 300px;
        width: 300px;
        margin-top: 15px;
        object-fit: cover;
        box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    }

    .btn {
        height: 80px;
        font-size: 36px;
        border-radius: 10px;
        background-color: transparent;
        color: #e08b5d;
        align-self: center;
        align-items: center;
        justify-content: center;
        border-radius: 150px;
        display: flex;
        padding-right: 35px;
        padding-left: 35px;
        transition: 0.3s;
    }
    .btn:hover {
        background-color: #3c3c3c;
    }
    .btn.play::before {
        margin-right: 15px;
        scale: 1.25;
        content: "▶";
    }
</style>
