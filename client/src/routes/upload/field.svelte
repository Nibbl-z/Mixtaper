<script lang="ts">
	import type { HTMLTextareaAttributes } from "svelte/elements";
    import { onDestroy } from "svelte";
    export let label: string = ""
    export let required: boolean = false
    export let big: boolean = false
    export let file: boolean = false

    export let fileTitle: string = ""
    export let fileMaxSize: number = 0
    export let isImage: boolean = false
    
    let width: string = label == "" ? "100%" : "70%"
    
    let field: HTMLInputElement
    let textArea: HTMLTextAreaElement
    let fileUpload: File

    export function getField(): string {
        if (field != null) {
            return field.value
        } else {
            return textArea.value
        }
    }

    export function getFile(): File {
        return fileUpload
    }
    
    let imageUrl: string

    $: if (fileUpload && isImage) {
        imageUrl = URL.createObjectURL(fileUpload)
    }
    
    onDestroy(() => {
        if (imageUrl && isImage) {
            URL.revokeObjectURL(imageUrl)
        }
    });

    let hovering = false

    function hover(event: DragEvent) {
        event.preventDefault()
        hovering = true
    }

    function leave(event: DragEvent) {
        event.preventDefault()
        hovering = false
    }
    
    function drop(event: DragEvent) {
        event.preventDefault()
        hovering = false

        if (event.dataTransfer?.files) {
            fileUpload = event.dataTransfer.files[0]
        }
    }
</script>

<div class="flex flex-row">
    <p class="text-[2em]">
        {label}{#if required}
        <span class="text-[#FF0000]">*</span>
        {/if}
    </p>
    {#if big} 
        <textarea bind:this={textArea} class="w-[{width}] h-[6em] rounded-xl text-[2em] p-[0.2em] ml-auto"></textarea>
    {:else if file}
        <div 
        role="button" 
        tabindex=0 
        aria-label="Drop files here" 
        on:drop={drop} 
        on:dragover|preventDefault={hover}
        on:dragleave={leave}
        class="w-[{width}] ml-auto {hovering ? "bg-[#63555a]" : "bg-[#262329]"} border-dotted border-[#FFFFFF] border-4 p-2 flex items-center justify-center flex-col">
            <p class="w-full text-center text-2xl">{fileUpload ? fileUpload.name : fileTitle}</p>
            <img src={isImage ? (imageUrl ? imageUrl : "/upload.png") : "/upload.png"} alt="Upload Icon" class="w-[4em] my-2">
            <p class="w-full text-center text-xl">(max file size {fileMaxSize}mb)</p>
        </div>
    {:else}
        <input bind:this={field} type="text" class="w-[{width}] h-[1.5em] rounded-xl text-[2em] p-[0.2em] ml-auto"/>
    {/if}
    
</div>