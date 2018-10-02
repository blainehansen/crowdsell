<template lang="pug">

#project-story
	//- MdEditor(v-model="story")
	textarea(v-model="story")

</template>


<script>
import { sync } from 'vuex-pathify'

export default {
	// components: {
	// 	MdEditor
	// },
	computed: {
		...sync('project/', [
			'story',
		])
	}
}
</script>


<style lang="sass">
</style>
