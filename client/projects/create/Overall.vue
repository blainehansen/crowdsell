<template lang="pug">

.project-overall
	h1 Overall

	input(v-model="name", placeholder="give your project a name")

</template>

<script>
import { sync } from 'vuex-pathify'

export default {
	name: 'Overall',
	computed: {
		...sync('project/', [
			'name'
		])
	}
}
</script>

<style lang="sass">
</style>
