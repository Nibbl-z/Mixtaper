<script lang="ts">
    export let label: string;
    export let color: string
</script>

<div class="rounded-full px-2 py-1 text-[#FFFFFF]" style="background-color: {color}">
    {label}
</div>