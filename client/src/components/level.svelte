<script lang="ts">
    import Chip from "./chip.svelte"
    export let songName: string
    export let songArtist: string
    export let cover: string
</script>

<div class="w-[50em] h-[15em]">
    <div class="bg-item p-3 w-[50em] h-[15em] rounded-3xl flex items-left shadow-2xl">
        <img src={cover} alt="Cover Art" class="h-[100%] rounded-xl shadow-2xl">
        <div class="mx-5 flex flex-col space-y-[-0.2em]">
            <h1 class="text-[2em] leading-[125%]">{songName}</h1>        
            <h1 class="text-[1.5em] text-author truncate">{songArtist}</h1>
            
            <div class="flex-grow"></div>
            <div class="flex flex-row space-x-1 mt-auto">
                <Chip label="Mixtape"/>
            </div>
        </div>
    </div>
    
    <div class='flex flex-col items-end float-right top-0 relative -translate-y-[6em]'>
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
