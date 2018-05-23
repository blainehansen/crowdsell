<template lang="pug">

header
	.login
		template(v-if="userLoggedIn")
			p Current User: {{ userName }}
			a.btn.btn-link(@click="logout") Logout
		template(v-else)
			router-link(to="login") Login

</template>


<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
	name: 'Header',

	computed: {
		...mapGetters(['userName', 'userLoggedIn'])
	},

	methods: {
		...mapMutations(['logout'])
	},
}

</script>


<style lang="sass">


</style>
