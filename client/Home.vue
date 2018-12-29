<template lang="pug">

#home
	#banner.container.full-height.flex.align-center.justify-between
		.text-box
			h1#banner-header.display-2
				| Freeing
				br
				| Intellectual
				br
				| Property

		.picture-box
			fa-icon(icon="dove").fa-5x

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

		.picture-box
			fa-icon(icon="dove").fa-5x

	.normal-tile.container.flex.align-center.justify-evenly
		.text-box.limit-40
			h3.display-4
				| We Need
				br
				| Something
				br
				| New

			//- Crowdsell is a crowdfunding system that allows creators to sell their intellectual work to the whole world
			//- p Crowdsell is a new crowdfunding system that adds binding promises to protect backers. Now creators can sell their intellectual work to the whole world, rewarding them for contributing to the commons and the public good.
			p.display-paragraph Crowdsell is a new crowdfunding system that allows creators to fund and sell their intellectual work to the whole world.
			p.display-paragraph By adding binding promises and oversight to protect backers, communities can reward creators for contributing to the commons and the public good.

		.picture-box
			fa-icon(icon="dove").fa-5x

	.container.flex.column.align-center.pad-top-large
		h2.center-text.push-bottom-large.display-3 So how does it work?

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p A creator has a big idea for valuable intellectual work, like code, art, or designs, and needs support and resources to make it happen.
			.picture-box.item
				fa-icon(icon="dove").fa-3x

		.feature-box.limit-40.flex.reverse.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p First they create a proof of concept to show the work is possible.
				p Then they make a detailed list of promises about the intended work's nature and quality, and choose an open culture license.
			.picture-box.item
				fa-icon(icon="dove").fa-3x

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p They make a budget for the project, and choose an amount to act as a reward once it's complete.
				p The whole world can contribute to raise the budget and the reward.
			.picture-box.item
				fa-icon(icon="dove").fa-3x

		.feature-box.limit-40.flex.reverse.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p When the amount is raised, the budget is released month by month, and the community has oversight as the project progresses.
				p When the work is complete and meets the original promises, the creator is also given the reward amount.
			.picture-box.item
				fa-icon(icon="dove").fa-3x

		.feature-box.limit-40.flex.align-center.justify-between.push-bottom-medium
			.text-box.limit-50.item
				p The creator has been rewarded for their creativity and skill, and the community has access to valuable work free of legal restrictions.
				p It belongs to everyone now.
			.picture-box.item
				fa-icon(icon="dove").fa-3x


	.normal-tile.container.flex.align-center.justify-evenly
		.text-box.limit-40
			h2.display-3 Join Us
			p.display-paragraph We're on a mission to make open knowledge, open culture, and intellectual collaboration the norm.
			router-link(to="/sign-up", tag="a").btn.btn-dark Join the Private Beta

		.picture-box
			fa-icon(icon="dove").fa-5x


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
				h4 How does Crowdsell work exactly?

		.page-box
			router-link(to="/why-is-crowdsell-different")
				h4 How is Crowdsell different than other crowdfunding systems?


	footer

</template>

<script>

export default {
	name: 'home',
}

</script>

<style lang="sass" scoped>

// .text-box
// 	font-size: 2rem

.picture-box
	font-size: 3rem

#banner-header
	line-height: 1.6

</style>
