<template lang="pug">

#sign-up.full-height.container.flex.align-center.justify-evenly

	.text-box.limit-40
		h2.display-4 Mailing List

		p.display-paragraph Join the Private Beta to be one of the first to launch a project, or to contribute

		p.display-paragraph We're trying to make

	#email-form
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
