<template lang="pug">

.signup-form(@keyup.enter="submitEmail")
	template(v-if="!separateMode")
		//- label(for="sign-up-email")
		.fake-input.flex.flex-row.flex-nowrap.items-center.justify-start.bg-pale-grey.border.border-solid.border-reddish-grey.rounded-more.p-10.pl-30
			input.text-small.text-blue-grey.bg-pale-grey.outline-none.leading-none(
				v-model="email",
				placeholder="you@internet.com",
				type="email",
				aria-describedby="email-assurance",
				:class="{ 'is-invalid': emailInvalid }",
			)
			button(@click="submitEmail").bg-greenish-teal.rounded-normal.text-white.font-bold.join-button.leading-none.py-20.px-42.ml-auto
				| Join us

		#email-assurance.mt-20.text-pastel-red.text-tiny(v-if="emailInvalid") Email Invalid!
		#email-assurance.mt-20.text-tiny(v-else, :class="[assuranceClass]")
			| Your email will never be shared.

	template(v-else)
		#email-assurance.mb-40.text-tiny.text-pastel-red(v-if="emailInvalid") Email Invalid!
		#email-assurance.mb-40.text-tiny.text-steel(v-else, :class="[assuranceClass]")
			| Your email will never be shared.

		input.text-steel.text-small.rounded-more.p-20.border.border-solid.border-silver.w-full.mb-40.outline-none(
			v-model="email",
			placeholder="you@internet.com",
			type="email",
			aria-describedby="email-assurance",
			:class="{ 'is-invalid': emailInvalid }",
		)
		button(@click="submitEmail").bg-greenish-teal.rounded-normal.text-white.font-bold.join-button.leading-none.py-20.w-full
			| Join us

</template>


<script>

export default {
	props: {
		assuranceClass: {
			type: String,
			required: false,
		},

		separateMode: {
			type: Boolean,
			required: false,
			default: false,
		},
	},

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
