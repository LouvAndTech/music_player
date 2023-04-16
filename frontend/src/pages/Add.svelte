<script lang='ts'>
    import { onMount } from 'svelte';
    import { GetAlbums, GetArtists} from "../../wailsjs/go/main/App";
    import { PostSongs, PostAlbums, PostArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    let errorMessage : string = "";
    let successMessage : string = "";

    let albums: main.Album[] = [];
    let artists: main.Artist[] = [];
    let param : main.Param[] = []

    onMount(async () => {
        await updateAlbArt();
	}); 

    function updateAlbArt(){
        try{
            GetArtists(param).then((result) => {
                if (result == null)
                    artists = [];
                else
                    artists = result;
            });
            GetAlbums(param).then((result) => {
                if (result == null)
                    albums = [];
                else
                    albums = result;
            });
        }
        catch(e){
            console.log(e);
        }
    }

    /*New song*/
    let newSong : main.Song = {
        name: "",
        artist: -1,
        album: -1,
        path: "",
    }

   async function submitSong(){
        successMessage = "";
        newSong.artist = albums.find((album) => album.id == newSong.album)?.artist;
        if (newSong.album == -1)
            errorMessage = "Please select an album";
        else if (newSong.name == "")
            errorMessage = "Please add a name";
        else{
            PostSongs(newSong).then((result) => {
                if (result) {
                    successMessage = "Song added";
                    updateAlbArt();
                    newSong.artist = -1;
                    newSong.album = -1;
                    newSong.name = "";
                    newSong.path = "";
                }
            });
        }
    }

    /*New album*/
    let newAlbun : main.Album = {
        name: "",
        artist: -1,
        imagePath: "",
    }

    function submitAlbum(){
        successMessage = "";
        if(newAlbun.artist == -1)
            errorMessage = "Please select an artist";
        else if (newAlbun.name == "")
            errorMessage = "Please add a name";
        else{
            PostAlbums(newAlbun).then((result) => {
                if (result) {
                    successMessage = "Album added";
                    updateAlbArt();
                    newAlbun.artist = -1;
                    newAlbun.name = "";
                    newAlbun.imagePath = "";
                }
            });
        }
    }

    /*New artist*/
    let newArtist : main.Artist = {
        firstName: "",
        lastName: "",
    }

    async function submitArtist(){
        successMessage = "";
        if(newArtist.firstName == "")
            errorMessage = "Please add a name";
        else{
            PostArtists(newArtist).then((result) => {
                if (result) {
                    successMessage = "Artist added";
                    updateAlbArt();
                    newArtist.firstName = "";
                    newArtist.lastName = "";
                }
            });
        }
    }


</script>

<main>
    <h1>Add Entry</h1>
    {#if errorMessage != ""}
        <h3 class="error">{errorMessage}</h3>
    {:else if successMessage != ""}
        <h3 class="succes">{successMessage}</h3>
    {/if}

    <h2>Songs :</h2>
    <form>
        <div style="grid-column: 1;">
            <label for="title">Title</label>
            <input type="text" name="title" id="title" bind:value={newSong.name}>    
        </div>
        <div style="grid-column: 2;">
            <label for="album">Album</label>
            <select name="album" id="album" bind:value={newSong.album}>
                <option value="">None</option>
                {#each albums as album}
                    <option value="{album.id}">{album.name}</option>
                {/each}
            </select>    
        </div>
        <div style="grid-column: 3;">
            <input type="button" on:click={submitSong} value="Submit"/>
        </div>
    </form>
    <hr>
    <h2>Album :</h2>
    <form>
        <div style="grid-column: 1;">
            <label for="title">Name</label>
            <input type="text" name="title" id="title" bind:value={newAlbun.name}>    
        </div>
        <div style="grid-column: 2;">
            <label for="artist">Artist</label>
            <select name="artist" id="artist" bind:value={newAlbun.artist}>
                <option value="">None</option>
                {#each artists as artist}
                    <option value="{artist.id}">{artist.firstName}</option>
                {/each}
            </select>
        </div>
        <div style="grid-column: 3;">
            <input type="button" on:click={submitAlbum} value="Submit"/>
        </div>
    </form>
    <hr>
    <h2>Artist :</h2>
    <form>
        <div style="grid-column: 1;">
            <label for="firstName">First Name</label>
            <input type="text" name="firstName" id="firstName" bind:value={newArtist.firstName}>    
        </div>
        <div style="grid-column: 2;">
            <label for="lastName">Last Name</label>
            <input type="text" name="lastName" id="lastName" bind:value={newArtist.lastName}>    
        </div>
        <div style="grid-column: 3;">
            <input type="button" on:click={submitArtist} value="Submit"/>
        </div>
    </form>
    <hr>
</main>

<style lang='scss'>
    main {
        display: flex;
        flex-direction: column;
        align-items: ceter;
        justify-content: start;
        height: 100%;
        width: 100%;
        padding: 1rem;
        box-sizing: border-box;

        h1 {
            font-size: 2rem;
            margin-bottom: 1rem;
        }
        h2 {
            font-size: 1.5rem;
            margin-bottom: 1rem;
        }
        form {
            display: flex;
            flex-direction: row;
            align-items: center;
            justify-content: space-between;
            
            width: 100%;
            margin-bottom: 1rem;

            div {
                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: center;
                width: 100%;
                margin-bottom: 1rem;

                label {
                    font-size: 1rem;
                    margin-bottom: 0.5rem;
                }
                input {
                    font-size: 1rem;
                    margin-bottom: 0.5rem;
                    padding: 0.5rem;
                    border: 1px solid ;
                    border-radius: 0.5rem var(--main-color-darker);
                    box-sizing: border-box;

                    &[type="button"] {
                        font-size: 1rem;
                        margin-bottom: 0.5rem;
                        padding: 0.5rem;
                        border-radius: 0.5rem;
                        background-color: var(--main-color-light);
                        color: var(--main-color-lighter);
                        cursor: pointer;
                    }
                }
                select {
                    font-size: 1rem;
                    margin-bottom: 0.5rem;
                    padding: 0.5rem;
                    border: 1px solid #000;
                    border-radius: 0.5rem;
                    box-sizing: border-box;
                }
            }
        }

        hr{
            width: 100%;
            margin-bottom: 1rem;
            color: var(--main-color-lighter);
        }
        
    }
    .error {
        color: red;
    }
    .succes {
        color: green;
    }
</style>