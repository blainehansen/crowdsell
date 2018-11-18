<template lang="pug">

.custom-select
	.option-box(
		v-for="(option, index) in internalOptions",
		:key="option.value",
		:class="{ selected: option.selected }",
		@click="handleClick(index)",
	)
		slot(:option="option")
			span.option-text {{ option.name }}

</template>

<script>

export default {
	name: 'CustomSelect',

	props: {
		options: Array,
		multiple: {
			type: Boolean,
			default: false,
		}
	},

	data() {
		value: {
			type: [Array, String],
			validator(value) {
				if (Array.isArray(value) && !value.every(v => typeof v === 'string')) return false
				else return true
			}
		},
	},

	computed: {
		internalOptions() {
			const evaluator = this.getEvaluator()
			return this.options.map(option => ({ selected: evaluator(option), ...option }))
		}
	},

	methods: {
		getEvaluator() {
			const value = this.value

			return typeof value === 'string'
				? option => option.value === value
				: option => value.includes(option.value)
		},

		handleClick(index) {
			const internalOptions = this.internalOptions
			const option = internalOptions[index]

			option.selected = !option.selected

			this.$emit('input', internalOptions.filter(o => o.selected).map(o => o.value))
		},
	},
}

</script>
