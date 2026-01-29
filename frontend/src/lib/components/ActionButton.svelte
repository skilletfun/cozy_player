<button onclick={() => onclick()} class="row center hover h-40 p-15 br-20">
  <i class="flex fal fa-{icon}"></i>
  <p>{title}</p>
</button>

<script>
  let { title, icon, onclick } = $props();
</script>

<style>
  button {
    background-color: transparent;
    font-size: 12pt;
    color: #e08b5d;
  }
</style>
