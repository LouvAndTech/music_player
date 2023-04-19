<script lang='ts'>
    import { Link } from "svelte-navigator";
    import { useParams } from "svelte-navigator";
    import { GetAlbums, GetSongs, GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    const params = useParams();
    let id = $params.id;
    console.log(id);

    let artist: main.Artist = {id: 0, firstName: "", lastName: ""};
    let albums: main.Album[][] = [[]];
    
    console.log("Init")

    try{
        let param : main.Param[] = [{key: "person_id", value: id}]
        GetArtists(param).then((result) => {
            console.log("Artist result :");
            console.log(result);
            if (result != null){
                artist = result[0];
            }
        });

        param = [{key: "artist_id", value: id}]
        GetAlbums(param).then((result) => {
            console.log("Album result :");
            console.log(result);
            if (result != null){
                albums = [];
                while(result.length) albums.push(result.splice(0,3));
            }
        });

    }catch(e){
        console.log(e);
    }
</script>

<main>
    <div class="container">
        {#await artist}
                <p>Chargement...</p>
        {:then album}
            <img src="https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTASYJZc6NswNmFLXxgg6rM2J7_yMRqkayqugwAA1T6a2HCGCuo" alt="">
            <div class="info">
                <h1>{artist.firstName}
                {#if artist.lastName != ""}
                    {artist.lastName}
                {/if}
                </h1>
                
            </div>
        {/await}
        <div class='grid'>
            {#each albums as row , x}
              <div class='row'>
                {#each row as album , y}
                  <div class='element wrapper'>
                    <Link to="/Albums/{album.id}" class="my-link">
                        <img src="https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTASYJZc6NswNmFLXxgg6rM2J7_yMRqkayqugwAA1T6a2HCGCuo" alt="">
                        <p>{album.name}</p>
                    </Link>
                  </div>
                {/each}
              </div>
            {/each}
        </div>
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

                border-radius: 50%;
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

            .grid{
                grid-column-start: 1;
                grid-column-end: 7;
                grid-row: 2;

                width: 100%;
                height: 100%;

                display: flex;
                flex-direction: column;
                align-items: start;
                justify-content: space-evenly;

                scroll-behavior: smooth;
                overflow-y: scroll;

                .row{
                    display: flex;
                    flex-direction: row;
                    justify-content: space-evenly;
                    align-items: center;

                    .element{
                        display: flex;
                        flex-direction: column;
                        align-items: center;
                        justify-content: center;
                        padding: 1.5rem;

                        img{
                            width: 200px;
                            height: 200px;
                            border-radius: 5px;
                            object-fit: cover;
                        }

                        transition: 0.2s;
                        cursor: pointer;
                        &:hover{
                            transform: scale(1.05);
                        }
                    }
                }
            }
        }
        .wrapper :global(.my-link) {
            all: unset;
            color: var(--main-color-lighter);
            cursor: pointer;
        }
    }

</style>