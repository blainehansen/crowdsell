<template lang="pug">

#home
	h1 Welcome

	MdEditor(v-model="myText")

	//- button(@click="sendCardInfo") send
	//- p(v-if="cardInfoSuccess") {{ cardInfoSuccess }}
	//- p(v-if="cardInfoError") {{ cardInfoError }}

</template>

<script>
import { privateApi } from '@/api'

import MdEditor from '@/components/MdEditor'

export default {
	name: 'home',

	components: {
		MdEditor
	},

	data() {
		return { myText: '## hello' }
	},

	// data() {
	// 	return {
	// 		cardInfoSuccess: null,
	// 		cardInfoError: null,
	// 	}
	// },
	// methods: {
	// 	async sendCardInfo() {
	// 		const { data: cardToken } = await privateApi.generateCardToken()
	// 		console.log(cardToken)

	// 		promisepay.createCardAccount(cardToken, {
	// 			full_name: "Bella Buyer",
	// 			number: "4111111111111111",
	// 			expiry_month: "02",
	// 			expiry_year: "2022",
	// 			cvv: "123"
	// 		}, (data) => {
	// 			this.cardInfoSuccess = data
	// 		}, (data) => {
	// 			this.cardInfoError = data
	// 		})
	// 	}
	// }
}
</script>

<style lang="sass">
</style>
