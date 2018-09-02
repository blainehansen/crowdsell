<template lang="pug">

#app
	Header

	router-view

	Footer

</template>

<script>
import Header from '@/layout/Header'
import Footer from '@/layout/Footer'

import promisepay from 'promisepay'

export default {
	name: 'App',
	components: {
		Header, Footer
	},
}
</script>

<style lang="sass">

#app
	font-family: 'Avenir', Helvetica, Arial, sans-serif
	-webkit-font-smoothing: antialiased
	-moz-osx-font-smoothing: grayscale
	text-align: center

@import './styles/main'

</style>
