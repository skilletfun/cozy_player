<script>
  import { onMount } from "svelte";
  import { Notifications, acts } from "@tadashi/svelte-notification";
  import Header from "$lib/components/Header.svelte";
  import { setupPlayer, removeHandlers } from "$lib/services/player.js";

  let { children } = $props();

  onMount(() => {
    setupPlayer();
    return () => {
      removeHandlers();
      if ("mediaSession" in navigator) {
        navigator.mediaSession.metadata = null;
      }
    };
  });
</script>

<main class="column">
  <Header />

  <div class="column content">
    {@render children()}
  </div>

  <Notifications />
</main>

<style>
  div.content {
    margin-left: 100px;
    margin-right: 100px;
  }
</style>
