<script>
    import PasswordPrompt from "./PasswordPrompt.svelte";
    import Stats from "./Stats.svelte";
    import Approval from "./Approval.svelte";

    let screen = "auth";
</script>

<div>
    {#if screen === "auth"}
        <PasswordPrompt onAuth={() => screen = 'dashboard'}></PasswordPrompt>
    {/if}
    {#if screen === "dashboard"}
        <div style="align-items: center">
            <Stats></Stats>
            <Approval></Approval>
        </div>
    {/if}
</div>

<style>
</style>