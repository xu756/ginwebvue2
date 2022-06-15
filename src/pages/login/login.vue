<template>
  <div>登录页面</div>
</template>
<script>
export default {
  name: 'Login',
  mounted() {
   this.$http.post('/api/login', {
      username: 'admin',
      password: 'admin'
    }).then(response => {
      console.log(response.data);
    });
  },
}
</script>