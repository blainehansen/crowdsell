<template lang="pug">

.create-project
	p Create a project
	p
		button(:disabled="!project$anyTouched", @click="saveProject") create this project

	Overall(:createMode="true")

</template>

<script>
import { get } from 'vuex-pathify'

import Overall from './edit/Overall'

export default {
	name: 'CreateProject',

	components: {
		Overall,
	},

	computed: {
		project$anyTouched: get('project/$anyTouched'),
	},

	methods: {
		async saveProject() {
			const projectId = await this.$store.dispatch('project/saveProject')

			this.$router.replace({ name: 'projectEditOverall', params: { projectId } })
		},
	}
}
</script>

<style lang="sass">
</style>
