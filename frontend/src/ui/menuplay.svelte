<script lang="ts">
    import { connectEcho } from "../net/ws";
    import { onMount } from "svelte";

    let echoResponse = "Waiting for echo..."

    onMount(() => {
        connectEcho((msg) => {
            echoResponse = 'Received: ${msg}'
        })
    })
</script>

<Panel title="WebSocket Test">
    <p>{echoResponse}</p>
</Panel>