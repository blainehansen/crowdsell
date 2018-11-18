<template lang="pug">

b-navbar(toggleable="md")
	router-link.navbar-brand(:to="{ name: 'home' }") Home

	b-navbar-toggle(target="nav-collapse")

	b-collapse#nav-collapse(is-nav)
		ul.navbar-nav.mr-auto
			router-link.nav-item(:to="{ name: 'about' }", tag="li")
				a.nav-link About
		ul.navbar-nav
			template(v-if="userLoggedIn")
				router-link.nav-item(:to="{ name: 'userProfile' }", tag="li")
					a.nav-link {{ userName }}
				router-link.nav-item(:to="{ name: 'projectCreate' }", tag="li")
					a.nav-link Create Project
				li.nav-item(@click="logout")
					a.nav-link Logout
			template(v-else)
				router-link.nav-item(:to="{ name: 'projectsExplore' }", tag="li")
					a.nav-link Explore
				router-link.nav-item(:to="{ name: 'login' }", tag="li")
					a.nav-link Login

</template>


<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
	name: 'Header',

	computed: {
		...mapGetters('auth', [
			'userName', 'userLoggedIn'
		])
	},

	methods: {
		...mapMutations({
			logout: 'auth/logout'
		})
	},
}

</script>


<style lang="sass">
</style>
