<template>
  <div>页面</div>
</template>
<script>
export default {
  name: "Login",
  created() {
    this.$http
      .post("/api/login", {
        username: "admin",
        password: "admin",
      })
      .then((response) => {
        console.log(response.data);
      });
  },
};
</script>