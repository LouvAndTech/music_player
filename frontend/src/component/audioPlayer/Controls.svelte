<script>
	import { createEventDispatcher } from 'svelte';
	
	let dispatch = createEventDispatcher();

	import play from '../../assets/images/media-playback-start.png';
	import pause from '../../assets/images/media-playback-pause.png';
	import rewind from '../../assets/images/media-skip-backward.png';
	import fwd from '../../assets/images/media-skip-forward.png';
	
  export let isPlaying;
	
</script>


<div id="btn-cont">
	<button id="rewind" on:click={() => dispatch('rewind')}>
		<img src={rewind} alt="rewind"/>
	</button>
	<button id="play" on:click={() => dispatch('playPause')}>
		<img src={isPlaying ? pause : play} 
				 alt="play"/>
	</button>
	<button id="forward" on:click={() => dispatch('forward')}>
		<img src={fwd} alt="forward"/>
	</button>
</div>

<style>
	#btn-cont {
		display: flex;
		justify-content: space-between;
	}
	
	button {
		width: 37px;
		border: none;
		font-size: 2.3rem;
		padding: 0;
		cursor: pointer;
		margin: 0;
		background-color: transparent;
	}
	
	button:active {
		font-size: 2.2rem;
		background-color: transparent;
	}
	
	button img {
		width: 100%;
	}
</style>