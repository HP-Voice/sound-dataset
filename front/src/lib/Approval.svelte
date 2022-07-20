<script>
    import {getSampleForApproval, postVerdict} from "./api.js";

    let sample;
    let autoplay = false;

    function nextSample() {
        getSampleForApproval().then(s => sample = s);
    }

    function approve() {
        postVerdict(sample.id, 1).then(() => {nextSample(); autoplay = true});
    }

    function decline() {
        postVerdict(sample.id, -1).then(() => {nextSample(); autoplay = true});
    }

    nextSample();
</script>

<div>
    {#if sample}
        <h4>{sample.labelName}</h4>
        <table>
            <tr>
                <td>
                    <audio controls {autoplay} src={`/samples/${sample.id}.ogg`}></audio>
                </td>
                <td class="small">
                    <button on:click={approve} class="approve">👍</button>
                    <button on:click={decline} class="decline">👎</button>
                </td>
            </tr>
        </table>
    {/if}
</div>

<style>
    .approve {
        background: #9ed058;
    }
    .decline {
        background: #CD5C5CFF;
    }
    .small {
        width: 40%;
    }
</style>