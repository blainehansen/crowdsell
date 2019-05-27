<template lang="pug">

#home
	#banner.relative.md_h-banner.header-adjacent
		.container.h-full.md_flex.md_flex-row.md_flex-no-wrap.md_justify-between.md_items-center
			#banner-text.text-white.md_w-55p.md_pr-4
				h1.text-hero.font-bold.mb-20.leading-relaxed
					| Freeing Intellectual Property

				p.text-heavy.mb-50
					| Using crowdfunding to change the way we think about intellectual work, and how people are compensated for it.

				SignUpForm#current-signup-form.md_w-85p

			#banner-picture.h-tile.md_w-45p.md_h-60p.relative
				img(src="@/assets/images/how-could-crowdsell-change-society_touched.svg").w-full.absolute.bottom-20p.left-0.right-0.m-auto.z-20
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

	#broken.mt-100.md_mt-175
		.container.flex.flex-col-reverse.md_flex-row.flex-no-wrap.md_justify-between.md_items-center.h-full
			#broken-text.md_w-50p
				h2.text-header.text-dark-blue-grey.font-bold.leading-tight-kinda.md_w-60p.mb-40.mt-30.md_mt-0
					| Intellectual Property is Broken

				p.text-lead.text-blue-grey.leading-tight.mb-30
					| It's stifled our culture and strangled innovation.
					//- In the world today, the one that's driven by Intellectual Property, the only way to make money doing creative work is with coercion, control, and walled gardens.
					//- That status quo has produced countless harmful businesses that focus more on controlling people than contributing to the public good.
				p.text-lead.text-blue-grey.leading-tight
					| Powerful platforms built on Intellectual Property have the power to create monopolies, trap us all in echo chambers, and extract wealth from society.

			#broken-picture.m-auto.md_ml-auto.md_mr-0.w-80p.md_w-40p
				img(src="@/assets/images/intellectual-property-is-broken_touched.svg")

	#something-new.mt-100.md_mt-175.mb-150.md_mb-300
		.container.flex.flex-col-reverse.md_flex-row-reverse.flex-no-wrap.md_justify-between.md_items-center.h-full
			#something-new-text.md_w-50p
				h2.text-header.text-dark-blue-grey.font-bold.leading-tight-kinda.md_w-60p.mb-40.mt-30.md_mt-0
					| We Need Something New

				p.text-lead.text-blue-grey.leading-tight.mb-30
					| Crowdsell is a crowdfunding system that allows creators to fund and sell their intellectual work to the whole world.
				p.text-lead.text-blue-grey.leading-tight
					| By adding binding promises and oversight to protect backers, communities can safely and predictably reward creators for contributing to the commons and the public good.

			#something-new-picture.m-auto.md_mr-auto.md_ml-0.w-80p.md_w-40p
				img(src="@/assets/images/we-need-something-new_touched.svg")


	#how-does-it-work.bg-white-two
		.container.tight
			h2.text-center.text-header.text-dark-blue-grey.font-heavy.mb-50.md_mb-125 So how does Crowdsell work?

			.explain-box.flex.flex-row.flex-no-wrap.justify-between.items-center.mb-80.md_mb-100
				.explain-box-text.w-60p
					h3.text-lead.text-dark-blue-grey.mb-10
						| The Work
					p.text-tiny.text-blue-grey.leading-airy.mb-15
						| Crowdsell is for valuable intellectual work, like code, art or designs.
					p.text-tiny.text-blue-grey.leading-airy
						| Creators can either sell finished work, or fundraise a budget for a big idea.

				img.explain-box-picture.w-30p(src="@/assets/images/completed-parchment_touched.svg")

			.explain-box.flex.flex-row-reverse.flex-no-wrap.justify-between.items-center.mb-80.md_mb-100
				.explain-box-text.w-60p
					h3.text-lead.text-dark-blue-grey.mb-10
						| Previews and Promises
					p.text-tiny.text-blue-grey.leading-airy.mb-15
						| First they create a proof of concept or a preview to show they are capable, and choose any kind of&nbsp;
						out-link(href="https://creativecommons.org/share-your-work/public-domain/freeworks/") open culture license
						| &nbsp;that is&nbsp;
						out-link(href="https://opensource.org/licenses") appropriate for their project
						| .
					p.text-tiny.text-blue-grey.leading-airy
						| Then they make a detailed list of promises about the work that the backers will expect it to live up to.

				img.explain-box-picture.w-30p(src="@/assets/images/a-plaque-previewing-the-parchment_touched.svg")

			.explain-box.flex.flex-row.flex-no-wrap.justify-between.items-center.mb-80.md_mb-100
				.explain-box-text.w-60p
					h3.text-lead.text-dark-blue-grey.mb-10
						| Budget and Reward
					p.text-tiny.text-blue-grey.leading-airy.mb-15
						| They choose a reward amount, and make a monthly budget if the work isn't complete yet.
					p.text-tiny.text-blue-grey.leading-airy
						| Anyone in the world can contribute to fundraise the total.

				img.explain-box-picture.w-30p(src="@/assets/images/accounting-paper-with-a-tip-jar_touched.svg")

			.explain-box.flex.flex-row-reverse.flex-no-wrap.justify-between.items-center.mb-80.md_mb-100
				.explain-box-text.w-60p
					h3.text-lead.text-dark-blue-grey.mb-10
						| The Backers Decide
					p.text-tiny.text-blue-grey.leading-airy.mb-15
						| If needed, the budget is released month by month as the backers give oversight.
					p.text-tiny.text-blue-grey.leading-airy
						| Once the backers decide that the work fulfills the original promises, then the creator is given the reward amount.

				img.explain-box-picture.w-30p(src="@/assets/images/the-accounting-paper-inspected-with-a-ballot-box_touched.svg")

			.explain-box.flex.flex-row.flex-no-wrap.justify-between.items-center.mb-100.md_mb-200
				.explain-box-text.w-60p
					h3.text-lead.text-dark-blue-grey.mb-10
						| The Work is Set Free
					p.text-tiny.text-blue-grey.leading-airy.mb-15
						| The creator has been rewarded for their creativity and skill, and the world has access to valuable work free of legal restrictions.
					p.text-tiny.text-blue-grey.leading-airy
						| It belongs to everyone now.

				img.explain-box-picture.w-30p(src="@/assets/images/the-parchment-flies-away_touched.svg")


	#join-us.bg-dark-blue-grey.relative
		.container.really-tight.text-white.pt-90.pb-100
			h2.text-center.text-header.font-heavy.mb-30
				| Join Us
			p.text-base.leading-loose.mb-20
				| We're on a mission to make&nbsp;
				out-link(href="https://creativecommonsusa.org/index.php/open-culture/") open knowledge, open culture, and intellectual collaboration
				| &nbsp;the norm.
			p.text-base.leading-loose.mb-40
				| Join the mailing list to hear from us and become a part of this new movement.

			//- SignUpForm(:center, submitText="Submit", assuranceColor="bluey-grey")
			SignUpForm(assuranceClass="text-bluey-grey")

		img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.top-22p.right-30p.h-8p
		img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.top-24p.left-25p.h-6p
		//- img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.top-40p.left-6p.h-8p
		img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.bottom-12p.left-6p.h-8p
		//- img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.top-20p.right-20p.h-8p
		img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.bottom-8p.right-30p.h-6p
		img(src="@/assets/images/cloud.svg").cloud.faded.z-10.absolute.bottom-25p.right-12p.h-8p


	#faq.bg-pale-grey-two.pt-100.pb-100
		.container.medium
			h2.text-center.text-header.text-dark-blue-grey.font-heavy.mb-20
				| FAQ and Resources
			p.text-center.text-base.text-blue-grey.mb-50
				| Something not clear? Want to learn more? Check out these links.

			nuxt-link(to="/faq").faq-link
				span
					| Frequently Asked Questions
				arrow-right-long.arrow

			nuxt-link(to="/why-we-need-it").faq-link
				span
					| Is there really a problem with Intellectual Property?
				arrow-right-long.arrow

			nuxt-link(to="/the-mission").faq-link
				span
					| How could Crowdsell change society?
				arrow-right-long.arrow

			nuxt-link(to="/how-it-works").faq-link
				span
					| How does Crowdsell work in detail?
				arrow-right-long.arrow

			nuxt-link(to="/who-is-crowdsell-for").faq-link
				span
					| Who is Crowdsell for? And what can they do with it?
				arrow-right-long.arrow

			nuxt-link(to="/why-is-crowdsell-different").faq-link
				span
					| How is Crowdsell different than other crowdfunding systems?
				arrow-right-long.arrow

			nuxt-link(to="/how-does-crowdsell-make-money").faq-link
				span
					| How does Crowdsell make money?
				arrow-right-long.arrow

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
			width: 110vw
			max-width: initial

	.faq-link
		@apply bg-white rounded-more px-38 py-30 mb-20 text-dark-blue-grey cursor-pointer
		@apply flex flex-row flex-no-wrap items-center
		span
			@apply text-base font-bold leading-none
			@apply w-80p
		.arrow
			@apply ml-auto w-10p
			@apply text-bright-blue

		&:hover
			@apply bg-white-two text-blue-grey no-underline
			.arrow
				@apply text-violet-blue

	#join-us
		background-image:
			url('../assets/images/ornament.svg'),
			linear-gradient(153deg, theme('colors.navy-blue'), theme('colors.dark-blue-grey'))

</style>
