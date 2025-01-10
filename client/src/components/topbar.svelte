<script lang="ts">
	import { goto } from "$app/navigation";
	import { deleteCookie, getCookie, getID, getPfp } from "$lib";
	import type { MessageResult } from "$lib/Types";
	import { onMount } from "svelte";

    let searchbar: HTMLInputElement

    function search(event: KeyboardEvent) {
        if (event.key !== "Enter") {return}
 
        window.location.href = `/search?query=${searchbar.value}`;
    }
    
    async function logout() {
        deleteCookie("token")
        goto("/")
        window.location.reload()
    }
</script>

<div class="overflow=hidden bg-topbar h-[4em] p-[0.2em] flex items-center justify-center relative">
    <a href="/" class="h-[100%]"><img src="/logo.png" class="text-[#FFFFFF] text-[30px] h-[90%] absolute left-0 ml-4 hover:brightness-75" alt="Mixtaper Logo"/></a>
    <div class="mx-auto flex-grow max-w-[50%] h-[75%]">
        <img src="/search.png" alt="Search icon" class="absolute h-[40%] brightness-75 py top-[50%] translate-y-[-50%] translate-x-2" />
        <input bind:this={searchbar} onkeydown={search} type="text" id="search" class="w-[100%] h-[100%] rounded-xl text-[2em] p-[0.5em] pl-[1.5em]">
    </div>
    
    <div class="flex items-center space-x-4 absolute right-0 h-[100%] mr-4">
        {#await getID()}
            <a href="/signup" class="h-[100%] flex items-center">
                <img src="/upload.png" class="text-[#FFFFFF] h-[70%] hover:brightness-75" alt="Upload"/>
            </a>
        {:then id}
            {#if id != ""}
            <a href="/upload" class="h-[100%] flex items-center">
                <img src="/upload.png" class="text-[#FFFFFF] h-[70%] hover:brightness-75" alt="Upload"/>
            </a>
            {:else}
            <a href="/signup" class="h-[100%] flex items-center">
                <img src="/upload.png" class="text-[#FFFFFF] h-[70%] hover:brightness-75" alt="Upload"/>
            </a>
            {/if}
        {/await}
        
        <div class="relative h-[100%] group">
            <a href="/" class="h-[100%] flex items-center">
                {#await getPfp()}
                    <img src="/account.png" class="text-[#FFFFFF] h-[70%] hover:brightness-75" alt="Account"/>
                {:then pfp} 
                    <img src={pfp} class="text-[#FFFFFF] h-[70%] hover:brightness-75" alt="Account"/>
                {/await}
            </a>
            <div class="flex flex-col text-center opacity-0 absolute right-0 bg-topbar group-hover:opacity-100 pointer-events-none group-hover:pointer-events-auto">
                {#await getID()}
                    <a class="text-3xl hover:bg-topbarHover p-4" href="/login">Login</a>
                    <a class="text-3xl hover:bg-topbarHover p-4" href="/signup">Signup</a>
                {:then id}
                    {#if id != ""}
                        <a class="text-3xl hover:bg-topbarHover p-4" href="/user?id={id}">Profile</a>
                        <a class="text-3xl hover:bg-topbarHover p-4" href="/settings">Settings</a>
                        <a class="text-3xl hover:bg-topbarHover p-4" onclick={logout} href="/">Log Out</a>
                    {:else}
                        <a class="text-3xl hover:bg-topbarHover p-4" href="/login">Login</a>
                        <a class="text-3xl hover:bg-topbarHover p-4" href="/signup">Signup</a>
                    {/if}
                {/await}
            </div>
        </div>
    </div>
</div>