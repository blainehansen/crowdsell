<template lang="pug">

#home
	#banner.container.full-height.flex.align-center.justify-center
		.text-box
			h1#banner-header.display-4
				| Freeing
				br
				| Intellectual
				br
				| Property

		.picture-box.big
			FreeingIntellectualProperty

	.normal-tile.container.flex.reverse.justify-evenly.pad-top-huge
		.text-box.limit-40.text-right
			h3.display-4
				| Intellectual
				br
				| Property
				br
				| is Broken

			p.display-paragraph It's stifled our culture and strangled innovation.
			p.display-paragraph Big corporations have the power to create monopolies, trap us all in echo chambers, and extract wealth from society.

		.picture-box.normal
			IntellectualPropertyIsBroken

	.normal-tile.container.flex.align-center.justify-evenly
		.text-box.limit-40
			h3.display-4
				| We Need
				br
				| Something
				br
				| New

			p.display-paragraph Crowdsell is a new crowdfunding system that allows creators to fund and sell their intellectual work to the whole world.
			p.display-paragraph
				| By adding binding promises and oversight to protect backers, communities can reward creators for contributing to the&nbsp;
				out-link(href="https://creativecommonsusa.org/index.php/open-culture/") commons and the public good
				| .

		.picture-box.normal
			WeNeedSomethingNew

	.container.flex.column.align-center.pad-top-large
		h2.center-text.push-bottom-large.display-3 So how does it work?

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p Crowdsell is for valuable intellectual work, like code, art or designs.
				p Creators can either sell finished work, or fundraise a budget for a big idea.
			.picture-box.item.feature
				CompletedParchment

		.feature-box.limit-40.flex.reverse.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p
					| First they create a proof of concept or a preview to show they are capable, and choose an&nbsp;
					out-link(href="https://creativecommons.org/share-your-work/public-domain/freeworks/") open culture license
					| .
				p Then they make a detailed list of promises about the work that the backers will expect it to live up to.
			.picture-box.item.feature
				APlaquePreviewingTheParchment

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p They choose a reward amount, and make a monthly budget if the work isn't complete yet.
				p Anyone in the world can contribute to fundraise the total.
			.picture-box.item.feature
				AccountingPaperWithATipJar

		.feature-box.limit-40.flex.reverse.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p If needed, the budget is released month by month as the backers give oversight.
				p Once the backers decide that the work fulfills the original promises, then the creator is given the reward amount.
			.picture-box.item.feature
				TheAccountingPaperInspectedWithABallotBox

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p The creator has been rewarded for their creativity and skill, and the world has access to valuable work free of legal restrictions.
				p It belongs to everyone now.
			.picture-box.item.feature
				TheParchmentFliesAway


	.normal-tile.container.flex.align-center.justify-evenly
		.text-box.limit-40
			h2.display-3 Join Us
			p.display-paragraph We're on a mission to make open knowledge, open culture, and intellectual collaboration the norm.
			router-link(to="/sign-up", tag="a").btn.btn-dark Join the Private Beta

		.picture-box.big
			HowCouldCrowdsellChangeSociety


	.container.pad-top-large.pad-bottom-huge
		h3 Learn more

		.page-box
			router-link(to="/faq")
				h4 Frequently Asked Questions

		.page-box
			router-link(to="/why-we-need-it")
				h4 Is there really a problem with Intellectual Property?

		.page-box
			router-link(to="/the-mission")
				h4 How could Crowdsell change society?

		.page-box
			router-link(to="/who-is-crowdsell-for")
				h4 Who is Crowdsell for? And what can they do with it?

		.page-box
			router-link(to="/how-it-works")
				h4 How does Crowdsell work in detail?

		.page-box
			router-link(to="/why-is-crowdsell-different")
				h4 How is Crowdsell different than other crowdfunding systems?


	footer

</template>

<script>

import FreeingIntellectualProperty from "@/assets/illustrations/large-scenes/freeing-intellectual-property.svg"
import IntellectualPropertyIsBroken from "@/assets/illustrations/large-scenes/intellectual-property-is-broken.svg"
import WeNeedSomethingNew from "@/assets/illustrations/large-scenes/we-need-something-new.svg"
import CompletedParchment from "@/assets/illustrations/feature-panels/completed-parchment.svg"
import APlaquePreviewingTheParchment from "@/assets/illustrations/feature-panels/a-plaque-previewing-the-parchment.svg"
import AccountingPaperWithATipJar from "@/assets/illustrations/feature-panels/accounting-paper-with-a-tip-jar.svg"
import TheAccountingPaperInspectedWithABallotBox from "@/assets/illustrations/feature-panels/the-accounting-paper-inspected-with-a-ballot-box.svg"
import TheParchmentFliesAway from "@/assets/illustrations/feature-panels/the-parchment-flies-away.svg"
import HowCouldCrowdsellChangeSociety from "@/assets/illustrations/small-scenes/how-could-crowdsell-change-society.svg"

export default {
	name: 'home',
	components: {
		FreeingIntellectualProperty,
		IntellectualPropertyIsBroken,
		WeNeedSomethingNew,
		CompletedParchment,
		APlaquePreviewingTheParchment,
		AccountingPaperWithATipJar,
		TheAccountingPaperInspectedWithABallotBox,
		TheParchmentFliesAway,
		HowCouldCrowdsellChangeSociety,
	}
}

</script>

<style lang="sass" scoped>


// .text-box
// 	font-size: 2rem

.picture-box
	&.big
		width: 70%

	&.normal
		width: 60%

	&.feature
		width: 40%

	img
		width: 100%
		height: auto

#home
	background: linear-gradient(to bottom, #a09ad6 0%, rgba(255,255,255,0) 800px)

#banner-header
	line-height: 1.6

</style>
