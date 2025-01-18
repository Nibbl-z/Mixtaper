<script lang="ts">
	import Topbar from "../../components/topbar.svelte"
    import Level from "../../components/level.svelte"
    import { page } from '$app/stores';
    export let query: string
    export let resultsCount: number
    query = $page.url.searchParams.get('query') || '';
    resultsCount = 0

    import type { LevelData, SearchResult } from "$lib/Types";
	import { onMount } from "svelte";
	import { PUBLIC_BACKEND_URL } from "$env/static/public";
	import { loadLevels } from "$lib/levels";
    
    let levels: LevelData[] = []
    let covers: Record<string, string> = {}
    
    async function getResults() {
        const params = new URLSearchParams({
            query: query
        })
        
        const response = await fetch(`${PUBLIC_BACKEND_URL}/search?${params.toString()}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json",
                "Origin": "http://localhost:5173"
            }
        })
        
        if (response.ok) {
            const data: SearchResult = await response.json()
        
            const { results, covers: levelCovers } = await loadLevels(data);
            levels = results;
            covers = levelCovers;
            resultsCount = results.length
            console.log(results)
        }
    }
    
    onMount(() =>{
        getResults()
    })
    
</script>

<Topbar/>

<div class="my-2 ml-8 p-[4rem]">
    <h1 class="text-4xl">Showing {resultsCount} results for {query}</h1>
</div>

<div class="flex items-center justify-center flex-col">
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8">
        {#each levels as level}
            <Level 
            id={level.$id}
            songName={level.songName} 
            songArtist={level.songArtist} 
            cover={covers[level.$id]}
            gamesUsed={level.gamesUsed} 
            bpm={level.bpm} 
            />
        {/each}
    </div>    
</div>