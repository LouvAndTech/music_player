<script lang='ts'>
    import { GetAlbums, GetSongs, GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    let songs: main.Song[][] = [[]];
    let param : main.Param[] = []

    try{
        GetSongs(param).then((result) => {
                console.log("Songs result :");
                console.log(result);
                if (result != null){
                    songs = [];
                    while(result.length) songs.push(result.splice(0,3));
                }
                console.log("Songs :")
                console.log(songs);
            });

    }catch(e){
        console.log(e);
    }
</script>

<main>
    <div class='grid'>
        {#each songs as row , x}
          <div class='row'>
            {#each row as song , y}
              <div class='element'>
                <p>{song.name}</p>
              </div>
            {/each}
          </div>
        {/each}
    </div>
</main>

<style lang='scss'>
    main{
        .grid{
            display: flex;
            flex-direction: column;
            align-items: start;
            justify-content: space-evenly;

            scroll-behavior: smooth;
            overflow-y: scroll;

            width: 100%;

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
                    margin: 1rem;

                    width: 10rem;
                    height: 10rem;

                    background-color: var(--main-color-light);
                    border-radius: 5px;

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