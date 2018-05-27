<template lang="pug">

//- title
//- blurb
//- image
//- video?
//- categories and tags
//- funding goal
//- funding timeline and structure
//- add people to project
//- description
//- the files themselves

//- the demo materials
//- the release guarantees
//- the license

//- financial information

</template>

<script>
import api from '@/api'

export default {
	name: 'CreateProject',
	asyncData: {
		projects: () => api.getProjects()
	}
}
</script>

<style lang="sass">

</style>
