<template lang="pug">

#app
	Header

	router-view

	//- Footer

</template>

<script>
import Header from '@/layout/Header'
// import Footer from '@/layout/Footer'

export default {
	name: 'App',
	components: {
		Header,
		// Footer,
	},
}
</script>

<style lang="sass">

#app
	-webkit-font-smoothing: antialiased
	-moz-osx-font-smoothing: grayscale

@import './styles/main'

</style>
