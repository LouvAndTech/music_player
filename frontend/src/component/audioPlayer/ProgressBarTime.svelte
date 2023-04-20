<script>
	export let currTimeDisplay;
	export let totalTimeDisplay;
	export let progress;
</script>


<main>
	<div id="time">
		<span class="time" id="progress-time">{currTimeDisplay}</span>
			<div id="progress-bar-cont">
				<span id="bar" style="width: {progress}%"></span>
			</div>
		<span class="time" id="track-duration">{totalTimeDisplay}</span>
	</div>
</main>

<style>
	#time {
		width: 100%;
		display: flex;
		justify-content: space-between;
	}

	.time {
		margin-left: 10px;
		margin-right: 10px;
	}

	#progress-bar-cont {
		width: 15vw;
		height: 10px;
		margin: .5rem 0 .5rem;
		display: flex;
		align-items: center;
		border: 1px solid #888;
		cursor: pointer;
	}

	#bar {
		width: 0%;
		height: 10px;
		background-color: goldenrod;
	}
</style>