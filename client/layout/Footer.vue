<template lang="pug">

//- footer.pad-top-medium.pad-bottom-small.bg-dark
footer.pad-top-small.pad-bottom-small
	.container
		.links.text-muted.row
			//- .col-12
			//- 	h5 Resources

			//- .col-md-5
			//- 	.page-box
			//- 		nuxt-link(to="/faq")
			//- 			span Frequently Asked Questions

			//- 	.page-box
			//- 		nuxt-link(to="/why-we-need-it")
			//- 			span Is there really a problem with Intellectual Property?

			//- 	.page-box
			//- 		nuxt-link(to="/the-mission")
			//- 			span How could Crowdsell change society?

			//- 	.page-box
			//- 		nuxt-link(to="/how-it-works")
			//- 			span How does Crowdsell work in detail?

			//- .col-md-5
			//- 	.page-box
			//- 		nuxt-link(to="/who-is-crowdsell-for")
			//- 			span Who is Crowdsell for? And what can they do with it?

			//- 	.page-box
			//- 		nuxt-link(to="/why-is-crowdsell-different")
			//- 			span How is Crowdsell different than other crowdfunding systems?

			//- 	.page-box
			//- 		nuxt-link(to="/how-does-crowdsell-make-money")
			//- 			span How does Crowdsell make money?
			//- two columns
			.col-12
				h3.push-bottom-less Join the Mailing List
				SignUpForm

			.col-12.flex.justify-between.align-center.push-top-small

				.item
					//- out-link.push-right-less(href="https://twitter.com/crowdsell_io")
					//- 	twitter(width="1.25rem", height="1.25rem", href="")
					a(href="mailto:contact@crowdsell.io") contact@crowdsell.io

				.item
					nuxt-link(to="/faq")
						span FAQ and Resources

				.item.license.text-right.text-muted
					| All content licensed as
					br
					a(rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/")
						| Creative Commons Attribution-ShareAlike
					br
					a(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
						img(alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/80x15.png")

</template>


<script>
export default {
	name: 'Footer'
}

</script>


<style lang="sass" scoped>

</style>
