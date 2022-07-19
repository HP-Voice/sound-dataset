<script>
    export let mode;
    export let prompt;
</script>

<div>
    <div class="center">
        {#if mode === "spells"}
            <h1>Скажіть "{prompt}"</h1>
            <p>Ви можете записати будь-яку кількість семплів. Чим більше - тим краще, особливо в різних інтонаціях.</p>
        {/if}
        {#if mode === "nonsense"}
            <h2>Говоріть що завгодно, або ось вам випадковий текст:</h2>
            <p>{prompt}</p>
        {/if}
    </div>
</div>

<style>
    .center {
        text-align: center;
    }
</style>