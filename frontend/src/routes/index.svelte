<script context="module">
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */
	import { variables } from '$utils/variables'
	export async function load({fetch}) {
		const url = variables.backendAPIURL + `/hostname`;
		const res = await fetch(url);

		if (res.ok) {
			return {
				props: {
					returnedObj: await res.json()
				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	}
</script>

<script lang="ts">
	import { Shortlink, shortlink } from '$utils/shortlink_store';

	export let returnedObj
	let submit
	let response = null
	let url = ''

	function handleSubmit() {
	// Send a POST request to src/routes/contact.js endpoint
	submit = fetch(variables.backendAPIURL + "/create", {
		method: 'POST',
		body: JSON.stringify({ url: url }),
		headers: { 'content-type': 'application/json' },
	})
		.then(resp => resp.json())
		.then(resp => response = resp)
		.then(response => {
				$shortlink = new Shortlink(response)
				console.log($shortlink)
			})
		.finally(() => setTimeout(() => (submit = null), 5000))
		.catch(error => {
				console.log(error);
		});
	}

	function saveToClipboard(url:string) {
		navigator.clipboard.writeText(url).then(function() {
			console.log("set clipboard: " + url)
		}, function() {
			console.log("set clipoard failed")
		});
	}
</script>

<svelte:head>
    <title>shortlink</title>
</svelte:head>

<body class="flex flex-col min-h-screen">
	<main class="flex-grow">
		<div class="flex justify-center px-10 py-2">
			<img class="w-1/5 flex-shrink-0" src="https://raw.githubusercontent.com/michaelpeterswa/ShortLink/master/shortlink_logo.png" alt="shortlink">
		</div>
		<div class="px-5">
			<h2 class="text-center font-sans font-light text-yellow-500 text-2xl">welcome to the link shortener of your dreams</h2>
		</div>
		<div class="flex flex-grow justify-center space-x-0 p-5">
			<form on:submit|preventDefault={handleSubmit} method="post" class="flex w-3/5">
				<input 
					type="url"
					id="url"
					aria-label="url"
					name="url"
					autocomplete="url"
					placeholder="https://github.com"
					required
					bind:value={url}
					class="inline-block w-4/5 border-2 form-input py-3 rounded-tl-lg rounded-bl-lg border-yellow-700 border-solid">
				<button type="submit" class="inline-block flex-grow border-t-2 border-b-2 border-r-2 form-input py-3 px-2 rounded-tr-lg rounded-br-lg border-yellow-700 bg-yellow-500 text-white border-solid">submit</button>
			</form>
		</div>
		{#if $shortlink != null}
		<div class="flex justify-center">
			<div class="p-5 rounded-lg border-2 border-yellow-700">
				<p class="font-sans font-light text-yellow-500 text-md">shortlink: {$shortlink.shortlink}</p>
				<p class="font-sans font-light text-yellow-500 text-md">points to: {$shortlink.url}</p>
				<br>
				<button class="p-2 rounded-lg border-2 border-yellow-700 bg-yellow-500 text-white border-solid font-sans font-light text-md" on:click|once={() => saveToClipboard($shortlink.shortlink)}>save to clipboard</button>
			</div>
		</div>
		{/if}
	</main>

	<footer>
		<div class="p-5">
			<h2 class="text-center font-sans font-light text-yellow-500 text-2xl">node: {returnedObj.hostname}</h2>
		</div>
	</footer>
</body>
<style>
	/* This is a single-line comment */
</style>