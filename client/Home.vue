<template lang="pug">

#home
	#banner.full-height.container-fluid.color-background
		CloudSpecies1.cloud(style="left: 30%; top: 20%; width: 20%;")
		CloudSpecies2.cloud(style="left: 15%; bottom: 10%; width: 30%;")
		//- CloudSpecies3.cloud(style="right: 20%; bottom: 15%; width: 22%;")
		//- CloudSpecies4.cloud(style="right: 20%; bottom: 10%; width: 22%;")
		CloudSpecies5.cloud(style="right: 6%; bottom: 6%; width: 22%;")
		//- CloudSpecies6.cloud(style="right: 20%; bottom: 10%; width: 22%;")
		//- CloudSpecies7.cloud(style="right: 20%; bottom: 10%; width: 22%;")

		.row.align-items-center.justify-content-evenly.fill-height
			.text-box.col-md-4
				h1#banner-header.display-3
					| Freeing
					br
					| Intellectual
					br
					| Property

				p.display-paragraph Using crowdfunding to change the way we think about intellectual work, and how people are compensated for it.

				p
					h3.push-bottom-less Join the Mailing List

					SignUpForm

			#hero-picture.picture-box.col-md-5
				FreeingIntellectualProperty

	#broken.normal-tile
		CloudSpecies1.cloud(style="left: 6%; top: 20%; width: 30%;")
		CloudSpecies2.cloud(style="left: 4%; bottom: 10%; width: 28%;")
		//- CloudSpecies3.cloud(style="right: 20%; bottom: 15%; width: 22%;")
		//- CloudSpecies4.cloud(style="right: 20%; bottom: 10%; width: 22%;")
		CloudSpecies6.cloud(style="right: 14%; top: 10%; width: 35%;")
		CloudSpecies5.cloud(style="right: 6%; bottom: 6%; width: 38%;")
		//- CloudSpecies7.cloud(style="right: 20%; bottom: 10%; width: 22%;")

		.container.fill-height
			.row.flex-row-reverse.align-items-center.justify-content-around.fill-height
				.text-box.col-md-5
					h3.display-4
						| Intellectual
						br
						| Property
						br
						| is Broken

					p.display-paragraph It's stifled our culture and strangled innovation.
					p.display-paragraph Powerful platforms built on Intellectual Property have the power to create monopolies, trap us all in echo chambers, and extract wealth from society.

				.picture-box.col-md-5
					IntellectualPropertyIsBroken

	#new.normal-tile.color-background
		.container.fill-height
			.row.align-items-center.justify-content-around.fill-height
				.text-box.col-md-5
					h3.display-4
						| We Need
						br
						| Something
						br
						| New

					p.display-paragraph Crowdsell is a crowdfunding system that allows creators to fund and sell their intellectual work to the whole world.
					p.display-paragraph By adding binding promises and oversight to protect backers, communities can safely and predictably reward creators for contributing to the commons and the public good.

				.picture-box.col-md-5
					WeNeedSomethingNew

	.container.pad-top-large
		.row.justify-content-center
			h2.text-center.push-bottom-large.display-4 So how does it work?

		.row.justify-content-evenly.align-items-center.push-bottom-large
			.text-box.feature-box.col-md-4
				h4 The Work
				p Crowdsell is for valuable intellectual work, like code, art or designs.
				p Creators can either sell finished work, or fundraise a budget for a big idea.
			.picture-box.col-md-3
				CompletedParchment

		.row.justify-content-evenly.align-items-center.push-bottom-large.flex-row-reverse
			.text-box.feature-box.col-md-4
				h4 Previews and Promises
				p
					| First they create a proof of concept or a preview to show they are capable, and choose any kind of&nbsp;
					out-link(href="https://creativecommons.org/share-your-work/public-domain/freeworks/") open culture license
					| &nbsp;that is&nbsp;
					out-link(href="https://opensource.org/licenses") appropriate for their project
					| .
				p Then they make a detailed list of promises about the work that the backers will expect it to live up to.
			.picture-box.col-md-3
				APlaquePreviewingTheParchment

		.row.justify-content-evenly.align-items-center.push-bottom-large
			.text-box.feature-box.col-md-4
				h4 Budget and Reward
				p They choose a reward amount, and make a monthly budget if the work isn't complete yet.
				p Anyone in the world can contribute to fundraise the total.
			.picture-box.col-md-3
				AccountingPaperWithATipJar

		.row.justify-content-evenly.align-items-center.push-bottom-large.flex-row-reverse
			.text-box.feature-box.col-md-4
				h4 The Backers Decide
				p If needed, the budget is released month by month as the backers give oversight.
				p Once the backers decide that the work fulfills the original promises, then the creator is given the reward amount.
			.picture-box.col-md-3
				TheAccountingPaperInspectedWithABallotBox

		.row.justify-content-evenly.align-items-center.push-bottom-large
			.text-box.feature-box.col-md-4
				h4 The Work is Set Free
				p The creator has been rewarded for their creativity and skill, and the world has access to valuable work free of legal restrictions.
				p It belongs to everyone now.
			.picture-box.col-md-3
				TheParchmentFliesAway


	.normal-tile.color-background
		.container.fill-height
			.row.fill-height.justify-content-evenly.align-items-center
				.text-box.col-md-5
					h2.display-3 Join Us
					p.display-paragraph
						| We're on a mission to make&nbsp;
						out-link(href="https://creativecommonsusa.org/index.php/open-culture/") open knowledge, open culture, and intellectual collaboration
						| &nbsp;the norm.
					p.display-paragraph.push-bottom-less
						| Join the mailing list to hear from us and become a part of this new movement.
					//- router-link(to="/sign-up", tag="a").btn.btn-dark Join the Private Beta
					SignUpForm

				.picture-box.col-md-5
					HowCouldCrowdsellChangeSociety


	.container.pad-top-large.pad-bottom-huge
		.row.justify-content-center
			.col-md-10
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
					router-link(to="/how-it-works")
						h4 How does Crowdsell work in detail?

				.page-box
					router-link(to="/who-is-crowdsell-for")
						h4 Who is Crowdsell for? And what can they do with it?

				.page-box
					router-link(to="/why-is-crowdsell-different")
						h4 How is Crowdsell different than other crowdfunding systems?

				.page-box
					router-link(to="/how-does-crowdsell-make-money")
						h4 How does Crowdsell make money?

