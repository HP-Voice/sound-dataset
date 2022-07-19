<script>
    import Prompt from "./Prompt.svelte";
    import Recorder from "./Recorder.svelte";
    import List from "./List.svelte";
    import SaveButton from "./SaveButton.svelte";
    import {getLabels, getRandomSpell, getSentence, postSample} from "./api.js";

    export let mode;
    let labels = [];
    let labelId;
    let prompt = '';

    if (mode === "spells") {
        getLabels().then(l => {
            labels = l;
            getRandomSpell().then((id) => {
                labelId = id;
                let randomSpell = labels.find(label => label.id === id);
                if (randomSpell === null)
                    return;
                prompt = randomSpell.name;
            });
        });
    }

    if (mode === "nonsense") {
        getSentence().then(s => prompt = s);
    }

    let blobs = [];
    let urls = [];

    function onBlobAvailable(blob) {
        blobs.push(blob);
        urls.push(URL.createObjectURL(blob));
        if (mode === "nonsense")
            getSentence().then(s => prompt = s);
        urls = urls;
    }

    function onDeleteBlob(index) {
        blobs.splice(index, 1);
        urls.splice(index, 1);
        urls = urls;
    }

    function createSamplePoster(blob, next) {
        return () => postSample(labelId, blob).then(next);
    }

    function onSaveClick() {
        if (mode === "spells") {
            let func = () => {
                getRandomSpell().then((id) => {
                    labelId = id;
                    let randomSpell = labels.find(label => label.id === id);
                    if (randomSpell === null)
                        return;
                    prompt = randomSpell.name;
                });
                blobs = [];
                urls = [];
            }
            for (let blob of blobs) {
                func = createSamplePoster(blob, func);
            }
            func();
        }
        if (mode === "nonsense") {
            labelId = 1;
            let func = () => {
                getSentence().then(s => prompt = s);
                blobs = [];
                urls = [];
            }
            for (let blob of blobs) {
                func = createSamplePoster(blob, func);
            }
            func();
        }
    }
</script>

<div>
    <Prompt {mode} {prompt}></Prompt>
    <Recorder {onBlobAvailable}></Recorder>
    <List {urls} {onDeleteBlob}></List>
    <SaveButton onClick={onSaveClick} enabled={urls.length > 0}></SaveButton>
</div>

<style>

</style>