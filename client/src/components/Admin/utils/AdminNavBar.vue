<template>
	<div>
		<v-navigation-drawer
				v-model="isDrawer"
				fixed
				clipped
				app
		>
		</v-navigation-drawer>
		<v-toolbar
				color="primary"
				dense
				fixed
				clipped-left
				app
		>
			<v-toolbar-side-icon @click.stop="isDrawer = !isDrawer" style="color: white">
			</v-toolbar-side-icon>
			<v-icon class="mx-3">fab fa-youtube</v-icon>
			<v-toolbar-title class="mr-5 align-center">
				<span class="title" style="color: white">Insomnia's Admin</span>
			</v-toolbar-title>
			<v-spacer></v-spacer>
			<v-layout row align-center style="max-width: 350px">
				<v-text-field class="search-text-field  lighten-1"
				              placeholder="搜索..."
				              single-line
				              append-icon="search"
				              color="white"
				              hide-details
				></v-text-field>
			</v-layout>
		</v-toolbar>
	</div>
</template>

<script>
	import {
		VApp,
		VIcon,
		VList,
		VListTile,
		VListTileAction,
		VListTileContent,
		VListTileTitle,
		VNavigationDrawer,
		VToolbar,
		VToolbarTitle,
		VToolbarSideIcon,
		VTextField,
		VLayout,
		VSpacer
	} from 'vuetify/lib'

	export default {
		name: 'AdminNavBar',
		components: {
			VApp,
			VIcon,
			VNavigationDrawer,
			VToolbar,
			VToolbarTitle,
			VToolbarSideIcon,
			VList,
			VListTile,
			VListTileAction,
			VListTileContent,
			VListTileTitle,
			VTextField,
			VLayout,
			VSpacer
		},
		data() {
			return {
				isDrawer: false
			}
		}
	}
</script>

<style scoped lang="scss">

</style>
