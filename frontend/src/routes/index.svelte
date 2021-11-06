<script context="module">
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */
	import { variables } from '$utils/variables'
	export async function load({fetch}) {
		const url = variables.backendAPIURL + `/versions`;
		const res = await fetch(url);

		if (res.ok) {
			return {
				props: {
					versions: await res.json()
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
	import { Count, count } from '$utils/count_store';

	export let versions = []
	async function getCount() {
			fetch(variables.frontendAPIURL + `/count`)
			.then(response => response.json())
			.then(data => {
				$count = new Count(data)
				console.log($count.count)
			}).catch(error => {
				console.log(error);
				return "0";
		});
	}
</script>

<svelte:head>
    <title>sigma</title>
</svelte:head>

<main>
	<div class="p-10">
		<h1 class="text-center font-mono font-normal text-yellow-700 text-6xl">σ</h1>
	</div>
	<div class="p-5">
		<h2 class="text-center font-sans font-light text-yellow-500 text-2xl">welcome to the sigma stack</h2>
	</div>
	<div class="p-5">
		<h4 class="text-center font-sans font-normal text-yellow-500 text-xl">apps:</h4>
		<ul>
			{#each versions as app}
				<li class="text-center font-sans font-light text-yellow-500 text-xl">{app.name} - {app.version}</li>
			{/each}
		</ul>
	</div>
	<div class="p-5 flex justify-center">
		<button class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded-full" on:click={getCount}>
			{$count.count}
		</button>
	</div>
</main>

<style>
	/* This is a single-line comment */
	div {
		align-content: center;
	} 
</style>