<script>
  let { title, text } = $props();
</script>

<div>
  <h2>{title}</h2>
  <h1>{text}</h1>
</div>

<style>
  div {
    width: 460px;
    height: 195px;
    background-color: #3c3c3c;
    border-radius: 10px;
    margin-top: 45px;
    margin-right: 45px;
    max-width: 550px;
    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    align-content: space-around;
    border: 2px white;
  }
  h2 {
    margin-top: 35px;
    display: flex;
    justify-content: center;
  }
  h1 {
    margin-top: 25px;
    display: flex;
    justify-content: center;
  }

</style>