<script>
    let {index, track} = $props();
</script>

<div style="display: flex;margin-left: -10px;align-items: center">
    <p style="margin-right: 20px;color: #777777;white-space: break-spaces">{`${index + 1}`.padStart(5, ' ')}</p>
    <a style="font-size: 16px">{track.title}</a>
</div>