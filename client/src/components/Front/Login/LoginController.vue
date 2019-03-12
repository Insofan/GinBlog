<template>
  <Layout>
    <div class="login-bg">
      <div>
        <login-card></login-card>
      </div>
    </div>
  </Layout>
</template>

<script>
import LoginCard from './view/LoginCard'
export default {
  name: 'LoginController',
  components: {LoginCard}
}
</script>

<style scoped lang="scss">
  .login-bg {
      @include flex-row-center;
      height: 100vh;
      background: linear-gradient(0deg, white 50%, $info-color 50%);
  }
</style>
