<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
    import Stat from "./stat.svelte";
    import Chip from "../../components/chip.svelte";
    export let songName: string
    export let songArtist: string
    export let cover: string
    export let description: string

    songName = "Awesome Song"
    songArtist = "Awesome Artist"
    description = "Even more amazing description. Did I ever mention my extreme dislike for CSS?"
    cover = "/PLACEHOLDER.png"
</script>

<Topbar/>

<div class="flex items-center justify-center flex-col mt-10 mb-10">
    <div class="bg-item p-3 w-[80%] rounded-3xl flex items-left shadow-2xl mb-10">
        <img src={cover} alt="Cover Art" class="h-auto max-h-[15em] rounded-xl shadow-2xl">
        <div class="mx-5 flex flex-col space-y-[-0.2em]">
            <h1 class="text-[2em] leading-[125%]">{songName}</h1>        
            <h1 class="text-[1.5em] text-author truncate">{songArtist}</h1>
            <p class="text-[1em] text-description">{description}</p>
        </div>
    </div>
    
    <div class="flex flex-row justify-between w-[80%] h-full">
        <div class="w-[30%]">
            <div class="w-full bg-item self-start shadow-2xl rounded-3xl p-4">
                <div class="flex flex-col space-y-[-1em]">
                    <Stat name="Author" value="Nibbles"/>
                    <Stat name="BPM" value="128"/>
                    <Stat name="Duration" value="1:23"/>
                    <Stat name="Downloads" value="1,000"/>
                    <Stat name="Upload Date" value="October 1, 1999"/>
                </div>
               
                <h1 class="w-full text-center text-[3em]">Games Used</h1>
                
                <div class="flex flex-wrap gap-2 mt-4">
                    <Chip label="Flipper Snapper"/> <Chip label="Meet & Tweet"/> <Chip label="Hammer Time"/> <Chip label="A really long chip name"/> <Chip label="shortchip"/> 
                </div>
            </div>
            
            <button class="w-full rounded-2xl bg-interactable text-[#FFFFFF] text-2xl p-2 mt-5 flex items-center justify-center">
                <img src="/download.png" alt="Search icon" class="h-6 w-6 mr-2"/>
                Download
            </button>
        </div>
        
        <div class="flex-grow h-fit bg-item ml-10 rounded-3xl shadow-2xl p-5">
            <div class="relative w-full pt-[56.25%]">
                <iframe class="absolute top-0 left-0 w-full h-full" src="https://www.youtube.com/embed/TjKslu3v97I" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen title="YouTube video player"></iframe>
            </div>
        </div>
    </div>
</div>
