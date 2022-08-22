<template>
    <div>
      微信公众号消息
      <ckeditor  @getdata="getData"></ckeditor>
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
      container: "<p>Hello World!</p>",
    };
  },
  methods: {
    getData(data) {
      this.container = data;
    },
  },
 
}
</script>

<style>

</style>