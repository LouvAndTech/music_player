<script lang="js">
	import { audioData } from './audioData.js';
	
	import TrackHeading from './TrackHeading.svelte';
	import ProgressBarTime from './ProgressBarTime.svelte';
	import Controls from './Controls.svelte';
	import VolumeSlider from './VolumeSlider.svelte';
	import PlayList from './PlayList.svelte';
	
	// Get Audio track
	let trackIndex = 0;
	// $: console.log(trackIndex)
	let audioFile = new Audio(audioData[trackIndex].url);
	let trackTitle = audioData[trackIndex].name;
	
	const loadTrack = () => {
		audioFile = new Audio(audioData[trackIndex].url);
		audioFile.onloadedmetadata = () => {
			totalTrackTime = audioFile.duration;
			updateTime();
		}
		trackTitle = audioData[trackIndex].name;
	}
	
	const autoPlayNextTrack = () => {
		if (trackIndex <= audioData.length-1) {
			trackIndex += 1;
			loadTrack();
			audioFile.play();
		} else {
			trackIndex = 0;
			loadTrack();
			audioFile.play();
		}
	}
	
	
	// Track Duration and Progress Bar
	let totalTrackTime;
	$: console.log(totalTrackTime)
	audioFile.onloadedmetadata = () => {
		totalTrackTime = audioFile.duration;
		updateTime();
	}
	
	let totalTimeDisplay = "loading...";
	let currTimeDisplay = "0:00:00";
	let progress = 0;
	let trackTimer;
	
	function updateTime() {
		progress = audioFile.currentTime * (100 / totalTrackTime);
		
		let currHrs = Math.floor((audioFile.currentTime / 60) / 60);
		let currMins = Math.floor(audioFile.currentTime / 60);
		let currSecs = Math.floor(audioFile.currentTime - currMins * 60);
		
		let durHrs = Math.floor( (totalTrackTime / 60) / 60 );
		let durMins = Math.floor( (totalTrackTime / 60) % 60 );
		let durSecs =  Math.floor(totalTrackTime - (durHrs*60*60) - (durMins * 60));
		
		currTimeDisplay  = ((currHrs > 0 )? `${currHrs}:`: '' )+`${currMins}:`+ ((currSecs < 10) ? `0${currSecs}` : `${currSecs}`);
		totalTimeDisplay = ((durHrs > 0 ) ? `${durHrs}:` : '' )+ `${durMins}:`+ ((durSecs < 10) ? `0${durSecs}` : `${durSecs}`);
		
		if (audioFile.ended) {
			toggleTimeRunning();
		}
	}
	
	const toggleTimeRunning = () => {
		if (audioFile.ended) {
			isPlaying = false;
			clearInterval(trackTimer);
			console.log(`Ended = ${audioFile.ended}`);	
		} else {
			trackTimer = setInterval(updateTime, 100);
		}
	}
	

	// Controls
	let isPlaying = false;
	$: console.log(`isPlaying = ${isPlaying}`)
	
	const playPauseAudio = () => {
		if (audioFile.paused) {
			toggleTimeRunning()
			audioFile.play();
			isPlaying = true;
		} else {
			toggleTimeRunning()
			audioFile.pause();
			isPlaying = false;
		}	 	
	}
	
	const rewindAudio = () => audioFile.currentTime -= 10;
	const forwardAudio = () => audioFile.currentTime += 10;
	
	// Volume Slider
	let vol = 50;
	const adjustVol = () => audioFile.volume = vol / 100; 
	
	
	// Playlist
	const handleTrack = (e) => {
		if (!isPlaying) {
			trackIndex = Number(e.target.dataset.trackId);
			loadTrack();
			playPauseAudio(); // auto play
		} else {
			isPlaying = false;
			audioFile.pause();
			trackIndex = Number(e.target.dataset.trackId);
			loadTrack();
			playPauseAudio(); // auto play
		}
	}
</script>


<main>
	<div class="container">
		<div class="sub controls">
			<Controls {isPlaying} 
				on:rewind={rewindAudio}
				on:playPause={playPauseAudio}
				on:forward={forwardAudio} />
		
			<VolumeSlider bind:vol
				on:input={adjustVol} />	
		</div>	
		<div class="sub infos">
			<TrackHeading {trackTitle} />
			<ProgressBarTime {currTimeDisplay}
				{totalTimeDisplay}
				{progress} />
		</div>	
		<!--div class="sub playlist">
			<PlayList on:click={handleTrack} />
		</div-->
	</div>
</main>


<style lang="scss">
	main {
		.container{
			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: space-between;

			.sub{
				display: flex;
				flex-direction: column;
				align-items: center;

				margin: 15px;

				&.controls{
					justify-content: space-between;
				}

				&.infos{
					justify-content: space-between;
				}

				&.playlist{
					justify-content: space-between;
				}
			}
		}
	}
</style>