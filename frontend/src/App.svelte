<script lang="ts">
  import { Router, links, Route } from "svelte-routing";

  /* Routes */
  import Songs from "./pages/Songs.svelte";
  import Albums from "./pages/Albums.svelte";
  import Artists from "./pages/Artists.svelte";
  import Add from "./pages/Add.svelte";

  export let url = "Songs";
  
</script>

<main>
  <Router url="{url}">
    <div class="navbar" use:links>
      <h1>Music player</h1>
      <ul class="page">
        <div>• Categories :</div>
        <li><a href="Songs">Songs</a></li>
        <li><a href="Albums">Albums</a></li>
        <li><a href="Artists">Artists</a></li>
        <div>• Tools :</div>
        <li><a href="Add">Add Entry</a></li>
      </ul>
    </div>
    <div class="page">
      <Route path="Songs"><Songs/></Route>
      <Route path="Albums"><Albums/></Route>
      <Route path="Artists"><Artists/></Route>
      <Route path="Add"><Add/></Route>
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
      display: flex;
      flex-direction: column;
      align-items: start;
      justify-content: start;
      height: 100%;
      width: 20%;
      background-color: var(--main-color-dark);
      box-shadow: 0 0 10px rgba(0,0,0,0.2);

      h1{
        font-size: 1.5rem;
        font-weight: 600;
        margin: 0;
        padding: 1rem;
      }

      .page{
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

          a{
            text-decoration: none;
            color: currentColor;
            width: 100%;
            height: 100%;
          }
        }
      }
    }

    .page{
      display: flex;
      align-items: start;
      justify-content: center;
      width: 80%;
      height: 100%;
    }
  }

</style>
