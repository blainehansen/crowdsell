<template lang="pug">

#home
	h1 Hello world!
	div
		button(@click="incr") yoyo
	div {{ counter }}


</template>

<script lang="ts">

import { createComponent, ref } from '@marketdial/composition-api'

export default createComponent({
	setup() {
		const counter = ref(0)
		function incr() {
			counter.value++
		}

		return { counter, incr }
	},
})

</script>
