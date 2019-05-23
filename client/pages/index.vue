<template lang="pug">

#home.bg-white

	//- .banner-cloud

	#banner.relative.h-banner
		.container.h-full.flex.flex-row.flex-no-wrap.justify-between.items-center
			#banner-text.text-white.w-55p.pr-4
				h1.text-hero.font-bold.mb-20.leading-relaxed
					| Freeing Intellectual Property

				p.text-heavy.mb-50
					| Using crowdfunding to change the way we think about intellectual work, and how people are compensated for it.

				SignUpForm.w-85p

			#banner-picture.w-45p.h-60p.flex.justify-center.items-center.relative
				img(src="@/assets/images/how-could-crowdsell-change-society_touched.svg").w-full.absolute.left-0.right-0.m-auto.z-20
				img(src="@/assets/images/scroll-big-split.svg").absolute.top-25p.left-5p.h-22p.z-20
				img(src="@/assets/images/scroll-big-together.svg").mirror.absolute.top-22p.right-5p.h-22p.z-20
				img(src="@/assets/images/scroll-small.svg").rotate-little.absolute.top-12p.left-30p.z-20
				img(src="@/assets/images/scroll-small.svg").rotate-more.absolute.top-4p.right-35p.h-10p.z-20

				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.top-0.right-0.h-8p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.top-8p.left-0.h-6p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.top-22p.left-25p.h-6p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.top-40p.left-6p.h-8p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.bottom-18p.left-8p.h-5p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.bottom-25p.right-8p.h-6p
				img(src="@/assets/images/cloud.svg").cloud.z-10.absolute.bottom-20p.right-35p.h-4p

		img#banner-wave(src="@/assets/images/download.svg").absolute.bottom--1

	#broken
		.container
			div
				#broken-text
					h2.text-header.text-dark-blue-grey
						| Intellectual Property is Broken

					p.text-lead.text-blue-grey
						| It's stifled our culture and strangled innovation.
						//- In the world today, the one that's driven by Intellectual Property, the only way to make money doing creative work is with coercion, control, and walled gardens.
						//- That status quo has produced countless harmful businesses that focus more on controlling people than contributing to the public good.
					p.text-lead.text-blue-grey
						| Powerful platforms built on Intellectual Property have the power to create monopolies, trap us all in echo chambers, and extract wealth from society.

				#broken-picture
					//- IntellectualPropertyIsBroken

	#something-new
		.container
			div
				#something-new-text
					h2.text-header.text-dark-blue-grey
						| We Need Something New

					p.text-lead.text-blue-grey
						| Crowdsell is a crowdfunding system that allows creators to fund and sell their intellectual work to the whole world.
					p.text-lead.text-blue-grey
						| By adding binding promises and oversight to protect backers, communities can safely and predictably reward creators for contributing to the commons and the public good.

				#something-new-picture
					//- WeNeedSomethingNew


	#how-does-it-work.bg-white-two
		.container.tight
			h2.text-center.text-dark-blue-grey So how does Crowdsell work?

			.explain-box
				.explain-box-text
					h3 The Work
					p Crowdsell is for valuable intellectual work, like code, art or designs.
					p Creators can either sell finished work, or fundraise a budget for a big idea.

				.explain-box-picture
					//- CompletedParchment

			.explain-box
				.explain-box-text
					h3 Previews and Promises
					p
						| First they create a proof of concept or a preview to show they are capable, and choose any kind of&nbsp;
						out-link(href="https://creativecommons.org/share-your-work/public-domain/freeworks/") open culture license
						| &nbsp;that is&nbsp;
						out-link(href="https://opensource.org/licenses") appropriate for their project
						| .
					p Then they make a detailed list of promises about the work that the backers will expect it to live up to.

				.explain-box-picture
					//- APlaquePreviewingTheParchment

			.explain-box
				.explain-box-text
					h3 Budget and Reward
					p They choose a reward amount, and make a monthly budget if the work isn't complete yet.
					p Anyone in the world can contribute to fundraise the total.

				.explain-box-picture
					//- AccountingPaperWithATipJar

			.explain-box
				.explain-box-text
					h3 The Backers Decide
					p If needed, the budget is released month by month as the backers give oversight.
					p Once the backers decide that the work fulfills the original promises, then the creator is given the reward amount.

				.explain-box-picture
					//- TheAccountingPaperInspectedWithABallotBox

			.explain-box
				.explain-box-text
					h3 The Work is Set Free
					p The creator has been rewarded for their creativity and skill, and the world has access to valuable work free of legal restrictions.
					p It belongs to everyone now.

				.explain-box-picture
					//- TheParchmentFliesAway


	#join-us.bg-bright-blue
		.container.tight.text-center.text-white
			h2.text-header Join Us
			p.text-heavy
				| We're on a mission to make&nbsp;
				out-link(href="https://creativecommonsusa.org/index.php/open-culture/") open knowledge, open culture, and intellectual collaboration
				| &nbsp;the norm.
			p.text-heavy
				| Join the mailing list to hear from us and become a part of this new movement.

			//- SignUpForm(:center, submitText="Submit", assuranceColor="bluey-grey")
			SignUpForm


	#faq.bg-pale-grey-two
		.container.medium
			h2.text-center.text-header.text-dark-blue-grey
				| FAQ and Resources
			p.text-center.text-lead.text-blue-grey
				| Something not clear? Want to learn more? Check out these links.

			nuxt-link(to="/faq", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| Frequently Asked Questions
				//- the arrow svg
				//- .text-bright-blue .fill-current

			nuxt-link(to="/why-we-need-it", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| Is there really a problem with Intellectual Property?
				//- the arrow svg

			nuxt-link(to="/the-mission", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| How could Crowdsell change society?
				//- the arrow svg

			nuxt-link(to="/how-it-works", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| How does Crowdsell work in detail?
				//- the arrow svg

			nuxt-link(to="/who-is-crowdsell-for", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| Who is Crowdsell for? And what can they do with it?
				//- the arrow svg

			nuxt-link(to="/why-is-crowdsell-different", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| How is Crowdsell different than other crowdfunding systems?
				//- the arrow svg

			nuxt-link(to="/how-does-crowdsell-make-money", tag="div").bg-white.rounded-more
				span.text-base.font-bold.text-dark-blue-grey
					| How does Crowdsell make money?
				//- the arrow svg

</template>

<!-- <style lang="sass"> -->
<style lang="sss">
#home
	#banner
		background-image:
			url('../assets/images/ornament.svg'),
			linear-gradient(233deg, theme('colors.violet-blue'), theme('colors.bright-blue'))

		background-repeat:
			no-repeat,
			no-repeat

		background-position:
			top left,
			bottom left

		background-size:
			80% 85%,
			cover

		#banner-wave
			margin-left: -10px
			margin-right: -10px
			width: 100vw
			max-width: initial


	.explain-box
		// @apply
		.explain-box-text
			h3
				@apply text-lead text-dark-blue-grey
			p
				@apply text-base text-blue-grey

	// #join-us
		// the decorative background

</style>
