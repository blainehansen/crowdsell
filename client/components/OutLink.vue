<template lang="pug">

a(:href="href", target="_blank")
	slot

</template>

<script>

export default {
	name: 'out-link',
	props: {
		href: {
			type: String,
			required: true,
		},
	}
}

</script>
