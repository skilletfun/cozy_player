<a {href} class={isCurrentTab() ? "active row center hover" : "row center hover"}>
  <i class="flex fal fa-{icon}"></i>
  <p>{title}</p>
</a>

<script>
  import { page } from "$app/state";
  let { title, icon, href } = $props();

  function isCurrentTab() {
    return page.route && page.route.id === href;
  }
</script>

<style>
  a {
    height: 40px;
    padding-right: 15px;
    padding-left: 15px;
    border-radius: 20px;
  }

  i {
    width: 32px;
    height: 32px;
    align-items: center;
  }
</style>
