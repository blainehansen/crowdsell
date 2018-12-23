import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/Home'
import SignUp from '@/pages/SignUp'

import HowItWorks from '@/pages/HowItWorks'
import WhyWeNeedIt from '@/pages/WhyWeNeedIt'
import WhyIsCrowdsellDifferent from '@/pages/WhyIsCrowdsellDifferent'
import TheMission from '@/pages/TheMission'
import WhoIsCrowdsellFor from '@/pages/WhoIsCrowdsellFor'
import FAQ from '@/pages/FAQ'

Vue.use(Router)


// I need pictures for:

// Home
// - the banner (Freeing Intellectual Property). we want to give people a sense of freedom and hope. something being released feels good, and something being widely available. I'm pretty set on representing intellectual property as pages
// removing it from some locked container
// unchaining it
// it falling from the sky
// someone rolling in a mountain of it
// the scroll birds flying around in a big beautiful swarm

// - (Intellectual Property is Broken). Want to evoke feelings of claustrophobia and control, and resentment towards greedy powerful people
// something about robots chaining people
// giant robot-like lawyers chasing people with lawsuits, people running away
// a robber baron type character holding it above people who are trying to get it
// people separated from it by a barred or other wall


// - (We Want to Fix That). Probably some version of the previous being undone

// - how does it work?
// all of these should be related, and depict different acts of the same scene

// 	- (done work). there's a paper, give it a "shiny mark". There's an important Letter or something
// 	- (preview). a piece of it is placed somewhere everyone can see. the important letter or seal is visible. something like it's put on a legal document or plaque or framed, and there's a checked checklist next to it
// 	- (goal). something about money or a tip jar, or a price. the previous plaque has a jar with money in it, one bill sticking out to indicate it might have just been put there
// 	- (released). focus on the release and duplication. the original paper is now a stack of them, with flying scrolls in the distance?

// - join us
// people carrying and about to throw the robber baron


// X Mission
// mountain top, possibly with the flying scrolls distantly fluttering around it

// X Why We Need (The Failure of Intellectual Property)
// a scroll with a chain on it

// X How it Works
// gears

// basic banner sponge background for all other top level?

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
	}, {
		path: '/faq',
		component: FAQ,
	}]

})

// TODO https://alligator.io/vuejs/vue-router-modify-head/
