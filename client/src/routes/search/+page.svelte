<script lang="ts">
	import Topbar from "../../components/topbar.svelte"
    import Level from "../../components/level.svelte"

    export let query: string
    export let results: number
    query = "Awesome Levels"
    results = 100000
</script>

<Topbar/>

<div class="my-2 ml-8 p-[4rem]">
    <h1 class="text-4xl">Showing {results} results for {query}</h1>
</div>

<div class="flex items-center justify-center flex-col">
    <div class="width-[100em] grid grid-cols-1 2cols:grid-cols-2 justify-items-center gap-8">
        <Level songName="Awesome Sauce" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="Awesome Sauceaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="I Hate CSS" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
        <Level songName="we should rewrite the web in lua tbh" songArtist="Josh, obviously" cover="/PLACEHOLDER.png"/>
    </div>    
</div>