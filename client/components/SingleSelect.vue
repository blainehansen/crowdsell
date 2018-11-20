<template lang="pug">

.single-select(v-click-outside="close")
	.currently-selected(@click="isOpen = !isOpen")
		span {{ currentOption.name }}

	//- .options(v-displaying="isOpen")
	.options(v-if="isOpen")
		.option(
			v-for="option in options",
			@click="handleClick(option.value)",
			:class="{ selected: option.value === value }",
		)
			span {{ option.name }}

</template>


<script>

import { optionsValidator } from './componentUtils'

export default {
	name: 'SingleSelect',

	props: {
		value: {
			type: String,
			required: null,
			validator: (value) => value !== undefined && (value === null || value.length)
		},
		options: {
			type: Array,
			validator: optionsValidator
		}
	},

	data() {
		return {
			isOpen: false,
		}
	},

	computed: {
		currentOption() {
			const value = this.value
			return this.options.find(option => option.value === value)
		}
	},

	methods: {
		close() {
			this.isOpen = false
		},

		handleClick(optionValue) {
			this.$emit('input', optionValue)
			this.close()
		},
	},
}

</script>


<style lang="sass" scoped>

.single-select
	display: inline-block
	border: 1px solid black
	position: relative

	.options
		position: absolute
		display: inline-block

		top: 100%
		left: 0
		right: 0

		.option
			&.selected
				background-color: blue

</style>
