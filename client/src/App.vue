<script setup lang="ts">
import Arrow from './components/sections/Arrow.vue';
// import Header from './components/sections/Header.vue';
import Invite from './components/sections/Invite.vue';
import LoveStory from './components/sections/LoveStory.vue';
import Map from './components/sections/Map.vue';
import Questions from './components/sections/Questions.vue';
import Schedule from './components/sections/Schedule.vue';
import Poll from './components/sections/Poll.vue';
import {onMounted} from "vue";
import axios from "axios";

onMounted(async () => {
  await axios.get(import.meta.env.VITE_API_URL + '/check')
})

</script>

<template>

<!--	<Header />-->

	<Invite />

	<Questions />

	<Poll />

	<LoveStory />

	<Schedule />

	<Map />

	<Arrow />

</template>

<style scoped>
#section-first .section__first-overlay {
	background: rgba(0, 0, 0, 0);
	-webkit-backdrop-filter:
		grayscale(0%) saturate(100%) contrast(100%) brightness(100%);
	backdrop-filter:
		grayscale(0%) saturate(100%) contrast(100%) brightness(100%);
}
</style>
