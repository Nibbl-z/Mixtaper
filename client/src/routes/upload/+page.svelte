<script lang="ts">
    import Topbar from "../../components/topbar.svelte"
    import Field from "./field.svelte";
</script>

<Topbar/>

<h1 class="w-full text-center text-[4em] my-10">Uploading Level</h1>

<div class="flex items-center justify-center mt-10">
    <div class="bg-item w-[50%] p-8 rounded-3xl shadow-2xl flex flex-col gap-4">
        <Field label="Song Name" id="songName" required/>
        <Field label="Song Artist" id="songArtist" required/>
        <Field label="Chart Name" id="chartName" />
        <Field label="Description" id="description" big/>
        <Field label="Cover Art" id="coverArt" file fileTitle="Drag and drop or select .png/.jpg file to upload" fileMaxSize={1}/>
        <Field label="YouTube Video" id="video" />
        <div class="my-8"></div>
        <Field id="riq" file fileTitle="Drag and drop or select .riq file to upload" fileMaxSize={20}/>

        <p class="w-full text-center text-2xl">Upload Status Message!</p>

        <div class="flex items-center justify-center">
            <button class="rounded-2xl bg-interactable text-[#FFFFFF] w-[20%] text-2xl p-2 mt-auto">Post</button>
        </div>
    </div>
</div>