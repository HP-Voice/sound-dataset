<script>
    export let onBlobAvailable;

    let state = "idle";
    let chunks = [];
    let mediaRecorder;

    if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
        navigator.mediaDevices.getUserMedia({audio: true})
            .then(stream => {
                mediaRecorder = new MediaRecorder(stream);
                mediaRecorder.ondataavailable = e => {
                    chunks.push(e.data);
                }

                mediaRecorder.onstop = e => {
                    let blob = new Blob(chunks, {type: "audio/ogg; codecs=opus"});
                    chunks = [];
                    onBlobAvailable(blob);
                }
            })
            .catch(e => {
                console.log(e);
            });
    } else {
        console.log("getUserMedia not supported :(");
    }

    function onRecordClick() {
        state = "recording";
        mediaRecorder.start();
    }

    function onPauseClick() {
        state = "recording paused";
        mediaRecorder.pause();
    }

    function onStopClick() {
        state = "idle";
        mediaRecorder.stop();
    }

</script>

<div class="recorder">
    {#if mediaRecorder !== null}
        {#if state === "idle" || state === "recording paused"}
            <button class="record" on:click={onRecordClick}>⏺</button>
        {/if}
        {#if state === "recording"}
            <button class="pause" on:click={onPauseClick}>⏸</button>
        {/if}
        {#if state === "recording" || state === "recording paused"}
            <button class="stop" on:click={onStopClick}>■</button>
        {/if}
    {/if}
</div>

<style>
    .recorder {
        text-align: center;
        margin-bottom: 1rem;
    }
    .record {
        background: #9ed058;
        height: 3rem;
        width: 4rem;
    }
    .pause {
        background: #DAA520FF;
        height: 3rem;
        width: 4rem;
    }
    .stop {
        background: #cd5c5c;
        height: 3rem;
        width: 4rem;
    }
</style>