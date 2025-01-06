<script lang="ts">
    import Topbar from "../components/topbar.svelte";
    import Level from "../components/level.svelte"

    import type { LevelData, SearchResult } from "$lib/Types";
	import { onMount } from "svelte";

    let results: LevelData[] = [];
    
    let covers: Record<string, string> = {}
    
    async function getResults() {
        const response = await fetch("http://localhost:2050/recent_levels", {
            method: "GET",
            headers: {
                "Origin": "http://localhost:5173"
            },
        })

        if (response.ok) {
            const data: SearchResult = await response.json()
            results = data.message
            
            const coverPromises = results.map(result => {
                const coverUrl = `https://cloud.appwrite.io/v1/storage/buckets/cover_art/files/${result.$id}/view?project=676f205d000370a15786&project=676f205d000370a15786`
                return checkCoverArt(coverUrl).then(exists => ({
                    id: result.$id,
                    url: exists ? coverUrl : "/PLACEHOLDER.png"
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

<div class="flex items-center justify-center flex-col">
    <div class="w-full 2cols:max-w-[100em] max-w-[50em] flex justify-between items-center">
        <h1 class="text-left text-3xl">Recently Uploaded</h1>
        <a href="/" class="text-right text-3xl">Show More &gt;</a>
    </div>
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8 my-4">
        {#each results as level}
            <Level 
            songName={level.songName} 
            songArtist={level.songArtist} 
            cover={covers[level.$id]}
            gamesUsed={level.gamesUsed} 
            bpm={level.bpm} 
            duration="1:20"
            />
        {/each}
    </div>    
</div>





<a href="https://www.flaticon.com/free-icons/search" title="search icons">Search icons created by Chanut - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/user" title="user icons">User icons created by Freepik - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/upload" title="upload icons">Upload/Download icons created by Ilham Fitrotul Hayat - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/song" title="song icons">Song icons created by Freepik - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/heart" title="heart icons">Heart icons created by Chanut - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/rythm" title="rythm icons">Rythm icons created by Freepik - Flaticon</a>
\
<a href="https://www.flaticon.com/free-icons/clock" title="clock icons">Clock icons created by Ilham Fitrotul Hayat - Flaticon</a>