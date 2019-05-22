<template lang="pug">

div
//- b-navbar(toggleable="md")#navbar.fixed-top.display-font
//- 	nuxt-link.navbar-brand.flex.justify-evenly.align-center(to="/")
//- 		.logo-box.item.push-right-less.self-start
//- 			//- Logo
//- 		span.item Crowdsell

//- 	b-navbar-toggle(target="nav-collapse")

//- 	b-collapse#nav-collapse(is-nav)
//- 		ul.navbar-nav.mr-auto
//- 			nuxt-link.nav-item(to="/why-we-need-it", tag="li")
//- 				a.nav-link The Problem
//- 			nuxt-link.nav-item(to="/the-mission", tag="li")
//- 				a.nav-link The Mission
//- 			nuxt-link.nav-item(to="/how-it-works", tag="li")
//- 				a.nav-link How It Works

//- 		ul.navbar-nav
//- 			nuxt-link.nav-item(to="/faq", tag="li")
//- 				a.nav-link FAQ

//- 			nuxt-link.nav-item(to="/sign-up", tag="li")
//- 				a.nav-link Join the Private Beta

			//- template(v-if="userLoggedIn")
			//- 	nuxt-link.nav-item(:to="{ name: 'userProfile' }", tag="li")
			//- 		a.nav-link {{ userName }}
			//- 	nuxt-link.nav-item(:to="{ name: 'projectCreate' }", tag="li")
			//- 		a.nav-link Create Project
			//- 	li.nav-item(@click="logout")
			//- 		a.nav-link Logout
			//- template(v-else)
			//- 	nuxt-link.nav-item(:to="{ name: 'projectsExplore' }", tag="li")
			//- 		a.nav-link Explore
			//- 	nuxt-link.nav-item(:to="{ name: 'login' }", tag="li")
			//- 		a.nav-link Login

</template>


<script>
// import { mapGetters, mapMutations } from 'vuex'

export default {
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
