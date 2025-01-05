<script lang="ts">
    export let label: string = ""
    export let id: string
    export let required: boolean = false
    export let big: boolean = false
    export let file: boolean = false

    export let fileTitle: string = ""
    export let fileMaxSize: number = 0

    let width: string = label == "" ? "100%" : "70%"
</script>

<div class="flex flex-row">
    <p class="text-[2em]">
        {label}{#if required}
        <span class="text-[#FF0000]">*</span>
        {/if}
    </p>
    {#if big} 
        <textarea {id} class="w-[{width}] h-[6em] rounded-xl text-[2em] p-[0.2em] ml-auto"></textarea>
    {:else if file}
        <div class="w-[{width}] ml-auto bg-[#262329] border-dotted border-[#FFFFFF] border-4 p-2 flex items-center justify-center flex-col">
            <p class="w-full text-center text-2xl">{fileTitle}</p>
            <img src="/upload.png" alt="Upload Icon" class="w-[4em] my-2">
            <p class="w-full text-center text-xl">(max file size {fileMaxSize}mb)</p>
        </div>
    {:else}
        <input type="text" {id} class="w-[{width}] h-[1.5em] rounded-xl text-[2em] p-[0.2em] ml-auto"/>
    {/if}
    
</div>