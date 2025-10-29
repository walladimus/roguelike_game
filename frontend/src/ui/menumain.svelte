<script lang="ts">
import MenuLobby from "./menulobby.svelte"
import { onMount } from "svelte";
import { getServerVersion } from "../net/api";
import MenuGame from "./menugame.svelte"

let serverStatus = "Checking server...";
let showPlayOptions = false
let inLobby = false
let inGame = false

onMount(async () => {
  const version = await getServerVersion();
  serverStatus = version
    ? 'Server Online - v${version}'
    : "Server Offline";
});

function togglePlayOptions() {
  showPlayOptions = !showPlayOptions
}

function createGame() {
  showPlayOptions = !showPlayOptions
}

function joinGame() {
  inLobby = true
}

function startGame() {
  inLobby = false
  inGame = true
}

</script>

{#if inGame}
  <MenuGame />
{:else if inLobby}
  <MenuLobby onLeave={() => (inLobby = false)} onStart={startGame} />
{:else}
<Panel title="Main Menu">
    <p>{serverStatus}</p>
    <Button label="Play" on:click={togglePlayOptions} />
    {#if showPlayOptions}
      <Panel title="Play Options">
        <Button label="Create Game" on:click={createGame} />
        <Button label="Join Game" on:click={joinGame} />
        </Panel>
    {/if}
    <Button label="Achievements" />
    <Button label="Friends" />
    <Button label="Settings" />
    <Button label="Notices & Requests" />
    <Button label="Buy Me a Coffee" />
</Panel>
{/if}