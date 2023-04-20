<script>
	import { audioData } from './audioData.js';
	
	let listIsShowing = false;
	
	const showPlayList = () => listIsShowing = !listIsShowing
</script>


<section id="playlist-cont">
	<ul class:show-list={listIsShowing}>
		{#each audioData as {name, url}, i}
		<!-- svelte-ignore a11y-no-visible -->
		<li data-track-id={i}
				on:click>
			{name}</li>
		{/each}
	</ul>
	<button class="accordion"
		class:active={listIsShowing}
		on:click={showPlayList}>&#9776; Playlist</button>
</section>


<style>
	section#playlist-cont {
		width: 298px;	
		box-shadow: 0 0 5px 0 rgba(0,0,0,0.5);
		border-radius: 0 0 5px 5px;
		background: var(--main-color-darker);
	}
	
	/* Style the buttons that are used to open and close the accordion panel */
	button.accordion {
		background: var(--main-color-darker);
		color: var(--main-color-lighter);
		margin: 0;
		padding: 5px 15px;
		width: 100%;
		border: none;
		outline: none;
		cursor: pointer;
		transition: 0.4s;
	}

	/* Add a background color to the button if it is clicked on (add the .active class with JS), and when you move the mouse over it (hover) */
	button.active {
		background-color: var(--main-color-darker);
		color: var(--main-color-lighter);
	}
	
	ul {
		list-style-type: none;
		background: var(--main-color-darker);
		color: var(--main-color-lighter);
		margin: 0;
		padding: 0;
		border-radius: 0 0 5px 5px;
		height: 0;
		overflow: auto;
		transition: height .3s;
	}
	
	li {
		line-height: 150%;
		padding: 5px 5px 5px 20px;
		border-bottom: 1px solid var(--main-color-lighter);
		cursor: pointer;
	}
	
	li:active, li:hover {
		background-color: var(--main-color-dark);
	}
	
	.show-list {
		height:50px;
	}
</style>