</template>

<script>

import FreeingIntellectualProperty from "@/assets/illustrations/large-scenes/freeing-intellectual-property_touched.svg"
import IntellectualPropertyIsBroken from "@/assets/illustrations/large-scenes/intellectual-property-is-broken_touched.svg"
import WeNeedSomethingNew from "@/assets/illustrations/large-scenes/we-need-something-new_touched.svg"

import CloudSpecies1 from "@/assets/illustrations/cloud_species_1.svg"
import CloudSpecies2 from "@/assets/illustrations/cloud_species_2.svg"
import CloudSpecies3 from "@/assets/illustrations/cloud_species_3.svg"
import CloudSpecies4 from "@/assets/illustrations/cloud_species_4.svg"
import CloudSpecies5 from "@/assets/illustrations/cloud_species_5.svg"
import CloudSpecies6 from "@/assets/illustrations/cloud_species_6.svg"
import CloudSpecies7 from "@/assets/illustrations/cloud_species_7.svg"

import CompletedParchment from "@/assets/illustrations/feature-panels/completed-parchment_touched.svg"
import APlaquePreviewingTheParchment from "@/assets/illustrations/feature-panels/a-plaque-previewing-the-parchment_touched.svg"
import AccountingPaperWithATipJar from "@/assets/illustrations/feature-panels/accounting-paper-with-a-tip-jar_touched.svg"
import TheAccountingPaperInspectedWithABallotBox from "@/assets/illustrations/feature-panels/the-accounting-paper-inspected-with-a-ballot-box_touched.svg"
import TheParchmentFliesAway from "@/assets/illustrations/feature-panels/the-parchment-flies-away_touched.svg"

import HowCouldCrowdsellChangeSociety from "@/assets/illustrations/small-scenes/how-could-crowdsell-change-society_touched.svg"

export default {
	name: 'home',
	components: {
		FreeingIntellectualProperty,
		IntellectualPropertyIsBroken,
		WeNeedSomethingNew,

		CloudSpecies1,
		CloudSpecies2,
		CloudSpecies3,
		CloudSpecies4,
		CloudSpecies5,
		CloudSpecies6,
		CloudSpecies7,

		CompletedParchment,
		APlaquePreviewingTheParchment,
		AccountingPaperWithATipJar,
		TheAccountingPaperInspectedWithABallotBox,
		TheParchmentFliesAway,
		HowCouldCrowdsellChangeSociety,
	}
}

</script>

<style lang="sass">

@import "~styles/variables"

.picture-box
	svg
		width: 100%
		height: auto


.cloud
	position: absolute
	opacity: 0.4
	height: auto

#banner svg.cloud
	path
		fill: $true-white

#broken
	position: relative
	svg.cloud path
		fill: $black
		opacity: 0.07

.feature-box
	font-size: 1.1rem

.horizontal-rule
	margin: auto
	height: 0.08rem
	width: 60vw
	background-color: $gray



.color-background
	// $color-primary-1
	$slightly-dark: darken($primary-almost-white, 1%)
	background: linear-gradient(180deg, $slightly-dark 5%, 40%, $primary-almost-white 90%)
	// $slightly-dark: darken($color-primary-1, 5%)
	// background: linear-gradient(0deg, $slightly-dark 5%, 40%, $color-primary-1 90%)


// #banner-background
// 	position: absolute
// 	top: 0
// 	left: 0
// 	right: 0
// 	z-index: -1000
// 	height: 140vh
// 	background: linear-gradient(175deg, $slightly-dark 5%, 40%, $true-white 90%)

#banner
	#banner-header
		line-height: 1.2


#broken
	// $start: transparentize(darken($black, 5%), 0.9)
	// $end: transparentize($black, 0.9)
	// background: linear-gradient(274deg, $start, $end)

#new
	// $slightly-dark: darken($primary-almost-white, 5%)
	// background: linear-gradient(274deg, $slightly-dark 5%, 70%, $primary-almost-white)

.normal-tile
	height: 70vh

.display-paragraph
	font-size: 1.2rem

</style>
