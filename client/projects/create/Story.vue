<template lang="pug">

#story
	MdEditor(v-model="myText")
	p {{ myText }}

</template>


<script>
import MdEditor from '@/components/MdEditor'

export default {
	components: {
		MdEditor
	},
	data() {
		return { myText: '## hello' }
	}
}
</script>


<style lang="sass">
</style>
