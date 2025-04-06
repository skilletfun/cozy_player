<script>
  import Navigation from "./lib/Navigation.svelte";
  import {router, currentRoute} from "./router.js";
</script>

<main>
  <Navigation/>
  <svelte:component this={$router[$currentRoute]}/>
</main>

<style>
  main {
    display: flex;
  }
</style>