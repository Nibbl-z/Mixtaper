<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
    import Field from "./field.svelte";
</script>

<Topbar/>

<h1 class="w-full text-center text-[4em] my-10">Account Settings</h1>

<div class="flex items-center justify-center mt-10">
    <div class="bg-item w-[50%] p-8 rounded-3xl shadow-2xl flex flex-col gap-4">
        <h1 class="w-full text-center text-[3em] mb-8">Account Details</h1>
        
        <Field label="Display Name" id="displayName"/>
        <Field label="Username" id="username"/>
        <Field label="Email" id="email"/>
        <Field label="Change Password" id="email" mutlipleFields={["previousPassword", "newPassword", "newPasswordAgain"]}/>

        <h1 class="w-full text-center text-[3em] mt-4">Profile</h1>
        
        <p class="text-[2em]">About Me</p>
        <textarea id="aboutMe" class="w-full h-[8em] rounded-xl text-[2em] p-[0.2em] ml-auto"></textarea>
        <button class="rounded-2xl bg-interactable text-[#FFFFFF] w-full] text-[2em] mt-auto">Change</button>
    
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