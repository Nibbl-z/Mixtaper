<script lang="ts">
	import { goto } from "$app/navigation";
	import { gameNames } from "$lib/levels";
	import type { Game } from "$lib/Types";
    import Chip from "./chip.svelte"
    export let songName: string
    export let songArtist: string
    export let cover: string
    export let bpm: number
    export let id: string
    
    export let gamesUsed: string[] = []
    let games: Game[] = []
    
    for (let game of gamesUsed) {
        if (gameNames[game] === undefined) continue
        if (game != "gameManager") {
            console.log(gameNames[game])
            games.push(gameNames[game])
        }
    }

    function onClick() {
        goto(`/level?id=${id}`)
    }
</script>

<div class="w-[50em] h-[15em] cursor-pointer" on:click={onClick} on:keydown={(e) => e.key === 'Enter' && onClick()} role="button" tabindex="0">
    <div class="bg-item hover:bg-itemHover p-3 w-[50em] h-[15em] rounded-3xl flex items-left shadow-2xl">
        <div class="h-[calc(15em-1.5rem)] w-[calc(15em-1.5rem)] overflow-hidden flex-shrink-0">
            <img src={cover != "" ? cover : "/placeholder_coverart.png"} alt="Cover Art" class="aspect-w-1 aspect-h-1 object-cover rounded-xl shadow-2xl w-full h-full">
        </div>
        <div class="mx-5 flex flex-col space-y-[-0.2em]">
            <h1 class="text-[2em] leading-[125%]">{songName}</h1>        
            <h1 class="text-[1.5em] text-author truncate">{songArtist}</h1>
            
            <div class="flex-grow"></div>
            <div class="flex flex-row space-x-1 mt-auto">
                {#each games as game}
                    <Chip label={game.name} color={game.color}/>
                {/each}
            </div>
        </div>
    </div>
    
    <div class='flex flex-col items-end float-right top-0 relative -translate-y-[4em] p-4'>
        <div class="flex-grow"></div>
        
        <div class="flex flex-row items-center space-x-2 justify-end">
            <p class="text-2xl">{bpm}</p>
            <img src="/bpm.png" alt="Duration Icon" class="h-[2em] p-1">
        </div>
    </div>
</div>
