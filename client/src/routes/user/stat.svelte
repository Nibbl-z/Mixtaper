<script lang="ts">
    export let image: string
    export let name: string
    export let value: string
</script>

<div class="flex flex-row items-center h-14 mx-4 mt-4 mb-2">
    <img src={image} alt={name} class="h-14 w-14"/>
    <p class="ml-4 text-3xl">{value} {name}</p>
</div>