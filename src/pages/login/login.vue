<template>
  <div id="Mainbox">
    <div id="bg"></div>
  </div>
</template>
<script>
export default {
  name: "Login",
  mounted() {
    this.$post
      ("login", {
        username: "admin",
        password: "admin",
      })
      .then((res) => {
        console.log(res);
      });
  },
};
</script>
<style lang="scss">
#Mainbox {
  width: 100%;
  height: 100%;
  text-align: center;

  color: #000;
  #bg {
    width: 100%;
    height: 100%;
    background-image: linear-gradient(135deg, #8bc6ec 0%, #9599e2 100%);
  }
}
</style>