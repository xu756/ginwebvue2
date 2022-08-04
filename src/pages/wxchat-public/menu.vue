<template>
  <div>微信公众号菜单</div>
</template>
<script>
export default {
  data() {
    return {};
  },
  mounted() {
    this.$get("getAccessToken").then((res) => {
      console.log(res);
    });
  },
};
</script>