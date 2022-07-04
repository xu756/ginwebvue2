<template>
  <div>{{ this.$store.state.userInfo.token }}</div>
</template>
<script>
export default {
  name: "Login",
  mounted() {
    this.$post("IsLogin").then((res) => {
      if (!res.is) {
        this.$router.push("/login");
      }
    });
  },
};
</script>