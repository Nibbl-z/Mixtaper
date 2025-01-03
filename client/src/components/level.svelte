<script lang="ts">
	import { onMount } from "svelte";
    import Chip from "./chip.svelte";

    export let songName: string
    export let songArtist: string
    export let description: string
    export let cover: string
</script>

<div class="bg-item w-[45em] h-[15em] p-3 rounded-3xl flex items-left m-4 shadow-2xl">
    <img src={cover} alt="Cover Art" class="h-[100%] rounded-xl shadow-2xl">
    <div class="mx-5 flex flex-col space-y-[-0.5em]">
        <h1 class="text-[3em] truncate">{songName}</h1>
        <h1 class="text-[2em] text-author truncate">{songArtist}</h1>
        <p class="text-description text-sm truncate">{description}</p>
        
        <div class="flex-grow"></div>
        <div class="flex flex-row space-x-1 mt-auto">
            <Chip label="Mixtape"/>
        </div>
    </div>
    <div class='flex flex-col items-end justify-end h-full ml-auto'>
        <div class="flex-grow"></div>
        
        <div class="flex flex-row items-center space-x-2 justify-end">
            <p class="text-2xl">2:00</p>
            <img src="/time.png" alt="Duration Icon" class="h-[2em] p-1">
        </div>

        <div class="flex flex-row items-center space-x-2 justify-end">
            <p class="text-2xl">120</p>
            <img src="/bpm.png" alt="Duration Icon" class="h-[2em] p-1">
        </div>
    </div>
</div>