<template lang="pug">

#project-promises
	h1 Promises

	.promises(v-for="(promise, index) in promises")
		input(:value="promise", @input="event => catchInput(event, index)", placeholder="make a promise to your backers")
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
			this.promises = this.promises
		},
		deletePromise(index) {
			this.promises.splice(index, 1)
			this.promises = this.promises
		},
		catchInput(event, index) {
			this.promises[index] = event.target.value
			this.promises = this.promises
		}
	}
}
</script>

<style lang="sass">
</style>
