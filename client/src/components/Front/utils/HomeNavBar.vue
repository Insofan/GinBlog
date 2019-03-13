<template>
  <div class="layout">
	  <v-navigation-drawer
			  v-model="isDrawer"
			  fixed
			  app
	  >
		  <v-list dense>
			  <v-list-tile @click="" to="/">
				  <v-list-tile-action>
					  <v-icon>home</v-icon>
				  </v-list-tile-action>
				  <v-list-tile-content>
					  <v-list-tile-title>Home</v-list-tile-title>
				  </v-list-tile-content>
			  </v-list-tile>
			  <v-list-tile @click="" to="login">
				  <v-list-tile-action>
					  <v-icon>account_box</v-icon>
				  </v-list-tile-action>
				  <v-list-tile-content>
					  <v-list-tile-title>
							  登录
					  </v-list-tile-title>
				  </v-list-tile-content>
			  </v-list-tile>
		  </v-list>
	  </v-navigation-drawer>
	  <v-toolbar color="primary" light fixed app>
		  <v-toolbar-side-icon @click.stop="isDrawer = !isDrawer" style="color: white"></v-toolbar-side-icon>
		  <v-toolbar-title>
			  <div style="color: white; font-size: 24px">
				  Insomnia's World
			  </div>
		  </v-toolbar-title>
	  </v-toolbar>
  </div>
</template>

<script>
	import {VApp,
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
	} from 'vuetify/lib'
export default {
  name: 'HomeNavBar',
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
	},
	data() {
		return {
			isDrawer: false,
		}
	},
}
</script>

<style scoped lang="scss">
  .nav-bg {
    background: $primary-color;
    min-width: 1200px;
    height: 8vh;
    min-height: 60px;
    .nav-title-label {
      @include flex-row-center;
      @include title-font;
      color: white;
      font-size: 4vw;
      font-weight: bold;
      overflow: hidden;
    }
    .menu {
      background: $primary-color;
      .menu-item {
        color: white;
      }
    }
  }
</style>
