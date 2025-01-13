<script lang="ts">
    export let label: string = ""
    export let placeholder: string
    export let mutlipleFields: string[] = []
    export let type: string = "text"
    let fields: HTMLInputElement[] = []

    let field: HTMLInputElement
    export function getField(): string {
        return field.value
    }

    export function getMultipleFields(): string[] {
        let fieldStrings: string[] = []
        
        for (let input of fields) {
            fieldStrings.push(input.value)
        }

        return fieldStrings
    }
    
    let button: HTMLButtonElement
    export function getButton(): HTMLButtonElement {
        return button
    }
    
</script>

<div class="flex flex-row">
    <p class="text-[2em]">{label}</p>
    
    <div class="ml-auto w-[65%]">
        
        {#if mutlipleFields.length == 0}
            <input bind:this={field} {type} placeholder={placeholder} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
            <button bind:this={button} class="rounded-2xl bg-interactable hover:bg-interactableHover text-[#FFFFFF] w-[30%] text-[2em] ml-2 mt-auto">Change</button>
        {:else}
            <div class="flex flex-col gap-4">
                {#each mutlipleFields as f, i}
                    {#if i == mutlipleFields.length - 1}
                        <div>
                            <input bind:this={fields[i]} {type} placeholder={f} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
                            <button bind:this={button} class="rounded-2xl bg-interactable hover:bg-interactableHover text-[#FFFFFF] w-[30%] text-[2em] ml-2 mt-auto">Change</button>
                        </div>
                    {:else}
                        <input bind:this={fields[i]} {type} placeholder={f} class="w-[65%] h-[1.5em] rounded-2xl text-[2em] p-[0.2em]"/>
                    {/if}
                {/each}
            </div>
        {/if}
    </div>
</div>