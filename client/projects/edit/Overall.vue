<template lang="pug">

#project-overall
	h1 Overall

	input(v-model="name", placeholder="give your project a name")
	input(v-model="description", placeholder="give your project a description")
	select(v-model="category")
		option(:value="null") Choose a Category
		option(v-for="option in PROJECT_CATEGORY_TYPE", :value="option.value") {{ option.text }}

	//- template(v-if="!createMode")

</template>

<script>
import { publicApi } from '@/api'
import { sync } from 'vuex-pathify'
import { PROJECT_CATEGORY_TYPE } from '@/constants'

export default {
	name: 'Overall',

	props: {
		createMode: {
			type: Boolean,
			default: false,
		}
	},

	data() {
		return {
			PROJECT_CATEGORY_TYPE,
		}
	},

	computed: {
		...sync('project/', [
			'name',
			'description',
			'category',
		])
	},
}
</script>

<style lang="sass">
</style>
