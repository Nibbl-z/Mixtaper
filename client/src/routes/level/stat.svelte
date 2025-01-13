<script lang="ts">
    export let name: string
    export let value: string
    export let hover: boolean = false
</script>

<div class="flex flex-row p-4">
    <p class="text-[1.5em]">{name}</p>
    <p class="ml-auto text-[1.5em] {hover ? "hover:underline" : ""}">{value}</p>
</div>