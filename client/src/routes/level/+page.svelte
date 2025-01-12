<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
    import Stat from "./stat.svelte";
    import Chip from "../../components/chip.svelte";
	import { page } from "$app/stores";
	import type { Game, GetLevelResult, GetUserResult, MessageResult } from "$lib/Types";
	import { onMount } from "svelte";
	import { gameNames } from "$lib/levels";
    let songName: string
    let songArtist: string
    let cover: string
    let description: string

    let bpm: string
    let author: string
    let authorId: string
    let uploadDate: string
    let youtubeVideo: string

    songName = "Loading..."
    songArtist = "Loading..."
    description = "Loading..."
    cover = "/PLACEHOLDER.png"

    let id = $page.url.searchParams.get('id') || '';

    let games: Game[] = []
    let loading = true
    
    async function load() {
        const params = new URLSearchParams({
            id: id
        })
        
        const response = await fetch(`http://localhost:2050/get_level?${params.toString()}`, {
            method: "GET",
            headers: {
                "Origin": "http://localhost:5173"
            },
        })

        const data: GetLevelResult = await response.json()
        
        if (response.ok && data.successful) {
            songName = data.message.songName
            songArtist = data.message.songArtist
            description = data.message.description
            bpm = data.message.bpm.toString()
            uploadDate = data.message.$createdAt.slice(0, 10)
            authorId = data.message.uploader
        }

        for (let game of data.message.gamesUsed) {
            if (gameNames[game] === undefined) continue
            if (game != "gameManager") {
                console.log(gameNames[game])
                games.push(gameNames[game])
            }
        }

        const userResponse = await fetch(`http://localhost:2050/get_user?id=${data.message.uploader}`, {
            method: "GET"
        })

        const userData: GetUserResult = await userResponse.json()
        
        author = userData.message.Username
        console.log(data.message)
        youtubeVideo = getYoutubeId(data.message.youtubeVideo || "") || ""
        loading = false
    }

    async function gamesUsed(): Promise<Game[]> {
        while(loading) {
            await new Promise(resolve => setTimeout(resolve, 100))
        }

        return games
    }
    
    let downloadingMessage = ""
    let downloadingColor = "white"

    async function downloadLevel() {
        downloadingMessage = "Downloading..."
        downloadingColor = "white"

        const params = new URLSearchParams({
            id: id
        })

        const response = await fetch(`http://localhost:2050/download_riq?${params}`, {
            method: "GET"
        })

        if (!response.ok) {
            const data: MessageResult = await response.json()

            downloadingMessage = data.message
            downloadingColor = "resultError"

            return
        }

        const blob = await response.blob()
        const url = window.URL.createObjectURL(blob)
        const link = document.createElement("a")
        
        link.href = url
        link.download = `${songName}.riq`
        document.body.appendChild(link)
        link.click()
        window.URL.revokeObjectURL(url)
        link.remove()
    
        downloadingMessage = "Downloaded successfully!"
        downloadingColor = "resultSuccess"
    }

    function getYoutubeId(url: string): string | null {
        const regex = /(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/;
        const match = url.match(regex);

        return match ? match[1] : null;
    }

    async function getYoutubeEmbedId() {
        while(loading) {
            await new Promise(resolve => setTimeout(resolve, 100))
        }

        return youtubeVideo
    }
    
    onMount(() => {
        load()
    })
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
                    <a href={`/user?id=${authorId}`}><Stat name="Author" value="{author || "Loading..."}"/></a>
                    <Stat name="BPM" value="{bpm || "Loading..."}"/>
                    <Stat name="Upload Date" value={uploadDate || "Loading..."}/>
                </div>
               
                <h1 class="w-full text-center text-[3em]">Games Used</h1>
                
                <div class="flex flex-wrap gap-2 mt-4">
                    {#await gamesUsed()}
                    <p class="w-full h-full text-center text-3xl">Loading...</p>
                    {:then gamesUsed} 
                        {#each gamesUsed as game}
                            <Chip label={game.name} color={game.color}/>
                        {/each}
                    {/await}
                </div>
            </div>
            
            <button onclick={downloadLevel} class="w-full rounded-2xl bg-interactable text-[#FFFFFF] text-2xl p-2 mt-5 flex items-center justify-center">
                <img src="/download.png" alt="Search icon" class="h-6 w-6 mr-2"/>
                Download
            </button>

            <p class="w-full text-center text-2xl text-{downloadingColor}">{downloadingMessage}</p>
        </div>
        
        <div class="flex-grow h-fit bg-item ml-10 rounded-3xl shadow-2xl p-5">
            {#await getYoutubeEmbedId()}
            <p class="w-full h-full text-center text-3xl">Loading...</p>
            {:then id}
                {#if id != ""}
                <div class="relative w-full pt-[56.25%]">
                    <iframe class="absolute top-0 left-0 w-full h-full" src="https://www.youtube.com/embed/{id}" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen title="YouTube video player"></iframe>
                </div>
                {:else}
                <p class="w-full h-full text-center text-3xl">No preview video was provided :(</p>
                {/if}
            {/await}
            
        </div>
    </div>
</div>