<template lang="pug">

b-navbar(toggleable="md")#navbar.fixed-top.display-font
	router-link.navbar-brand.flex.justify-evenly.align-center(to="/")
		.logo-box.item.push-right-less.self-start
			Logo
		span.item Crowdsell

	b-navbar-toggle(target="nav-collapse")

	b-collapse#nav-collapse(is-nav)
		ul.navbar-nav.mr-auto
			router-link.nav-item(to="/why-we-need-it", tag="li")
				a.nav-link The Problem
			router-link.nav-item(to="/the-mission", tag="li")
				a.nav-link The Mission
			router-link.nav-item(to="/how-it-works", tag="li")
				a.nav-link How It Works

		ul.navbar-nav
			router-link.nav-item(to="/faq", tag="li")
				a.nav-link FAQ

			router-link.nav-item(to="/sign-up", tag="li")
				a.nav-link Join the Private Beta

			//- template(v-if="userLoggedIn")
			//- 	router-link.nav-item(:to="{ name: 'userProfile' }", tag="li")
			//- 		a.nav-link {{ userName }}
			//- 	router-link.nav-item(:to="{ name: 'projectCreate' }", tag="li")
			//- 		a.nav-link Create Project
			//- 	li.nav-item(@click="logout")
			//- 		a.nav-link Logout
			//- template(v-else)
			//- 	router-link.nav-item(:to="{ name: 'projectsExplore' }", tag="li")
			//- 		a.nav-link Explore
			//- 	router-link.nav-item(:to="{ name: 'login' }", tag="li")
			//- 		a.nav-link Login

</template>


<script>
// import { mapGetters, mapMutations } from 'vuex'

import Logo from "@/assets/illustrations/logo_touched.svg"

export default {
	name: 'Header',

	components: {
		Logo,
	},

// 	computed: {
// 		...mapGetters('auth', [
// 			'userName', 'userLoggedIn'
// 		])
// 	},

// 	methods: {
// 		...mapMutations({
// 			logout: 'auth/logout'
// 		})
// 	},
}

</script>


<style lang="sass" scoped>

#navbar
	background-color: white
	box-shadow: 0 2px 5px 0 rgba(0,0,0,0.2)
	height: 4.5rem

.navbar-brand
	font-size: 1.2rem

.logo-box svg
	width: 2rem
	height: auto

</style>
