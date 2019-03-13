<template>
	<v-app>
    <div class="login-bg">
      <div>
        <login-card></login-card>
      </div>
    </div>
	</v-app>
</template>

<script>
import LoginCard from './view/LoginCard'
import {VApp} from 'vuetify/lib'
export default {
  name: 'LoginController',
  components: {LoginCard,
  VApp}
}
</script>

<style scoped lang="scss">
  .login-bg {
      @include flex-row-center;
      height: 100vh;
      background: linear-gradient(0deg, white 50%, $primary-color 50%);
  }
</style>
