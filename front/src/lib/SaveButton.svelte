<script>
    export let enabled = false;
    export let onClick;
</script>

<div class="center">
    {#if enabled}
        <button on:click={onClick} class="save">💾</button>
    {:else}
        <button disabled class="disabled">💾</button>
    {/if}
</div>

<style>
    .save {
        background: #53805b;
    }
    .disabled {
        background: rgb(48, 58, 49);
    }
    .center {
        text-align: center;
    }
</style>