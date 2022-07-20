<script>
    import {getStats} from "./api.js";

    let stats;

    function refresh() {
        getStats().then(s => {stats = s});
        setTimeout(refresh, 3000);
    }
    refresh();
</script>

<div>
    {#if stats}
        <table>
            <thead>
            <tr>
                <th>
                    #
                </th>
                <th>
                    Name
                </th>
                <th>
                    Amount
                </th>
            </tr>
            </thead>
            <tbody>
            {#each stats.labels as label}
                <tr>
                    <td>{label.id}</td>
                    <td>{label.name}</td>
                    <td>{label.amount}</td>
                </tr>
            {/each}
            </tbody>
        </table>
        <p>Total amount: {stats.totalAmount}</p>
    {/if}
</div>

<style>

</style>