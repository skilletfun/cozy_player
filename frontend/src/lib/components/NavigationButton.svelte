<script>
  let {title, icon, href} = $props();
  import { page } from '$app/state';

  function isCurrentTab() {
    return page.route && page.route.id === href;
  }
</script>

<a href="{href}" class="{isCurrentTab() ? 'active' : ''}">
  <i class="fal fa-{icon}"></i>
  <p>{title}</p>
</a>

<style>
  a {
    height: 40px;
    margin-top: 5px;
    display: flex;
    align-items: center;
    padding-right: 15px;
    padding-left: 15px;
    background-color: transparent;
    border-radius: 150px;
    transition: 0.3s;
  }
  a.active {
    color: chartreuse;
  }

  a:hover {
    background-color: #3c3c3c;
  }

  p {
    font-size: 12pt;
  }

  i {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
  }
</style>