<template>
  <el-form :model="form">
    <el-form-item label="用户名"></el-form-item>
  </el-form>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
    };
  },
};
</script>
<style lang="scss">
</style>