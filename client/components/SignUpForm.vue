<template lang="pug">

.signup-form(@keyup.enter="submitEmail")
	//- label(for="sign-up-email")
	.fake-input.flex.flex-row.flex-nowrap.items-center.justify-start.bg-pale-grey.border.border-solid.border-reddish-grey.rounded-more.p-10.pl-30
		input.text-small.text-blue-grey.bg-pale-grey.outline-none.leading-none(
			v-model="email",
			placeholder="you@internet.com",
			type="email",
			aria-describedby="email-assurance",
			:class="{ 'is-invalid': emailInvalid }",
		)
		button(@click="submitEmail").rounded-normal.bg-bright-blue.text-white.font-bold.join-button.leading-none.py-20.px-42.ml-auto
			| Join the Mission

	#email-assurance.mt-20.text-pastel-red.text-tiny(v-if="emailInvalid") Email Invalid!
	//- .text-white not necessary, just letting the default of the container take over?
	#email-assurance.mt-20.text-tiny(v-else)
		| Your email will never be shared.

</template>


<script>

export default {
	// TODO accept props that can change the colors of things
	// TODO accept a prop for the button text

	data() {
		return {
			email: "",
			attemptedSubmit: false,
		}
	},

	computed: {
		emailInvalid() {
			return this.attemptedSubmit && !(/^.+@.+\..+$/.test(this.email))
		},
	},

	methods: {
		submitEmail() {
			this.attemptedSubmit = true

			// if (this.emailInvalid) console.log('nope')
			// else {
			// 	this.attemptedSubmit = false
			// 	console.log(this.email)
			// }
		}
	},
}

</script>
