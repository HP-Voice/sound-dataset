<script>
    import {auth} from "./api.js";
    export let onAuth;

    let value = "";
    function onClick() {
        auth(value).then(onAuth);
    }
</script>

<div>
    <label>
        Password:
        <input type="password" bind:value={value}>
    </label>
    <button on:click={onClick}>OK</button>
</div>