<template lang="pug">

.signup-form
	.form-group(@keyup.enter="submitEmail")
		label(for="sign-up-email")
		input#sign-up-email.form-control(
			type="email",
			placeholder="Enter email",
			aria-describedby="email-assurance",
			v-model="email",
			:class="{ 'is-invalid': emailInvalid }",
		)
		.invalid-feedback(v-if="emailInvalid") Email Invalid!
		small(v-else)#email-assurance.form-text.text-muted
			| Your email will never be shared.

	button(type="submit", @click="submitEmail").btn.btn-dark Submit

</template>


<script>

export default {
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
