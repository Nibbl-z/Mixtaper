<script lang="ts">
    import Topbar from "../components/topbar.svelte";
    import Level from "../components/level.svelte"
</script>

<Topbar/>

<div class="flex items-center flex-col justify-center my-10">
    <img src="/logo.png" alt="Logo" class="w-[30%] my-5">
    <h1 class="xl:text-5xl text-2xl font-Itim">Custom Levels and Mixtapes for Bits and Bops</h1>
</div>

<div class="flex items-center justify-center flex-col">
    <div class="w-full 2cols:max-w-[100em] max-w-[50em] flex justify-between items-center">
        <h1 class="text-left text-3xl">Recently Uploaded</h1>
        <a href="/" class="text-right text-3xl">Show More &gt;</a>
    </div>
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8 my-4">
        <Level songName="Awesome Sauce" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="Awesome Sauceaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="I Hate CSS" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="we should rewrite the web in lua tbh" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
    </div>    
</div>





<a href="https://www.flaticon.com/free-icons/search" title="search icons">Search icons created by Chanut - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/user" title="user icons">User icons created by Freepik - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/upload" title="upload icons">Upload/Download icons created by Ilham Fitrotul Hayat - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/song" title="song icons">Song icons created by Freepik - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/heart" title="heart icons">Heart icons created by Chanut - Flaticon</a>

<a href="https://www.flaticon.com/free-icons/rythm" title="rythm icons">Rythm icons created by Freepik - Flaticon</a>
\
<a href="https://www.flaticon.com/free-icons/clock" title="clock icons">Clock icons created by Ilham Fitrotul Hayat - Flaticon</a>