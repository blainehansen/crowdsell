<template lang="pug">

//- .content-between
nav.absolute.top-0.right-0.left-0.bg-transparent.w-full.z-50
	.container.flex.flex-row.flex-no-wrap.items-center
		nuxt-link(to="/").logo.my-30.flex.flex-row.flex-no-wrap.justify-start.items-center.cursor-pointer.text-white.hover_text-light-teal.hover_no-underline
			Logo.mr-10.fill-current
			span.font-bold.leading-none.text-tricky-logo.mt-10 Crowdsell


		nuxt-link(to="/why-we-need-it").ml-auto.font-heavy.my-25.mx-30.text-white.hover_no-underline.hover_text-light-teal
			| The Problem
		nuxt-link(to="/the-mission").font-heavy.my-25.mx-30.text-white.hover_no-underline.hover_text-light-teal
			| The Mission
		nuxt-link(to="/how-it-works").font-heavy.my-25.mx-30.text-white.hover_no-underline.hover_text-light-teal
			| How It Works
		nuxt-link(to="/faq").font-heavy.my-25.mx-30.text-white.hover_no-underline.hover_text-light-teal
			| FAQ

		nuxt-link(to="/sign-up")
			button.bg-greenish-teal.rounded-very.text-white.font-heavy.my-25.mx-30.py-15.px-30.leading-none.mr-0.hover_bg-light-teal
				| Join us


//- b-navbar(toggleable="md")#navbar.fixed-top.display-font

//- 	b-navbar-toggle(target="nav-collapse")

//- 	b-collapse#nav-collapse(is-nav)

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
