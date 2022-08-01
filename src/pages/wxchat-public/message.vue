<template>
    <div>
      微信公众号消息
      <ckeditor :editorData="container" ></ckeditor>
    </div>
</template>

<script>
import ckeditor from "@/components/ckeditor.vue";
export default {
  components: {
    ckeditor,
  },
  data() {
    return {
      container: "8585",
    };
  },
}
</script>

<style>

</style>