<script lang='ts'>
    import { onMount } from 'svelte';
    import { Link } from 'svelte-routing';
    import { GetAlbums, GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    let albums: main.Album[] = [];
    let artists: main.Artist[] = [];
    let param : main.Param[] = []

    onMount(async () => {
        try{
            GetArtists(param).then((result) => {
                artists = result;
            });
            GetAlbums(param).then((result) => {
                albums = result;
            });
        }
        catch(e){
            console.log(e);
        }
	}); 

    let title = "";
    let artist = "";
    let album = "";

    function submit(){
        console.log("submit");
        console.log("title :" +title + " artist :" + artist + " album :" + album);

    }

</script>

<main>
    <form>
        <label for="title">Title</label>
        <input type="text" name="title" id="title" bind:value={title}>
        <label for="artist">Artist</label>
        <select name="artist" id="artist" bind:value={artist}>
            <option value="">None</option>
            {#each artists as artist}
                <option value="{artist.id}">{artist.firstName}</option>
            {/each}
        </select>
        <label for="album">Album</label>
        <select name="album" id="album" bind:value={artist}>
            <option value="">None</option>
            {#each albums as album}
                <option value="{album.id}">{album.name}</option>
            {/each}
        </select>
        <input type="button" on:click={submit} value="Submit"/>
    </form>
</main>

<style lang='scss'>

</style>