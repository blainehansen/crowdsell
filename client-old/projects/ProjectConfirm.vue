<template lang="pug">

#project-confirm
	template(v-if="stage === stages.WHICH")
		button(@click="submitGood") Everything looks good!

		button(@click="stage = stages.ALMOST") It's good, but I have some little misgivings.

		button(@click="stage = stages.BROKEN") They've broken some promises.


	template(v-if="stage === stages.BROKEN")
		p Say what's broken.

		p Is the license correct? (It should be {{ project.license }})
		select(v-model="wrongLicense")
			option(:value="null")
			option(:value="false") The license applied looks good
			option(:value="true") The license applied is wrong

		.promise(v-for="(promise, index) in project.promises", :key="promise.text")
			span {{ promise.text }}
			.checkbox(@click="toggleBroken(index)") {{ isBroken(index) ? 'broken' : 'fine' }}

		input(v-model="commentary", placeholder="general comments")

		button(:disabled="!canSubmitBroken", @click="submitBroken") Submit this feedback


	template(v-if="stage === stages.ALMOST")
		p Say what didn't quite make it for you.

		.promise(v-for="(promise, index) in project.promises", :key="project.id")
			span {{ promise.text }}
			.checkbox(@click="toggleAlmost(index)") {{ isAlmost(index) ? 'almost' : 'fine' }}

		input(v-model="commentary", placeholder="general comments")

		button(:disabled="!canSubmitAlmost", @click="submitAlmost") Submit this feedback


	template(v-if="stage === stages.DONE")
		p Alright! Thanks for giving feedback!

		router-link(:to="{ name: 'projectsExplore' }") Explore other projects

		router-link(:to="{ name: 'project', params: { userSlug, projectSlug } }") Head back to the project page

</template>

<script>

import { delay } from '@/api'

const stages = {
	WHICH: Symbol('WHICH'),
	ALMOST: Symbol('ALMOST'),
	BROKEN: Symbol('BROKEN'),
	DONE: Symbol('DONE'),
}

export default {
	name: 'ProjectsConfirm',

	props: {
		userSlug: String,
		projectSlug: String,
	},

	data() {
		return {
			stages,
			stage: stages.WHICH,

			brokenIndices: [],
			almostIndices: [],

			wrongLicense: null,
			commentary: "",
		}
	},

	asyncData: {
		project: {
			get() {
				console.log(this.userSlug, this.projectSlug)
				return delay({ license: 'MIT', promises: [{ text: "Promise one" }, { text: "Promise two" }] })
			},
			default: {},
		}
	},

	computed: {
		canSubmitAlmost() {
			return this.almostIndices.length > 0
				&& this.commentary.length > 0
		},
		canSubmitBroken() {
			return this.brokenIndices.length > 0
				&& this.commentary.length > 0
				&& this.wrongLicense !== null
		},
	},

	methods: {
		toggleIndex(index, list) {
			const foundIndex = list.indexOf(index)
			if (foundIndex === -1)
				list.push(index)
			else
				list.splice(foundIndex, 1)
		},
		isIndex(index, list) {
			return list.includes(index)
		},

		toggleBroken(index) {
			this.toggleIndex(index, this.brokenIndices)
		},
		isBroken(index) {
			return this.isIndex(index, this.brokenIndices)
		},

		toggleAlmost(index) {
			this.toggleIndex(index, this.almostIndices)
		},
		isAlmost(index) {
			return this.isIndex(index, this.almostIndices)
		},

		async submitGood() {
			console.log('all good')
			await delay()
			this.stage = stages.DONE
		},
		async submitBroken() {
			const promises = this.project.promises
			console.log({
				brokenPromises: this.brokenIndices.reduce((array, indexToAdd) => array.concat(promises[indexToAdd]), []),
				commentary: this.commentary,
				wrongLicense: this.wrongLicense
			})
			await delay()
			this.stage = stages.DONE
		},
		async submitAlmost() {
			const promises = this.project.promises
			console.log(promises)
			console.log({
				almostPromises: this.almostIndices.reduce((array, indexToAdd) => array.concat(promises[indexToAdd]), []),
				commentary: this.commentary
			})
			await delay()
			this.stage = stages.DONE
		},
	},
}

</script>

<style lang="sass">
</style>
