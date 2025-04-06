<script>
  import {currentRoute} from "../router.js";
  let {title, icon} = $props();
</script>

<div onclick={() => currentRoute.set(title)} class="{$currentRoute == title ? 'active' : ''}">
  <i class="fal fa-{icon}"></i>
  <p>{title}</p>
</div>

<style>
  div {
    height: 40px;
    margin-top: 5px;
    display: flex;
    align-items: center;
    cursor: pointer;
    padding-right: 15px;
    padding-left: 15px;
    background-color: transparent;
    border-radius: 150px;
    transition: 0.3s;
  }
  div.active {
    color: chartreuse;
  }

  div:hover {
    background-color: #3c3c3c
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