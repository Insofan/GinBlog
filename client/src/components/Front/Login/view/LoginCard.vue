<template>

  <div class="card-bg">
	  <v-icon color="#c5c8ce" size="80">account_circle</v-icon>
	  <v-text-field class="input"
			  label="用户名"
			  style="width: 260px"
	  ></v-text-field>
	  <v-text-field class="input"
			  label="密码"
			  style="width: 260px"
	  ></v-text-field>
    <div class="button-bg">
	    <v-btn color="primary" class="button">登录</v-btn>
	    <v-btn color="secondary" to="/" >返回</v-btn>
    </div>
  </div>
</template>

<script>
	import {
		VIcon,
		VTextField,
		VBtn
	} from 'vuetify/lib'

	export default {
		name: 'LoginCard',
		components: {
			VIcon,
			VTextField,
			VBtn
	}
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

    .input {
      margin-top: 25px;
    }

    .button-bg {
      margin-top: 25px;
      .button {
        width: 90px;
        margin-left: 20px;
        margin-right: 20px;
      }
    }

  }
</style>
