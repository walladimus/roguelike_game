<script lang="ts">
    export let onLeave: () => void
    let players = ["Alice", "Bob", "Charlie"]
    export let onStart: () => void
</script>

<Panel title="Lobby">
    <ul>
        {#each players as player}
            <li>{player}</li>
        {/each}
    </ul>

    <Button label="Start Game" on:click={onStart} />
    <Button label="Leave Lobby" on:click{onLeave} />
</Panel>