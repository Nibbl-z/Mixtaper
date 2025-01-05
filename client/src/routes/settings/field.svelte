<script lang="ts">
    export let label: string = ""
    export let id: string
    export let mutlipleFields: string[] = []


</script>

<div class="flex flex-row">
    <p class="text-[2em]">{label}</p>
    
    <div class="ml-auto w-[65%]">
        
        {#if mutlipleFields.length == 0}
            <input type="text" {id} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
            <button class="rounded-2xl bg-interactable text-[#FFFFFF] w-[30%] text-[2em] ml-2 mt-auto">Change</button>
        {:else}
            <div class="flex flex-col gap-4">
                {#each mutlipleFields as field, i}
                    {#if i == mutlipleFields.length - 1}
                        <div>
                            <input type="text" id={field} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
                            <button class="rounded-2xl bg-interactable text-[#FFFFFF] w-[30%] text-[2em] ml-2 mt-auto">Change</button>
                        </div>
                    {:else}
                        <input type="text" id={field} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
                    {/if}
                {/each}
            </div>
        {/if}
    </div>
</div>