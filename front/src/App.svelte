<script>
    import HelloScreen from "./lib/HelloScreen.svelte";
    import CategorySelectScreen from "./lib/CategorySelectScreen.svelte";
    import RecordScreen from "./lib/RecordScreen.svelte";
    import DashboardScreen from "./lib/DashboardScreen.svelte";

    let screen = "hello";

    function recorded(blob) {
        let a = document.createElement("a");
        a.href = window.URL.createObjectURL(blob);
        a.click();
        document.removeChild(a);
    }

</script>

<main>
    {#if screen === "hello"}
       <HelloScreen onStart="{() => screen = 'category'}"></HelloScreen>
    {/if}
    {#if screen === "category"}
        <CategorySelectScreen onSelect="{(c) => screen = c}"></CategorySelectScreen>
    {/if}
    {#if screen === "spells" || screen === "nonsense"}
        <RecordScreen mode={screen}></RecordScreen>
    {/if}
    {#if screen === 'dashboard'}
        <DashboardScreen></DashboardScreen>
    {:else}
        <small on:click={() => screen = 'dashboard'}><u>Admin?</u></small>
    {/if}
</main>

<style>
</style>
