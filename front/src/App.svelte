<script>
    import HelloScreen from "./lib/HelloScreen.svelte";
    import CategorySelect from "./lib/CategorySelect.svelte";
    import Recorder from "./lib/Recorder.svelte";
    import RecordScreen from "./lib/RecordScreen.svelte";

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
       <HelloScreen onStart="{() => {screen = 'category'}}"></HelloScreen>
    {/if}
    {#if screen === "category"}
        <CategorySelect onSelect="{(c) => {screen = c}}"></CategorySelect>
    {/if}
    {#if screen === "spells" || screen === "nonsense"}
        <RecordScreen mode={screen}></RecordScreen>
    {/if}
</main>

<style>
</style>
