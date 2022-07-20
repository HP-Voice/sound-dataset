<script>
    export let onSelect;
</script>

<div>
    <div class="center">
        <p>
            Нейромережу треба навчити відрізняти закляття від випадкових розмов, тому мені треба
            назбирати окремо записи заклять і записи розмов на будь-які теми (без згадки заклять).
        </p>
        <p>
            Тож оберіть будь ласка, що вам більше до вподоби:
        </p>
        <button class="continue" on:click={() => onSelect('spells')}>Закляття</button>
        <button class="continue" on:click={() => onSelect('nonsense')}>Випадкові розмови</button>
    </div>
</div>

<style>
    .center {
        text-align: center;
    }
    .continue {
        background: cadetblue;
    }
</style>
