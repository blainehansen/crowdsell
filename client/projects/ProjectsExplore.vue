<template lang="pug">

#explore

	#input-area
		input(v-model="internalQuery", placeholder="search")

		select(:value="categories", @input="changeCategories", multiple)
			option(v-for="option in categoryOptions", :value="option.value") {{ option.name }}

		select(:value="tags", @change="changeTags", multiple)
			option(v-for="tag in tagOptions", :value="tag.value") {{ tag.name }}


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
			internalQuery: '',

			categoryOptions,
			tagOptions,
		}
	},

	created() {
		this.internalQuery = this.query
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
					categories: categories.length ? categories : undefined,
					tags: tags.length ? tags : undefined,
					query: query.length ? query : undefined,
				}
			})
		},
	},

	asyncComputed: {
		foundProjects: {
			get() {
				const projects = cannedProjects
					.filter(project => this.categories.includes(project.category) || this.tags.includes(project.tag) || (!this.categories.length && !this.tags.length))
					.filter(project => !this.query.length || project.title.toLowerCase().includes(this.query.toLowerCase()))

				return delay(projects)
			},
			eager: true,
			watch: 'query',
			watchClosely() {
				return { categories: this.categories, tags: this.tags, }
			},
		},
	},
}

</script>

<style lang="sass">
</style>
