<template lang="pug">

#project-pledge
	template(v-if="stage === stages.AMOUNT")
		//- TODO make this a numeric
		input(v-model="amount", placeholder="0.00")

		button(v-for="suggestedAmount in [500, 250, 100, 50, 25, 10, 5]", @click="amount = suggestedAmount")

		button(@click="transitionAndClear(stages.ACCOUNT)") Yes that's how much

	template(v-else-if="stage === stages.ACCOUNT")
		template(v-if="paymentError")
			p Oh no! There was a problem with your payment.

			p Either
				button(@click="transitionAndClear(stages.AMOUNT)") change your amount
				| or change something about the account you use.

		template(v-if="user.accounts")
			p Use an existing payment method

			//- TODO add something to edit
			.payment-option(v-for="account in user.accounts")
				p {{ account.source }}
				p {{ account.number }}

				button(@click="useAccount(account)") use this {{ account.type }}

		template
			p(v-if="user.accounts") Or create a new one
			p(v-else) Create a payment method

			label(for="creating-bank") Creating a bank account?
			input#creating-bank(type="checkbox", v-model="creatingBank")

			template(v-if="creatingBank")
				input(v-model="routingNumber")
				input(v-model="accountNumber")

				label(for="business-account") Is this business or personal?
				input#business-account(type="checkbox", v-model="isBusinessAccount")

				label(for="savings-account") Is this savings?
				input#savings-account(type="checkbox", v-model="isSavingsAccount")


			template(v-else)
				//- TODO make this nice
				input(v-model="cardNumber", placeholder="0000 0000 0000 0000")

				//- TODO make these selects?
				input(v-model="expirationMonth", placeholder="January")
				input(v-model="expirationYear", placeholder="2019")

				input(v-model="cvv", placeholder="000")

			button(@click="createAccount") Create this payment method


	template(v-else-if="stage === stages.REVIEW")
		p You're about to make a payment of {{ amount }} with {{ account.number }}

		p Continue?

		button(@click="makePayment") Let's go!

	template(v-else-if="stage === stages.PAID")
		p Great!

		router-link(:to="{ name: 'projectsExplore' }") Explore other projects

		router-link(:to="{ name: 'project', params: { userSlug, projectSlug } }") Head back to the project page

	template(v-else)
		p zuh?

</template>

<script>

import { delay } from '@/api'

const stages = {
	AMOUNT: Symbol('AMOUNT'),
	ACCOUNT: Symbol('ACCOUNT'),
	REVIEW: Symbol('REVIEW'),
	PAID: Symbol('PAID'),
}

export default {
	name: 'ProjectsPledge',

	props: {
		userSlug: String,
		projectSlug: String,
	},

	data() {
		amount: 20.0,
		account: null,

		stages,
		stage: stages.AMOUNT,

		routingNumber: '',
		accountNumber: '',
		cardNumber: '',
		expirationMonth: null,
		expirationYear: null,
		cvv: '',

		creatingBank: false,
		isBusinessAccount: false,
		isSavingsAccount: false,

		paymentError: false,
	},

	methods: {
		useAccount(account) {
			// TODO do something to use account
			console.log(account)
			this.account = account
			this.transitionAndClear(stages.REVIEW)
		},

		async createAccount() {
			const account = this.creatingBank
				? {
					source: "The Iron Bank",
					number: `${this.routingNumber}-${this.accountNumber}`,
					type: "bank account",
					accountType: this.isSavingsAccount ? 'savings': 'checking',
					holderType: this.isBusinessAccount ? 'business' : 'personal'
				}
				: {
					source: "Visa Dickholes",
					number: `${this.cardNumber}-${this.expirationMonth}-${this.expirationYear}-${this.cvv}`,
					type: "card"
				}

			console.log(account)

			// TODO do something to create account
			await delay()
			this.account = account

			this.transitionAndClear(stages.PAID)
		}

		async makePayment() {
			const payment = {
				amount: this.amount,
				account: this.account,
			}

			console.log(payment)

			// TODO do something to pay
			await delay()
			if (Math.random() > 0.5) {
				// error
				this.paymentError = true
				this.account = null
				this.stage = stages.ACCOUNT
			}
			else
				this.transitionAndClear(stages.PAID)
		},

		transitionAndClear(stage) {
			this.paymentError = false
			this.stage = stage
		}
	}
}

</script>

<style lang="sass">



</style>
