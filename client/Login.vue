<template lang="pug">

#login
	h2(v-if="inLoginMode") Login
	h2(v-else) Create Account

	template(v-if="apiError")
		h3 There was an error!
		p {{ apiError.message }}

	input(v-if="!inLoginMode", v-model="name", placeholder="name")
	input(v-model="email", placeholder="email")
	input(v-model="password", type="password", placeholder="password")

	button(@click="submit") {{ inLoginMode ? "Login" : "Create Account" }}

	label(for="creation-mode-checkbox") {{ inLoginMode ? "Logging in" : "Creating new account" }}
	input#creation-mode-checkbox(v-model="inLoginMode", type="checkbox")

	label(for="remember-me-checkbox") Keep me Logged In
	input#remember-me-checkbox(v-model="rememberMe", type="checkbox")

</template>


<script>
import { publicApi } from '@/api'

export default {
	name: 'login',
	data() {
		return {
			inLoginMode: true,
			name: '',
			email: '',
			password: '',

			rememberMe: false,

			apiError: null,
		}
	},

	methods: {
		async submit() {
			const { inLoginMode, name, email, password } = this

			const apiFunction = inLoginMode ? publicApi.login : publicApi.createUser
			const args = inLoginMode ? [email, password] : [name, email, password]

			let signedUser = null
			try {
				const { data } = await apiFunction(...args)
				signedUser = data
			}
			catch (e) {
				this.apiError = e
				this.$store.commit('auth/logout')
			}

			if (signedUser !== null) {
				this.$store.commit('auth/login', signedUser)

				const goingTo = await this.$store.dispatch('auth/grabGoingTo')
				const routeObj = !!goingTo ? { path: goingTo } : { name: 'home' }
				this.$router.replace(routeObj)
			}
		},
	}
}
</script>


<style lang="sass">
</style>
