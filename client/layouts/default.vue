<template lang="pug">

#app.antialiased.bg-azure
	Header

	nuxt

	Footer

</template>

<script lang="ts">
import Header from '../layout/Header.vue'
import Footer from '../layout/Footer.vue'

import { Component, Vue, Prop } from 'vue-property-decorator'

@Component({
	components: {
		Header,
		Footer,
	},
})
export default class App extends Vue {}

</script>
