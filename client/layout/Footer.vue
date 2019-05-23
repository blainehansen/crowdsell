<template lang="pug">

footer.bg-navy-blue
	.container
		#footer-links
			.logo

			.contact
				out-link(href="https://twitter.com/crowdsell_io").text-white
					twitter(width="1.25rem", height="1.25rem").fill-current
				out-link(href="mailto:contact@crowdsell.io").text-white contact@crowdsell.io

			//- nuxt-link(to="/faq").text-white.underline.hover:ice-blue FAQ and Resources
			nuxt-link(to="/faq").text-white.underline FAQ and Resources

			//- a(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
			//- some flex container
			out-link(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
				img(alt="Creative Commons License", style="border-width:0", src="https://i.creativecommons.org/l/by-sa/4.0/80x15.png")

			.text-tiny
				span.text-white All content licensed as&nbsp;
				//- out-link(href="http://creativecommons.org/licenses/by-sa/4.0/").text-bright-blue.underline.hover:text-light-teal
				out-link(href="http://creativecommons.org/licenses/by-sa/4.0/").text-bright-blue.underline
					| Creative Commons Attribution-ShareAlike


		#footer-mailing-list.text-white
			h3.text-header.font-bold
				| Join the Mailing List
			p.text-base
				| Get the latest updates & be the first to know when we officially launch.

			SignUpForm

</template>
