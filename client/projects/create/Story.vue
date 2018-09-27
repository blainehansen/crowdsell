<template lang="pug">

#story
	//- MdEditor(v-model="myText")
	p {{ myText }}

</template>


<script>
export default {
	// components: {
	// 	MdEditor
	// },
	data() {
		return { myText: '## hello' }
	}
}
</script>


<style lang="sass">
</style>
