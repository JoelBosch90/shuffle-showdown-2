<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_SHUFFLE_SHOWDOWN_API_PATH } from '$env/static/public';

	let message = 'loading message...';
	const connectionMessages: string[] = [];

	onMount(async () => {
		const apiPath = `/${PUBLIC_SHUFFLE_SHOWDOWN_API_PATH}`;

		const response = await fetch(`${apiPath}/hello`).catch(() => {});

		if (response?.ok) {
			message = await response.text();
		}

		const connection = new WebSocket(`wss://${window.location.host}${apiPath}/socket`);

		connection.onopen = () => {
			connectionMessages.push('WebSocket connection established');
			connection.send('Hello Server!');
		};

		connection.onmessage = (event) => {
			connectionMessages.push(`Message from server: ${event.data}`);
		};

		connection.onerror = (error) => {
			connectionMessages.push('WebSocket error occurred', JSON.stringify(error));
		};

		connection.onclose = () => {
			connectionMessages.push('WebSocket connection closed');
		};
	});
</script>

<h1>Hello?</h1>
<p>{message}</p>
<h2>WebSocket Messages:</h2>
<ul>
	{#each connectionMessages as connectionMessage}
		<li>{connectionMessage}</li>
	{/each}
</ul>
