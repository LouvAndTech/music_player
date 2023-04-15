<script lang='ts'>
    import { onMount } from 'svelte';
    import { Link } from 'svelte-routing';
    import { GetAlbums, GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    let albums: main.Album[] = [];
    let artists: main.Artist[] = [];
    let param : main.Param[] = []
    console.log(typeof param)

    /*onMount(async () => {
        console.log("onMount");
        GetArtists(param).then((result) => {
            console.log(result);
            artists = result;
        });
        GetAlbums(param).then((result) => {
            console.log(result);
            albums = result;
        });
	});*/
    
    function onclick() {
        console.log("onClick");
        try{
            GetArtists(param).then((result) => {
            console.log(result);
            artists = result;
            });
            GetAlbums(param).then((result) => {
                console.log(result);
                albums = result;
            });
        }
        catch(e){
            console.log(e);
        }
        
    }   

</script>

<main>
    <button on:click={onclick}>Click</button>
    <form action="">
        <label for="title">Title</label>
        <input type="text" name="title" id="title">
        <label for="artist">Artist</label>
        <select name="artist" id="artist">
            <option value="">None</option>
            {#each artists as artist}
                <option value="{artist.id}">{artist.firstName}</option>
            {/each}
        </select>
        <label for="album">Album</label>
        <select name="album" id="album">
            <option value="">None</option>
            {#each albums as album}
                <option value="{album.id}">{album.name}</option>
            {/each}
        </select>

    </form>
</main>

<style lang='scss'>

</style>