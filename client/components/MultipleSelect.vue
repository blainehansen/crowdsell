<template lang="pug">

.multiple-select(v-click-outside="{ handler: close, active: popover }")
	.select-option(v-for="option in internalOptions", :key="option.value")
		slot(:option="option", :clickOption="clickOption", :yesOption="yesOption", :noOption="noOption")
			.option(@click="clickOption(option)", :class="{ selected: option.selected }") {{ option.name }}

</template>

<script>

import { optionsValidator } from './componentUtils'

export default {
	name: 'MultipleSelect',

	props: {
		value: {
			type: Array,
			validator: (value) => value.every(v => typeof v === 'string')
		},
		options: {
			type: Array,
			validator: optionsValidator,
		},

		popover: {
			type: Boolean,
			default: false,
		},
	},

	data() {
		return {
			isOpen: false,
		}
	},

	computed: {
		internalOptions() {
			const value = this.value
			return this.options.map(option => ({ selected: value.includes(option.value), ...option }))
		},
	},

	methods: {
		close() {
			this.isOpen = false
		},

		makeEmitValue(option) {
			return this.value.filter(v => v !== option.value)
		},

		clickOption(option) {
			const emitValue = this.makeEmitValue(option)
			if (emitValue.length === this.value.length)
				emitValue.push(option.value)

			this.$emit('input', emitValue)
		},

		yesOption(option) {
			const emitValue = this.makeEmitValue(option)
			emitValue.push(option.value)

			this.$emit('input', emitValue)
		},

		noOption(option) {
			const emitValue = this.makeEmitValue(option)
			this.$emit('input', emitValue)
		},
	},
}

</script>


<style lang="sass">

.multiple-select
	.select-option
		.option
			&.selected
				background-color: blue

</style>
