<script lang="ts">
  import { Router, links, Route, Link } from "svelte-navigator";

  /* Routes */
  import Songs from "./pages/Songs.svelte";
  import Albums from "./pages/Albums.svelte";
  import Artists from "./pages/Artists.svelte";
  import Add from "./pages/Add.svelte";
  import SoloAlbum from "./pages/SoloAlbum.svelte";
  import SoloArtist from "./pages/SoloArtist.svelte";

  export let url = "/Songs";
  console.log(url);

  
</script>

<main>
  <Router url="{url}">
    <div class="navbar" use:links>
      <h1>Music player</h1>
      <ul class="pages">
        <div>• Categories :</div>
        <li class="wrapper"><Link to="/Songs" replace class="my-link"><span class="link-content">Songs</span></Link></li>
        <li class="wrapper"><Link to="/Albums" replace class="my-link">Albums</Link></li>
        <li class="wrapper"><Link to="/Artists" replace class="my-link">Artists</Link></li>
        <div>• Tools :</div>
        <li class="wrapper"><Link to="/Add" replace class="my-link">Add Entry</Link></li>
      </ul>
    </div>
    <div class="page">
      <Route path="*"><Songs/></Route>
      <Route path="/Songs"><Songs/></Route>
      <Route path="/Albums"><Albums/></Route>
      <Route path="/Albums/:id" component="{SoloAlbum}"/>
      <Route path="/Artists"><Artists/></Route>
      <Route path="/Artists/:id" component="{SoloArtist}"/>
      <Route path="/Add"><Add/></Route>
    </div>
  </Router>
</main>


<style lang="scss">

  main{
    display: flex;
    flex-direction: row;
    align-items: start;
    justify-content: start;
    height: 100vh;
    width: 100vw;

    .navbar{

      position: fixed; 
      top: 0;
      overflow: hidden;

      flex-direction: column;
      align-items: start;
      justify-content: start;
      
      height: 100%;
      width: 20vw;

      background-color: var(--main-color-dark);
      box-shadow: 0 0 10px rgba(0,0,0,0.2);

      h1{
        font-size: 1.5rem;
        font-weight: 600;
        margin: 0;
        padding: 1rem;
      }

      .pages{
        display: flex;
        flex-direction: column;
        align-items: start;
        justify-content: start;
        height: 100%;
        width: 100%;
        padding: 1rem;

        animation: 0.5s;

        div{
          font-size: 1.2rem;
          font-weight: 600;
          margin: 0.5rem;
        }

        li{
          font-size: 1.2rem;
          font-weight: 400;
          margin: 0.5rem;
          padding: 0.5rem;
          list-style: none;

          background-color: var(--main-color-light);
          border-radius: 5px;

          cursor: pointer;

          transition: 0.2s;

          &:hover{
            transform: scale(1.1);
          }

          
        }
      }
    }

    .wrapper :global(.my-link) {
      all: unset;
      color: var(--main-color-lighter);
    }

    .page{
      display: flex;
      align-items: start;
      justify-content: center;
      margin-left: 20vw;
      width: 80%;
      height: 100%;
    }
  }

</style>
