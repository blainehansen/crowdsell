<template lang="pug">

#sign-up.h-screen.md_flex.md_flex-row.md_flex-no-wrap.md_items-stretch

	#mailing-list.md_w-40p.flex.flex-col.relative.justify-center.items-center.text-left.py-40.md_py-0
		.thing.md_w-50p
			nuxt-link(to="/").logo.mb-60.flex.flex-row.flex-no-wrap.justify-start.items-center.cursor-pointer.text-bright-blue.hover_text-blue-blue.hover_no-underline
				Logo.mr-10.fill-current
				span.font-bold.leading-none.text-tricky-logo.mt-10 Crowdsell

			h2.text-heavy.text-dark-blue-grey.font-heavy.mb-24 Mailing List

			SignUpForm(:separateMode="true")

		.md_absolute.bottom-5p.flex.flex-row.flex-no-wrap.justify-between.items-center.mt-40.md_mt-0
			span.text-bluey-grey.text-tiny.font-heavy Stay In Touch!

			.flex.flex-row.flex-no-wrap.justify-start.items-center.text-tiny.ml-50
				out-link(href="https://twitter.com/crowdsell_io").text-bright-blue.hover_text-white.hover_no-underline
					twitter(width="1.25rem", height="1.25rem").fill-current
				out-link(href="mailto:contact@crowdsell.io").text-bright-blue.hover_no-underline.ml-5
					| contact@crowdsell.io

	#pretty-box.md_w-60p.flex.flex-col.relative.text-center.justify-center.items-center.py-40.md_py-0
		img(src="@/assets/images/freeing-intellectual-property_pieces.svg").w-50p.mb-10p

		img(src="@/assets/images/dots.svg").dots.absolute.w-20p.left-5p.top-12p
		img(src="@/assets/images/dots.svg").dots.absolute.w-18p.right-6p.top-4p
		img(src="@/assets/images/dots.svg").dots.absolute.w-14p.bottom-14p.right-14p

		img(src="@/assets/images/cloud.svg").cloud.absolute.top-45p.left-32p.h-10p
		img(src="@/assets/images/cloud.svg").cloud.absolute.top-40p.right-30p.h-8p

		p.text-white.w-60p
			i Join the Private Beta to be one of the first to launch or contribute to a project.
		p.text-white.w-60p
			i We’re trying to make the world an open culture place, and we’re looking forward to working together.

		.flex.flex-col.md_flex-row.flex-no-wrap.justify-between.w-80p.md_absolute.bottom-5p.mt-40.md_mt-0
			nuxt-link(to="/faq").text-white.underline.text-tiny.hover_text-white.mb-20.md_mb-0
				| FAQ and Resources

			.flex.flex-col.md_flex-row.flex-no-wrap.justify-start.items-center
				out-link(rel="license", href="http://creativecommons.org/licenses/by-sa/4.0/")
					img(alt="Creative Commons License", style="border-width:0", src="https://i.creativecommons.org/l/by-sa/4.0/80x15.png")

				.text-tiny.mt-5.md_mt-0.md_ml-15
					span.text-white All content licensed as&nbsp;
					out-link(href="http://creativecommons.org/licenses/by-sa/4.0/").text-white.underline.hover_text-white
						| Creative Commons Attribution-ShareAlike

</template>

<script>

export default {
	layout: 'bare',
}

</script>

<style lang="sss">

#pretty-box
	background-image:
		url('../assets/images/ornament.svg'),
		linear-gradient(218deg, theme('colors.violet-blue'), theme('colors.blue-blue-two'))

</style>
