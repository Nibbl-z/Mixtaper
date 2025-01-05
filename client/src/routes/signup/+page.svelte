<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
    import type { MessageResult } from "$lib/Types"
    let emailField: HTMLInputElement
    let usernameField: HTMLInputElement
    let passwordField: HTMLInputElement
    
    let result = ""
    let resultColor = "resultSuccess"
    
    async function signup() {
        const response = await fetch("http://localhost:2050/signup", {
            method : "POST",
            body : JSON.stringify({
                email : emailField.value,
                username : usernameField.value,
                password : passwordField.value
            })
        })
        const data: MessageResult = await response.json()
        
        if (data.successful) {
            resultColor = "resultSuccess"
        } else {
            resultColor = "resultError"
        }

        result = data.message
    } 
</script>

<Topbar/>

<div class="flex items-center justify-center min-h-screen">
    <div class="bg-item w-[400px] h-[600px] p-8 rounded-3xl shadow-2xl flex flex-col">
        <h1 class="w-[100%] text-center text-4xl">Create an Account</h1>
        
        <p class="text-[1.25em] mt-8 mb-1">Email</p>
        
        <input bind:this={emailField} type="text" id="email" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">
        
        <p class="text-[1.25em] mt-2 mb-1">Username</p>
        
        <input bind:this={usernameField} type="text" id="username" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <p class="text-[1.25em] mt-2 mb-1">Password</p>
        
        <input bind:this={passwordField} type="password" id="password" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <div class="mt-auto">
            <p class="w-full text-center text-xl text-{resultColor}">{result}</p>
            <button on:click={signup} class="rounded-2xl bg-interactable text-[#FFFFFF] w-[100%] text-2xl p-2">Signup</button>
        </div>
    </div>
</div>