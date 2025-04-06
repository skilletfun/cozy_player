<script>
    import {rescanLibrary} from "../stores/library.js";
</script>

<div class="container">
    <div class="buttons">
        <button onclick={() => rescanLibrary()}>Rescan Library</button>
        <button>Play</button>
    </div>
    <p>Home</p>
</div>

<style>
    .container {
        margin-top: 200px;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: start;
        flex-direction: column;
    }
    button {
        width: 250px;
        height: 75px;
        font-size: 16pt;
        background-color: transparent;
        border: 2px solid rgb(75 70 74);
        border-radius: 10px;
        cursor: pointer;
        transition: 0.1s;
        margin-inline: 15px;
    }
    button:hover {
        font-size: 16.5pt;
        background-color: #3c3c3c;
    }
    .buttons {
        margin-bottom: 100px;
    }
</style>