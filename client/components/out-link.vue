<template lang="pug">

a(:href="href", target="_blank")
	slot

</template>

<script>

export default {
	props: {
		href: {
			type: String,
			required: true,
		},
	}
}

</script>
