<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
    import type { LoginResult, MessageResult } from "$lib/Types"

    let identifierField: HTMLInputElement
    let passwordField: HTMLInputElement
    
    let result = ""
    let resultColor = "resultSuccess"

    async function login() {
        const response = await fetch("http://localhost:2050/login", {
            method : "POST",
            body : JSON.stringify({
                identifier : identifierField.value,
                password : passwordField.value
            })
        })
        
        const data: LoginResult = await response.json()
        result = data.message
        if (data.successful) {
            resultColor = "resultSuccess"
            document.cookie = `token=${data.secret}; path=/;`
        } else {
            resultColor = "resultError"
            return
        }

        const verifyResponse = await fetch("http://localhost:2050/send_email_verification", {
            method : "POST",
            headers : {
                "Authorization" : data.secret || ""
            }
        })
        
        const verifyData: MessageResult = await response.json()
        if (verifyData.successful) {
            console.log(verifyData.message)
        }
    }
</script>

<Topbar/>

<div class="flex items-center justify-center min-h-screen">
    <div class="bg-item w-[400px] h-[600px] p-8 rounded-3xl shadow-2xl flex flex-col">
        <h1 class="w-[100%] text-center text-4xl">Login</h1>
        
        <p class="text-[1.25em] mt-8 mb-1">Username / Email</p>
        
        <input bind:this={identifierField} type="text" id="email" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <p class="text-[1.25em] mt-2 mb-1">Password</p>
        
        <input bind:this={passwordField} type="password" id="password" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em] mb-2">
        
        <u class="text-center"><a href="/" class="text-xl">Forgot password?</a></u>
        
        <div class="mt-auto">
            <p class="text-xl mb-2">Don't have an account? <u class="text-center"><a href="/signup" class="text-xl">Sign up here</a></u></p>
            <p class="w-full text-center text-xl text-{resultColor}">{result}</p>
            <button on:click={login} class="rounded-2xl bg-interactable hover:bg-interactableHover text-[#FFFFFF] w-[100%] text-2xl p-2">Login</button>
        </div>
       
    </div>
</div>