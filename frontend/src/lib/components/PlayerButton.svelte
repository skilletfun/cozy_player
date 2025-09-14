<button {onclick} aria-label="player control button" class="row-center">
  <i class="row-center fas fa-{icon}"></i>
</button>

<script>
  let { icon, onclick } = $props();
</script>

<style>
  button {
    height: 64px;
    width: 64px;
  }

  i {
    width: 32px;
    height: 32px;
    scale: 2.25;
  }
</style>
