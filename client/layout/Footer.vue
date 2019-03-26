<template lang="pug">

//- footer.pad-top-medium.pad-bottom-small.bg-dark
footer.pad-top-small.pad-bottom-small
	.container
			//- two columns
			.flex.justify-between.align-center
				.item
					out-link.push-right-less(href="https://twitter.com/crowdsell_io")
						twitter(width="1.25rem", height="1.25rem", href="")
					a(href="mailto:contact@crowdsell.io") contact@crowdsell.io

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

@import "~styles/variables"

footer
	font-size: 0.80rem

	// border-top: 5px solid $gray

	background-color: $gray
	color: $primary-almost-white

	a
		color: $color-primary-0

</style>
