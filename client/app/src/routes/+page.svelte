<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_STAGE, PUBLIC_SHUFFLE_SHOWDOWN_API_PATH } from '$env/static/public';

	let message = 'loading message...';

	onMount(async () => {
		const apiPath = `/${PUBLIC_STAGE}/${PUBLIC_SHUFFLE_SHOWDOWN_API_PATH}`;

		const response = await fetch(`${apiPath}/hello`).catch((error) => {});

		if (response?.ok) {
			message = await response.text();
		}
	});
</script>

<p>{message}</p>
