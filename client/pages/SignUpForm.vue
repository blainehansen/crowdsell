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

import launchEmailMixin from '@/launchEmailMixin'

export default {
	mixins: [launchEmailMixin],
}

</script>
