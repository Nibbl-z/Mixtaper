<script lang="ts">
    import Topbar from "../components/topbar.svelte";
    import Level from "../components/level.svelte"

    import type { LevelData, SearchResult } from "$lib/Types";
	import { onMount } from "svelte";
	import { PUBLIC_BACKEND_URL, PUBLIC_PROJECT_ID } from "$env/static/public";

    let results: LevelData[] = [];
    
    let covers: Record<string, string> = {}
    
    async function getResults() {
        const response = await fetch(PUBLIC_BACKEND_URL + "/recent_levels", {
            method: "GET",
        })

        if (response.ok) {
            const data: SearchResult = await response.json()
            results = data.message
            
            const coverPromises = results.map(result => {
                const coverUrl = `https://cloud.appwrite.io/v1/storage/buckets/cover_art/files/${result.$id}/view?project=${PUBLIC_PROJECT_ID}&project=${PUBLIC_PROJECT_ID}`
                return checkCoverArt(coverUrl).then(exists => ({
                    id: result.$id,
                    url: exists ? coverUrl : "/placeholder_coverart.png"
                }))
            })

            const loadedCovers = await Promise.all(coverPromises)
            covers = loadedCovers.reduce((acc: Record<string, string>, cover) => {
                acc[cover.id] = cover.url
                return acc
            }, {})
        }
    }
    
    async function checkCoverArt(url: string): Promise<boolean> {
        const response = await fetch(url)
        
        return response.ok
    }
    
    onMount(() =>{
        getResults()
    })
</script>

<Topbar/>

<div class="flex items-center flex-col justify-center my-10">
    <img src="/logo.png" alt="Logo" class="w-[30%] my-5">
    <h1 class="xl:text-5xl text-2xl font-Itim">Custom Levels and Mixtapes for Bits and Bops</h1>
</div>

<div class="flex items-center justify-center flex-col mt-10">
    <div class="w-full 2cols:max-w-[100em] max-w-[50em] flex justify-between items-center">
        <h1 class="text-left text-4xl">Recently Uploaded</h1>
    </div>
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8 my-4">
        {#each results as level}
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

<div class="flex flex-col min-h-screen">
    <div class="flex-grow"></div>

    <footer class="text-center py-10">
        <a class="text-2xl" href="/attributions">Attributions</a>
    </footer>
</div>


