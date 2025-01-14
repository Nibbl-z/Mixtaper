<script lang="ts">
	import { onMount, type SvelteComponent } from "svelte";
    import Topbar from "../../components/topbar.svelte"
    import Field from "./field.svelte";
	import type { MessageResult, PostResult } from "$lib/Types";
    import { getCookie } from "$lib";
	import { json } from "@sveltejs/kit";
	import { PUBLIC_BACKEND_URL } from "$env/static/public";
    
    interface FieldData {
        label: string
        required: boolean
        big: boolean
        file: boolean
        fileTitle?: string
        fileMaxSize?: number
        fileIsImage?: boolean
        element?: InstanceType<typeof Field>
    }
    
    let fields: Record<string, FieldData> = {
        songName: {
            label: "Song Name",
            required: true,
            big: false,
            file: false,
        },
        songArtist: {
            label: "Song Artist",
            required: true,
            big: false,
            file: false
        },
        description: {
            label: "Description",
            required: false,
            big: true,
            file: false
        },
        coverArt: {
            label: "Cover Art",
            required: false,
            big: false,
            file: true,
            fileTitle: "Drag and drop or select .png/.jpg file to upload",
            fileMaxSize: 5,
            fileIsImage: true
        },
        youtubeVideo: {
            label: "YouTube Video",
            required: false,
            big: false,
            file: false,
        },
        riq: {
            label: "",
            required: false,
            big: false,
            file: true,
            fileTitle: "Drag and drop or select .riq file to upload",
            fileMaxSize: 20,
            fileIsImage: false
        }
    }

    let result = ""
    let resultColor = "resultSuccess"
    
    async function uploadCoverArt(file: File, id?: string): Promise<MessageResult> {
        const arrayBuffer = await file.arrayBuffer();
        const bytes = new Uint8Array(arrayBuffer);

        const token = getCookie("token");

        const response = await fetch(PUBLIC_BACKEND_URL + "/upload_cover_art", {
            method : "POST",
            headers : (token && id) ? {
                "Authorization" : token,
                "ID" : id
            } : undefined,
            body : bytes
        })
        
        const data: MessageResult = await response.json()
        return data
    }

    async function uploadRiq(file: File, id?: string): Promise<MessageResult> {
        const arrayBuffer = await file.arrayBuffer();
        const bytes = new Uint8Array(arrayBuffer);
        
        const token = getCookie("token");

        const data: MessageResult = await fetch(PUBLIC_BACKEND_URL + "/upload_riq", {
            method : "POST",
            headers : (token && id) ? {
                "Authorization" : token,
                "ID" : id
            } : undefined,
            body : bytes
        }).then(response => response.json())
        .catch(error => {
            console.log(error)
            throw error
        })
        
        return data
    }

    async function deleteUpload(id: string | undefined) {
        const token = getCookie("token");

        const data: MessageResult = await fetch(PUBLIC_BACKEND_URL + "/delete_level", {
            method: "POST",
            headers : token ? {
                "Authorization" : token
            } : undefined,
            body: id
        }).then(response => response.json())
        .catch(error => {
            console.log(error)
            throw error
        })
    }

    async function upload() {
        let riq = fields["riq"].element?.getFile()
        
        if (riq == undefined) {
            return
        }
        
        const token = getCookie("token");
        
        const postData: Record<string, string> = {}

        for (const [id, field] of Object.entries(fields)) {
            if (id !== "riq" && id !== "coverArt") {
                let input = field.element?.getField()
                
                if (input !== undefined && input !== "") {
                    postData[id] = input
                }
            }
        }

        const postLevelData: PostResult = await fetch(PUBLIC_BACKEND_URL + "/post_level", {
            method : "POST",
            headers : token ? {
                "Content-Type" : "application/json",
                "Authorization" : token
            } : undefined,
            body : JSON.stringify(postData)
        }).then(response => response.json())
        .catch(error => {
            console.log(error)
            throw error
        })
        
        if (!postLevelData.successful) {
            result = postLevelData.message
            resultColor = "resultError"
            return
        }
        
        if (fields["coverArt"].element?.getFile() !== undefined) {
            const response = await uploadCoverArt(fields["coverArt"].element?.getFile(), postLevelData.id)
                
            if (!response.successful) {
                result = response.message
                resultColor = "resultError"

                deleteUpload(postLevelData.id)

                return
            }
        }

        const riqResponse = await uploadRiq(riq, postLevelData.id)
                
        if (!riqResponse.successful) {
            result = riqResponse.message
            resultColor = "resultError"

            deleteUpload(postLevelData.id)

            return
        }
        
        result = postLevelData.message
        resultColor = "resultSuccess"
    }
    
</script>

<Topbar/>

<h1 class="w-full text-center text-[4em] my-10">Uploading Level</h1>

<div class="flex items-center justify-center mt-10">
    <div class="bg-item w-[50%] p-8 rounded-3xl shadow-2xl flex flex-col gap-4">
        {#each Object.values(fields) as field}
            <Field
            bind:this={field.element}
            label={field.label} 
            required={field.required}
            big={field.big}
            file={field.file}
            fileTitle={field.fileTitle}
            fileMaxSize={field.fileMaxSize}
            isImage={field.fileIsImage}
            />
        {/each}
        
        <p class="w-full text-center text-2xl text-{resultColor}">{result}</p>
        
        <div class="flex items-center justify-center">
            <button on:click={upload} class="rounded-2xl bg-interactable hover:bg-interactableHover text-[#FFFFFF] w-[20%] text-2xl p-2 mt-auto">Post</button>
        </div>
    </div>
</div>