<script lang="ts">
    import Topbar from "../../components/topbar.svelte";
</script>

<Topbar/>

<div class="flex items-center justify-center min-h-screen">
    <div class="bg-item w-[400px] h-[600px] p-8 rounded-3xl shadow-2xl flex flex-col">
        <h1 class="w-[100%] text-center text-4xl">Create an Account</h1>
        
        <p class="text-[1.25em] mt-8 mb-1">Email</p>
        
        <input type="text" id="email" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <p class="text-[1.25em] mt-2 mb-1">Username</p>
        
        <input type="text" id="username" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <p class="text-[1.25em] mt-2 mb-1">Password</p>
        
        <input type="password" id="password" class="w-[100%] h-[1.5em] rounded-xl text-[2em] p-[0.2em]">

        <button class="rounded-2xl bg-interactable text-[#FFFFFF] w-[100%] text-2xl p-2 mt-auto">Signup</button>
    </div>
</div>