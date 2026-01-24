<button onclick={() => onclick()} class="row center hover">
  <i class="flex fal fa-{icon}"></i>
  <p>{title}</p>
</button>

<script>
  import { page } from "$app/state";
  let { title, icon, onclick } = $props();
</script>

<style>
  button {
    height: 40px;
    padding-right: 15px;
    padding-left: 15px;
    border-radius: 20px;
    background-color: transparent;
    font-size: 12pt;
    color: #e08b5d;
  }

  i {
    width: 32px;
    height: 32px;
    align-items: center;
  }
</style>
