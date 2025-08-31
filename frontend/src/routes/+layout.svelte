<main class=column>
  <Header />
  
  <div class="column content">
    {@render children()}
  </div>

  <Notifications />
</main>

<script>
  import { onMount } from "svelte";
  import { Notifications, acts } from "@tadashi/svelte-notification";
  import NavigationButton from "$lib/components/NavButton.svelte";
  import Header from "$lib/components/Header.svelte";
  import { playMainQueue, setupPlayer } from "$lib/player.js";
  import { API } from "$lib/api.js";
  import { Rescan } from "$lib/library.js";

  let { children } = $props();

  onMount(() => {
    setupPlayer();
    return () => {
      if ("mediaSession" in navigator) {
        navigator.mediaSession.metadata = null;
      }
    };
  });
</script>

<style>
  div.content {
    margin-left: 100px;
    margin-right: 100px;
  }
</style>
