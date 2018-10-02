<template lang="pug">

#project-promises
	h1 Promises

	.promises(v-for="(index, promise) in promises")
		input(v-model="promise", placeholder="make a promise to your backers")
		button(@click="deletePromise(index)") delete promise

	button(@click="addPromise") add promise

</template>

<script>
import { sync } from 'vuex-pathify'

export default {
	name: 'Overall',

	computed: {
		...sync('project/', [
			'promises',
		])
	},

	methods: {
		addPromise() {
			this.promises.push('')
		},
		deletePromise(index) {
			this.promises.splice(index, 1)
		},
	}
}
</script>

<style lang="sass">
</style>
