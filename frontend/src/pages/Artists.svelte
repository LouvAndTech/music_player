<script lang='ts'>
    import { Greet } from "../../wailsjs/go/main/App";
    import { main } from "../../wailsjs/go/models";

    let txt :string = "Hello World";
    let i :number = 0;

    function generate() {
        let person = new main.Person();
        person.name = "Peter";
        person.age = i;
        i++;
        Greet(person).then((result) => {
            console.log(result);
            txt = result;
        });
    }

</script>

<main>
    <h1>all Artist</h1>
    <p>{txt}</p>
    <button on:click={generate}>Generate</button>
</main>

<style lang='scss'>

</style>