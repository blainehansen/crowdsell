<template lang="pug">

#sign-up.content-tile

	h1 Join the Private Beta to be one of the first to launch a project, or to contribute

	h3 We're trying to make

	.form-group(@keyup.enter="submit")
		label(for="sign-up-email")
		input(type="email", placeholder="Enter email", aria-describedby="email-assurance", v-model="email")#sign-up-email.form-control
		small#email-assurance.form-text.text-muted
			| Your email will never be shared.

		button(type="submit", @click="submit").btn.btn-dark Submit

</template>

<script>

export default {
	data() {
		return {
			email: "",
		}
	},

	methods: {
		submit() {
			console.log(this.email)
		}
	},
}

</script>
