<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_STAGE, PUBLIC_SHUFFLE_SHOWDOWN_API_PATH } from '$env/static/public';

	let message = 'loading message...';

	onMount(async () => {
		const apiPath = `/${PUBLIC_STAGE}/${PUBLIC_SHUFFLE_SHOWDOWN_API_PATH}`;

		const response = await fetch(`${apiPath}/hello`).catch(() => {});

		if (response?.ok) {
			message = await response.text();
		}
	});
</script>

<h1>Hello?</h1>
<p>{message}</p>
