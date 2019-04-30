<template lang="pug">

#profile


</template>


<script>

export default {
	name: 'profile',

	props: {
		userSlug: {
			type: String,
			default: null,
		},
		previewing: {
			type: Boolean,
			default: false,
		}
	},
}
</script>


<style lang="sass">
</style>
