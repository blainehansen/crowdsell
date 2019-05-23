<template lang="pug">

//- .content-between
nav.absolute.top-0.right-0.left-0.bg-transparent.w-full.z-50
	.container.flex.flex-row.flex-no-wrap.items-center
		.logo.my-30.flex.flex-row.flex-no-wrap.justify-start.items-center
			img(src="@/assets/images/logo_touched.svg").mr-10
			span.text-white.font-bold.leading-none.text-tricky-logo.mt-10 Crowdsell


		nuxt-link(to="/why-we-need-it").ml-auto.text-white.font-heavy.my-25.mx-30
			| The Problem
		nuxt-link(to="/the-mission").text-white.font-heavy.my-25.mx-30
			| The Mission
		nuxt-link(to="/how-it-works").text-white.font-heavy.my-25.mx-30
			| How It Works
		nuxt-link(to="/faq").text-white.font-heavy.my-25.mx-30
			| FAQ

		nuxt-link(to="/sign-up", tag="button").bg-greenish-teal.rounded-very.text-white.font-heavy.my-25.mx-30.py-15.px-30.leading-none.mr-0
			| Join us


//- b-navbar(toggleable="md")#navbar.fixed-top.display-font
//- 	nuxt-link.navbar-brand.flex.justify-evenly.align-center(to="/")
//- 		.logo-box.item.push-right-less.self-start
//- 			//- Logo
//- 		span.item Crowdsell

//- 	b-navbar-toggle(target="nav-collapse")

//- 	b-collapse#nav-collapse(is-nav)
//-

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
