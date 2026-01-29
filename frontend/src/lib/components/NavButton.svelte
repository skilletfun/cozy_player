<a {href} class={`row center hover h-40 p-15 br-20 ${isCurrentTab() ? 'active' : ''}`}>
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
