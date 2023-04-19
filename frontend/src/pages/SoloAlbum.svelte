<script lang='ts'>
    import { useParams } from "svelte-navigator";
    import { GetAlbums, GetSongs, GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    const params = useParams();
    let id = $params.id;
    console.log(id);

    let songs: main.Song[] = [];
    let album: main.Album = {id: 0, name: "", artist: 0, imagePath: ""};
    let artist: main.Artist = {id: 0, firstName: "", lastName: ""};
    let param : main.Param[] = [{key: "album_id", value: id}]

    console.log("Init")

    try{
        GetAlbums(param).then((result) => {
            console.log("Album result :");
            console.log(result);
            if (result != null){
                album = result[0];
            }
            console.log("Albums :")
            console.log(album);
            console.log("Album name :")
            console.log(album.name);
            console.log("typeof album :")
            console.log(typeof album);

            GetSongs(param).then((result) => {
                console.log("Songs result :");
                console.log(result);
                if (result != null){
                    songs = result;
                }
                console.log("Songs :")
                console.log(songs);
            });

            param = [{key: "person_id", value: album.artist}]
            GetArtists(param).then((result) => {
                console.log("Artist result :");
                console.log(result);
                if (result != null){
                    artist = result[0];
                }
                console.log("Artist :")
                console.log(artist);
                console.log("Artist name :")
                console.log(artist.firstName);
                console.log("typeof artist :")
                console.log(typeof artist);
            });
            
        });
    }catch(e){
        console.log(e);
    }
</script>

<main>
    <div class="container">
        {#await album}
                <p>Chargement...</p>
        {:then album}
            <img src="https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTASYJZc6NswNmFLXxgg6rM2J7_yMRqkayqugwAA1T6a2HCGCuo" alt="">
            <div class="info">
                <h1>{album.name}</h1>
                {#await artist}
                    <p>Chargement...</p>
                {:then artist} 
                    <div class="artist">
                        <p>{artist.firstName}</p>
                        {#if artist.lastName != ""}
                            <p>{artist.lastName}</p>
                        {/if}
                    </div>
                {/await}
            </div>
        {/await}
        <ul class="list">

            {#await songs}
                <p>Chargement...</p>
            {:then songs} 
                {#each songs as song}
                    <li><p>{song.name}</p></li>
                {/each}
            {/await}
        </ul>
    </div>
</main>

<style lang='scss'>
    main{
        display: flex;
        flex-direction: column;
        align-items: start;
        justify-content: center;

        overflow: scroll;

        .container{
            display: grid;
            grid-template-columns: repeat(6, 1fr);
            grid-template-rows: 0.5fr 2fr;

            padding: 2rem;

            img{
                grid-column-start: 1;
                grid-column-end: 3;
                grid-row: 1;

                border-radius: 5px;
            }

            .info{
                grid-column-start: 4;
                grid-column-end: 6;
                grid-row: 1;

                padding: 2rem;

                display: flex;
                flex-direction: column;
                align-items: start;
                .artist{
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                    justify-content: space-between;

                    padding-right: 2rem;

                    *{
                        margin-right: 1rem;
                    }
                }
            }

            .list{
                grid-column-start: 2;
                grid-column-end: 7;
                grid-row: 2;

                width: 100%;
                height: 100%;

                display: flex;
                flex-direction: column;
                align-items: start;
                justify-content: start;

                margin: 0;
                margin-top: 2rem;
                padding: 0;

                

                li{
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                    justify-content: space-between;
                    margin-bottom: 1rem;
                    padding: 0.5rem;
                    width: 100%;
                    border-radius: 5px;
                    background-color: var(--main-color-light);

                    list-style-type: none;

                    color: var(--main-color-lighter);

                    transition: 0.2s;
                    cursor: pointer;
                    &:hover{
                        background-color: var(--main-color-dark);
                        transform: scale(1.05);
                    }
                }
            }
        }
    }

</style>