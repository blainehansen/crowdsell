<template lang="pug">

footer.bg-navy-blue
	.container.pt-60.pb-30.md_py-150.flex.flex-col-reverse.md_flex-row.flex-no-wrap.md_justify-between
		#footer-links.md_w-50p.mt-80.md_mt-0
			nuxt-link(to="/").inline-flex.flex-row.flex-no-wrap.justify-start.items-center.cursor-pointer.text-white.hover_text-light-teal.hover_no-underline.mb-30.md_mb-90
				Logo.mr-15.fill-current.w-tricky-big-logo.h-auto
				span.font-bold.leading-none.text-tricky-big-logo.mt-15 Crowdsell

			.contact.flex.flex-row.flex-no-wrap.justify-start.items-center.text-small.mb-10
				out-link(href="https://twitter.com/crowdsell_io").text-white.hover_text-ice-blue.hover_no-underline
					twitter(width="1.25rem", height="1.25rem").text-white.fill-current
				out-link(href="mailto:contact@crowdsell.io").text-white.hover_text-ice-blue.hover_no-underline.ml-15
					| contact@crowdsell.io

				nuxt-link(to="/faq").text-white.hover_text-ice-blue.hover_no-underline.text-small.ml-15.md_ml-100
					| FAQ and Resources

			//- a(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
			.flex.flex-row.flex-no-wrap.justify-start.items-center.mt-30.md_mt-70
				out-link(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
					img(alt="Creative Commons License", style="border-width:0", src="https://i.creativecommons.org/l/by-sa/4.0/80x15.png")

				.text-tiny.ml-15
					span.text-white All content licensed as&nbsp;
					out-link(href="http://creativecommons.org/licenses/by-sa/4.0/").text-bright-blue.hover_text-light-teal
						| Creative Commons Attribution-ShareAlike


		#footer-mailing-list.text-white.md_w-40p
			h3.text-header.font-heavy.mb-10
				| Join the Mailing List
			p.text-small.mb-50
				| Get the latest updates & be the first to know when we officially launch.

			SignUpForm

</template>
