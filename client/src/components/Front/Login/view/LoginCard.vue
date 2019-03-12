<template>
  <div class="card-bg">
    <div class="username-label">
      用户名
    </div>
    <Input class="username-input" v-model="value" placeholder="用户名..." style="width: 300px" />
    <div class="password-label">
      密码
    </div>
    <Input class="password-input" v-model="value" placeholder="密码..." style="width: 300px" />
    <Button type="primary" class="login-button">登录</Button>
  </div>
</template>

<script>
export default {
  name: 'LoginCard'
}
</script>

<style scoped lang="scss">
  .card-bg {
    @include flex-col-center;
    background: $background-color;
    min-height: 300px;
    min-width: 400px;
    box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12), 0 2px 4px -1px rgba(0, 0, 0, 0.2);
    padding: 50px;

    .username-label {
      @include title-font;
      align-self: flex-start;
    }
    .username-input {
      margin-top: 25px;
    }
    .password-label {
      margin-top: 25px;
      @include title-font;
      align-self: flex-start;
    }
    .password-input {
      margin-top: 25px;
    }
    .login-button {
      margin-top: 25px;
      width: 100px;
      height: 40px;
    }
  }
</style>
