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
    
    let results: LevelData[] = [];
    
    async function getResults() {
        const params = new URLSearchParams({
            query: query
        })
        
        const response = await fetch(`http://localhost:2050/search?${params.toString()}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json",
                "Origin": "http://localhost:5173"
            }
        })
        
        if (response.ok) {
            const data: SearchResult = await response.json()
        
            results = data.message
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
        {#each results as level}
            <Level 
            songName={level.songName} 
            songArtist={level.songArtist} 
            cover="/PLACEHOLDER.png" 
            gamesUsed={level.gamesUsed} 
            bpm={level.bpm} 
            duration="1:20"
            />
        {/each}
    </div>    
</div>