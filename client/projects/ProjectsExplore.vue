<template lang="pug">

#explore

	#input-area
		input(v-model="internalQuery", placeholder="search")

		MultipleSelect(v-model="fakeValue", :options="['one', 'two', 'three'].map(v => ({ name: v, value: v }))")

		p {{ fakeValue }}

		select(:value="categories", @input="changeCategories", multiple)
			option(v-for="option in categoryOptions", :value="option.value") {{ option.name }}

		select(:value="tags", @change="changeTags", multiple)
			option(v-for="tag in tagOptions", :value="tag.value") {{ tag.name }}


	p(v-if="loading")

	#found-projects
		.project(v-for="project in foundProjects")
			p {{ project.title }}

</template>

<script>

import titleCase from 'title-case'
import { delay } from '@/api'

const cannedProjects = [{
	title: 'Some Project',
	category: 'COMPUTER_SOFTWARE',
	tag: 'A_TAG',
}, {
	title: 'Some Different',
	category: 'COMPUTER_HARDWARE',
	tag: 'B_TAG',
}]

const categoryOptions = [
	'COMPUTER_SOFTWARE',
	'COMPUTER_HARDWARE',
].map(o => ({ value: o, name: titleCase(o) }))

const tagOptions = [
	'A_TAG',
	'B_TAG',
].map(t => ({ value: t, name: titleCase(t) }))

export default {
	name: 'ProjectsExplore',

	props: {
		query: {
			type: String,
			default: '',
		},
		categories: {
			type: Array,
			default: () => [],
		},
		tags: {
			type: Array,
			default: () => [],
		}
	},

	data() {
		return {
			fakeValue: [],

			internalQuery: '',

			categoryOptions,
			tagOptions,
		}
	},

	created() {
		this.internalQuery = this.query
	},

	asyncComputed: {
		foundProjects: {
			get() {
				const { query, categories, tags } = this
				const projects = cannedProjects
					.filter(project => (!categories.length && !tags.length) || categories.includes(project.category) || tags.includes(project.tag))
					.filter(project => !query.length || project.title.toLowerCase().includes(query.toLowerCase()))

				return delay(projects)
			},
			eager: true,
			watch: 'query',
			watchClosely() {
				return { categories: this.categories, tags: this.tags, }
			},
		},
	},

	computed: {
		loading() {
			return this.foundProjects$loading || this.foundProjects$pending
		},
	},

	watch: {
		internalQuery(newQuery) {
			this.switchParams(this.categories, this.tags, newQuery)
		}
	},

	methods: {
		changeCategories(event) {
			const selectedOptions = event.target.selectedOptions
			const newCategories = []
			for (let index = selectedOptions.length - 1; index >= 0; index--) {
				newCategories.push(selectedOptions[index].value)
			}
			console.log(newCategories)

			this.switchParams(newCategories, this.tags, this.query)
		},

		changeTags(newTags) {
			this.switchParams(this.categories, newTags, this.query)
		},

		switchParams(categories, tags, query) {
			this.$router.push({
				name: this.$route.name,
				query: {
					q: query.length ? query : undefined,
					c: categories.length ? categories : undefined,
					t: tags.length ? tags : undefined,
				}
			})
		},
	},
}

</script>

<style lang="sass">
</style>
