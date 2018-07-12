<template lang="pug">

.project-overall
	h1 Overall

	input(v-model="name", placeholder="give your project a name")
	input(v-model="description", placeholder="give your project a description")

</template>

<script>
import { sync } from '@/packages/vuex-pathify'

export default {
	name: 'Overall',
	computed: {
		...sync('project/', [
			'name',
			'description',
		])
	}
}
</script>

<style lang="sass">
</style>
