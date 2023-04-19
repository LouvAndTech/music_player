<script lang='ts'>
    import { Link } from "svelte-navigator";
    import { GetArtists} from "../../wailsjs/go/main/App";
    import type { main } from "../../wailsjs/go/models";

    let artists: main.Artist[][] = [[]];
    let param : main.Param[] = []

    try{
        GetArtists(param).then((result) => {
            console.log("Artist result :");
            console.log(result);
            if (result != null){
                artists = [];
                while(result.length) artists.push(result.splice(0,3));
            }
        });
    }catch(e){
        console.log(e);
    }

</script>

<main>
    <div class='grid'>
        {#each artists as row , x}
          <div class='row'>
            {#each row as artist , y}
              <div class='element wrapper'>
                <Link to="/Artists/{artist.id}" class="my-link">
                    <img src="https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTASYJZc6NswNmFLXxgg6rM2J7_yMRqkayqugwAA1T6a2HCGCuo" alt="">
                    <p>{artist.firstName}
                    {#if artist.lastName}
                        {artist.lastName}
                    {/if}
                    </p>
                    
                </Link>
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
                        border-radius: 50%;
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

        .wrapper :global(.my-link) {
            all: unset;
            color: var(--main-color-lighter);
            cursor: pointer;
        }
    }

</style>