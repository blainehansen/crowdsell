import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import SignUp from '@/pages/SignUp'

import HowItWorks from '@/pages/HowItWorks'
import WhyWeNeedIt from '@/pages/WhyWeNeedIt'
import WhyIsCrowdsellDifferent from '@/pages/WhyIsCrowdsellDifferent'
import TheMission from '@/pages/TheMission'
import WhoIsCrowdsellFor from '@/pages/WhoIsCrowdsellFor'

Vue.use(Router)

export default new Router({
	mode: 'history',

	linkActiveClass: 'active',
	linkExactActiveClass: 'active-exact',

	routes: [{
		path: '/',
		component: Home,
	}, {
		path: '/sign-up',
		component: SignUp,
	}, {
		path: '/the-mission',
		component: TheMission,
	}, {
		path: '/why-we-need-it',
		component: WhyWeNeedIt,
	}, {
		path: '/how-it-works',
		component: HowItWorks,
	}, {
		path: '/why-is-crowdsell-different',
		component: WhyIsCrowdsellDifferent,
	}, {
		path: '/who-is-crowdsell-for',
		component: WhoIsCrowdsellFor,
	}]

})

// TODO https://alligator.io/vuejs/vue-router-modify-head/
