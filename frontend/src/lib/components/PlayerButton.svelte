<script>
  let { icon, onclick } = $props();
</script>

<button onclick={ onclick } aria-label="player control button">
  <i class="fas fa-{icon}"></i>
</button>

<style>
  button {
    height: 64px;
    width: 64px;
    display: flex;
    align-items: center;
    background-color: transparent;
    justify-content: center;
    transition: 0.3s;
  }

  i {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    scale: 2.25;
  }
</style>