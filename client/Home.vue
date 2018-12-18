<template lang="pug">

#home
	#banner.content-tile
		h1 Crowdsell

		h2 Freeing Intellectual Property

	.content-tile
		h3 Intellectual Property is broken.
		p It's made our culture closed and stifled, and gives big corporations the power to create monopolies, close everyone off in walled gardens, and extract wealth from society.

	.content-tile
		h3 We want to fix that.

		p Crowdsell is a crowdfunding system that lets creators sell their intellectual work to the whole world, rewarding them for contributing to the commons and the public good.
		//- , rather than using lawyers and threats to sell copies of it.

	.content-tile
		h3 So how does it work?

		p Let's say a creator has done valuable creative work, like code, art, or designs, and wants to be rewarded for it.

		p They create a preview to show the world why the work is special, and they make a list of promises about the work's nature and quality.

		p They choose a price they believe is fair, and the whole world can contribute to raise that amount.

		p When the amount is reached, the creator is paid the money, and the work is released under an open culture license free of legal restrictions. It belongs to everyone now.

	.content-tile
		h3 Join us on our mission to make open knowledge, open culture, and intellectual collaboration the norm.

		router-link(to="/sign-up", tag="button") Join the Private Beta

	.content-tile
		h3 Learn more

		.page-box
			router-link(to="/why-we-need-it")
				h4 How does Crowdsell work exactly?

		.page-box
			router-link(to="/who-is-crowdsell-for")
				h4 Who is Crowdsell for? And what can they use it to do?

		.page-box
			router-link(to="/why-is-crowdsell-different")
				h4 How is Crowdsell different than other crowdfunding systems?

		.page-box
			router-link(to="/why-we-need-it")
				h4 Are society's Intellectual Property laws and customs really that bad?

		.page-box
			router-link(to="/the-mission")
				h4 What could Crowdsell do to change society?

	footer


</template>

<script>

export default {
	name: 'home',
}

</script>

<style lang="sass">
</style>
