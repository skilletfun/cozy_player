<div class="navigation row center">
  <NavButton title="Home" icon="home" href="/" />
  <NavButton title="Artists" icon="user-music" href="/artists" />
  <ActionButton title="Rescan" icon="redo" onclick={Rescan} />
</div>

<script>
  import NavButton from "$lib/components/NavButton.svelte";
  import ActionButton from "./ActionButton.svelte";
  import { Rescan } from "$lib/services/library";
</script>
