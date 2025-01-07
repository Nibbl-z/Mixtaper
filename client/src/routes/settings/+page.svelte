<script lang="ts">
	import { getCookie } from "$lib";
	import type { MessageResult } from "$lib/Types";
	import { onMount } from "svelte";
    import Topbar from "../../components/topbar.svelte";
    import Field from "./field.svelte";

    let displayName: InstanceType<typeof Field>
    let username: InstanceType<typeof Field>
    let password: InstanceType<typeof Field>

    let bio: HTMLTextAreaElement
    let bioSubmit: HTMLButtonElement
    
    let result = ""
    let resultColor = "resultSuccess"

    async function setDisplayName() {
        const token = getCookie("token");
        
        const response = await fetch("http://localhost:2050/change_display_name", {
            method: "POST",
            headers: token ? {
                "Authorization" : token
            } : undefined,
            body : displayName.getField()
        })
        
        const data: MessageResult = await response.json()
        
        resultColor = data.successful ? "resultSuccess" : "resultError"
        result = data.message
    }

    async function setUsername() {
        const token = getCookie("token");
        
        const response = await fetch("http://localhost:2050/change_username", {
            method: "POST",
            headers: token ? {
                "Authorization" : token
            } : undefined,
            body : username.getField()
        })
        
        const data: MessageResult = await response.json()
        
        resultColor = data.successful ? "resultSuccess" : "resultError"
        result = data.message
    }

    async function setPassword() {
        const passwordInput = password.getMultipleFields()
    
        const previousPassword = passwordInput[0]
        const newPassword = passwordInput[1]
        const repeatedPassword = passwordInput[2]
        
        if (newPassword != repeatedPassword) {
            result = "Passwords do not match!"
            resultColor = "resultError"
            return 
        }

        const token = getCookie("token");
        
        const response = await fetch("http://localhost:2050/change_password", {
            method: "POST",
            headers: token ? {
                "Authorization" : token
            } : undefined,
            body : JSON.stringify({
                previous: previousPassword,
                new: newPassword
            })
        })
        
        const data: MessageResult = await response.json()
        
        resultColor = data.successful ? "resultSuccess" : "resultError"
        result = data.message
    }

    async function setBio() {
        const token = getCookie("token");
        
        const response = await fetch("http://localhost:2050/change_bio", {
            method: "POST",
            headers: token ? {
                "Authorization" : token
            } : undefined,
            body : bio.value
        })
        
        const data: MessageResult = await response.json()
        
        resultColor = data.successful ? "resultSuccess" : "resultError"
        result = data.message
    }
    
    onMount(() => {
        displayName.getButton().onclick = setDisplayName
        username.getButton().onclick = setUsername
        password.getButton().onclick = setPassword
        bioSubmit.onclick = setBio
    })
</script>

<Topbar/>

<h1 class="w-full text-center text-[4em] my-10">Account Settings</h1>

<div class="flex items-center justify-center mt-10">
    <div class="bg-item w-[50%] p-8 rounded-3xl shadow-2xl flex flex-col gap-4">
        <h1 class="w-full text-center text-[3em] mb-8">Account Details</h1>
        <p class="w-full text-center text-[2em] text-{resultColor}">{result}</p>
        <Field bind:this={displayName} label="Display Name" placeholder="Display Name"/>
        <Field bind:this={username} label="Username" placeholder="Username"/>
        <Field bind:this={password} label="Change Password" placeholder="" type="password" mutlipleFields={["Previous Password", "New Password", "Repeat Password"]}/>

        <h1 class="w-full text-center text-[3em] mt-4">Profile</h1>
        
        <p class="text-[2em]">About Me</p>
        <textarea bind:this={bio} id="aboutMe" class="w-full h-[8em] rounded-xl text-[2em] p-[0.2em] ml-auto"></textarea>
        <button bind:this={bioSubmit} class="rounded-2xl bg-interactable text-[#FFFFFF] w-full] text-[2em] mt-auto">Change</button>
    
        <p class="text-[2em]">Profile Picture</p>
        
        <div class="h-[15em] flex flex-row flex-shrink-0">
            <div class="aspect-w-1 aspect-h-1 h-[15em] w-[15em]">
                <img src="/PLACEHOLDER.png" alt="" class="object-cover rounded-full shadow-2xl">
            </div>
            <div class="w-[70%] ml-auto bg-[#262329] border-dotted border-[#FFFFFF] border-4 p-2 flex items-center justify-center flex-col">
                <p class="w-full text-center text-2xl">Drag and drop or select .png/.jpg file to upload</p>
                <img src="/upload.png" alt="Upload Icon" class="w-[4em] my-2">
                <p class="w-full text-center text-xl">(max file size 1mb)</p>
            </div>
        </div>
       
    </div>
</div>