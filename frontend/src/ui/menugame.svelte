<script lang="ts">
import { onMount } from 'svelte';
import { fetchGameState, sendMove } from '$lib/net/api';
import { setGameState } from '$lib/state/LocalPlayerState';
import { getPlayerState, setPlayerState } from '$lib/state/LocalPlayerState';
import { tryMove } from '$lib/systems/MovementSystem';

let map = string[][] = [];
let player = getPlayerState();
let turn = 0:

onMount(() => {
    fetchGameState().then((response) => {
        setGameState(response);
        map = response.map;
    });

    window.addEventListener('keydown', handleKey);
    return () => window.removeEventListener('keydown', handleKey);
});

function handleKey(event: KeyboardEvent) {
    const keyMap = {
        ArrowUp: 'up',
        ArrowDown: 'down',
        ArrowLeft: 'left',
        ArrowRight: 'right',
        w: 'up',
        s: 'down',
        a: 'left'
        d: 'right'
    };

    const direction = keyMap[event.key];
    if (direction) {
        const moved = tryMove(direction);
        if (moved) {
            const state = getGameState();
            map = state.map;
            turn = state.turn;

            setPlayerState({
                hp:100,
                inventory: [],
                position: state.player
            });
            player = getPlayerState();
        }
    }
}
</script>

<!--<Panel title="Game View">
    <div class="grid">
        {#each map as row}
            {#each row as cell}
                <div class="tile {cell}">{cell === "player" ? "" : ""}</div>
            {/each}
        {/each}
    </div>
</Panel>

<div class="overlay">
    <p>HP: {player.hp}</p>
    <p>Turn: {turn}</p>
</div>

<div class="inventory">
    <h3>Inventory</h3>
    <ul>
        {#each player.inventory as item}
            <li>{item}</li>
        {/each}
    </ul>
</div>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(5, 40px);
        grid-template-rows: repeat(5, 40px);
        gap: 2px;
    }

    .tile {
        width: 40px;
        height: 40px;
        border: 1px solid #ccc;
    }

    .wall {
        background-color: #444;
    }

    .floor {
        background-color: #eee;
    }

    .player {
        background-color: gold;
    }

    .overlay {
        position: absolute;
        top: 1rem;
        left: 1rem;
        background: rgba(0, 0, 0.6);
        color: white;
        padding: 0.5rem;
        border-radius: 4px;
        font-family: monospace;
    }

    .inventory {
        position: absolute;
        top: 1rem;
        left: 1rem;
        background: rgba(0, 0, 0.6);
        color: white;
        padding: 0.5rem;
        border-radius: 4px;
        font-family: monospace;
    }
</style>-->