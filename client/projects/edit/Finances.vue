<template lang="pug">

#user-finances
	h1 Finances

	input(v-model="cardNumber", placeholder="0000 0000 0000 0000")

</template>

<script>
export default {
	name: 'Overall',

	data() {
		return {
			cardNumber: ''
		}
	},

	// asyncData: {
	// 	cards
	// }

	// methods: {
	// 	addCard() {

	// 	},
	// 	deleteCard() {

	// 	},
	// 	editCard() {

	// 	}
	// }
}
</script>

<style lang="sass">
</style>
