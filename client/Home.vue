<template lang="pug">

#home
	#banner.container.full-height.flex.align-center.column.justify-evenly
		.text-box
			h1#banner-header Freeing Intellectual Property

		.picture-box
			fa-icon(icon="dove").fa-5x

	.normal-tile.container.flex.justify-between.pad-top-huge
		.text-box.limit-60
			h3 Intellectual Property is broken.
			p It's stifled our culture and strangled innovation.
			p Big corporations have the power to create monopolies, trap us all in echo chambers, and extract wealth from society.

		.picture-box
			fa-icon(icon="dove").fa-5x

	.normal-tile.container.flex.column.align-center.justify-evenly
		.text-box.limit-40
			h3.text-center We need something new.

			//- Crowdsell is a crowdfunding system that allows creators to sell their intellectual work to the whole world
			//- p Crowdsell is a new crowdfunding system that adds binding promises to protect backers. Now creators can sell their intellectual work to the whole world, rewarding them for contributing to the commons and the public good.
			p Crowdsell is a new crowdfunding system that allows creators to sell their intellectual work to the whole world.
			p By adding binding promises to protect backers, communities can reward creators for contributing to the commons and the public good.

		.picture-box
			fa-icon(icon="dove").fa-5x

	.container.flex.column.align-center.pad-bottom-huge
		h2.center-text.push-bottom-large So how does it work?

		//- .text-box.limit-30.push-bottom-medium.item.self-start
		.text-box.limit-30.push-bottom-medium.item
			p A creator has finished valuable intellectual work, like code, art, or designs, and wants to be rewarded for it.

		.text-box.limit-30.push-bottom-medium
			p They create a preview to show the world why the work is special, and they make a list of promises about the work's nature and quality.

		.text-box.limit-30.push-bottom-medium
			p They choose a price they believe is fair, and the whole world can contribute to raise that amount.

		.text-box.limit-30.push-bottom-medium
			p When the amount is reached, the creator is paid, and the work is released under an open culture license free of legal restrictions.
			p It belongs to everyone now.

	.normal-tile.container
		h2 Join us

		h3 We're on a mission to make open knowledge, open culture, and intellectual collaboration the norm.

		router-link(to="/sign-up", tag="a").btn.btn-dark Join the Private Beta

	.normal-tile.container
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

<style lang="sass">

// .text-box
// 	font-size: 2rem

.picture-box
	font-size: 3rem

#banner-header
	font-size: 3rem


</style>
