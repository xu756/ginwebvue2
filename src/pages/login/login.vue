<template>
  <div id="Mainbox">登录页面</div>
</template>
<script>
export default {
  name: "Login",
  mounted() {
    this.$http
      .post("login", {
        username: "admin",
        password: "admin",
      })
      .then((response) => {
        console.log(response.data);
      });
  },
};
</script>
<style lang="scss">
#Mainbox {
  width: 100%;
  height: 100%;
  background-color: #fff;
  text-align: center;
  line-height: 100%;
  color: #000;
}
</style>