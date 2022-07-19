<script>
    export let urls;
    export let onDeleteBlob;
</script>

<div>
    <table class="full-width">
        {#each urls as url, i}
            <tr>
                <td>
                    <audio src={url} class="full-width" controls></audio>
                </td>
                <td class="small">
                    <button on:click={() => onDeleteBlob(i)} class="delete">✕</button>
                </td>
            </tr>
        {/each}
    </table>
</div>

<style>
    .full-width {
        width: 100%;
    }
    .small {
        width: 10%;
    }
    .delete {
        background: #cd5c5c;
        margin-top: -2rem;
    }
</style>