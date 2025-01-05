<script lang="ts">
	import { onMount } from "svelte";
    import Topbar from "../../components/topbar.svelte";
    import { page } from '$app/stores';
	import { getCookie } from "$lib";
	import type { MessageResult } from "$lib/Types";

    let userId = $page.url.searchParams.get('userId') || '';
    let secret = $page.url.searchParams.get('secret') || '';
    
    let message = ""

    async function verify() {
        const token = getCookie("token");
        const params = new URLSearchParams({
            userId : userId,
            secret : secret
        })
        const response = await fetch(`http://localhost:2050/verify_email?${params.toString()}`, {
            method : "PUT",
            headers : token ? {
                "Authorization" : token
            } : undefined
        })

        const data: MessageResult = await response.json()

        if (data.successful) {
            message = data.message
            console.log(data.message)
        } else {
            message = data.message
            console.log(data.message)
        }
    }
    
    onMount(() => {
        verify()
    })
</script>

<Topbar/>

<div class="flex items-center flex-col justify-center my-10">
    <h1 class="xl:text-5xl text-2xl font-Itim">{message}</h1>
</div>