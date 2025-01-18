<script lang="ts">
    import Topbar from "../../components/topbar.svelte"
    import Stat from "./stat.svelte"
    import Level from "../../components/level.svelte"
    import { page } from "$app/stores";
	import type { GetUserResult, LevelData, SearchResult } from "$lib/Types";
	import { loadLevels } from "$lib/levels";
	import { onMount } from "svelte";
	import { PUBLIC_BACKEND_URL } from "$env/static/public";
	import { getPfpFromId } from "$lib";

    export let displayName: string
    export let username: string
    export let bio: string
    export let pfp: string
    
    pfp = "/placeholder_pfp.png"
    displayName = "Loading..."
    username = "Loading..."
    bio = "Loading..."
    
    let levels: LevelData[] = [];
    let covers: Record<string, string> = {};
    
    let id = $page.url.searchParams.get('id') || '';

    async function load() {
        const params = new URLSearchParams({
            id: id
        })
        
        const response = await fetch(`${PUBLIC_BACKEND_URL}/get_user?${params}`, {
            method: "GET"
        })
    
        const userData: GetUserResult = await response.json()

        displayName = userData.message.DisplayName
        username = userData.message.Username
        pfp = await getPfpFromId(userData.message.ID, "/placeholder_pfp.png")
        bio = userData.message.Bio

        const levelsResponse = await fetch(`${PUBLIC_BACKEND_URL}/get_levels_from_user?${params}`, {
            method: "GET"
        })
    
        const levelData: SearchResult = await levelsResponse.json()
    
        const { results, covers: levelCovers } = await loadLevels(levelData);
        levels = results;
        covers = levelCovers;
    }
    
    onMount(() => {
        load()
    })
</script>

<Topbar/>

<div class="flex items-center justify-center flex-col mt-10 mb-10">
    <div class="bg-item w-[80%] rounded-3xl shadow-2xl mb-10 p-5">
        <div class="flex">
            <div class="w-[15%] flex-shrink-0">
                <h1 class="text-[2em] leading-[125%]">{displayName}</h1>        
                <h1 class="text-[1.5em] text-author truncate">@{username}</h1>
                <div class="aspect-w-1 aspect-h-1">
                    <img src={pfp} alt="{displayName}'s profile picture" class="object-cover rounded-full shadow-2xl w-full">
                </div>
            </div>
            <div class="flex-grow w-[85%]">
                <p class="text-[1.5em] text-description ml-10">{bio}</p>

                <!--<div class="flex flex-row items-center">
                    <Stat image="/heart.png" value="1,000" name="Downloads"/>
                    <Stat image="/song.png" value="100,000" name="Downloads"/>
                    <Stat image="/download.png" value="1,000,000" name="Downloads"/>
                </div> don't have the time rn to implement these, might do later--> 
            </div>
        </div>
    </div>
</div>

<div class="flex items-center justify-center flex-col">
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8 my-4">
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