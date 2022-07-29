<template>
  <div class="rich-editor">
    <div id="toolbar-container" />
    <div id="editor"></div>
  </div>
</template>
<script>
import DecoupledEditor from "@ckeditor/ckeditor5-build-decoupled-document";
import "@ckeditor/ckeditor5-build-decoupled-document/build/translations/zh-cn"; // 中文包
import UploadAdapter from "@/plugin/upload.js";
export default {
  data() {
    return {
      editor: "555", // 编辑器实例
    };
  },
  mounted() {
    this.initCKEditor();
  },
  methods: {
    initCKEditor() {
      // 初始化编辑器
      DecoupledEditor.create(document.querySelector("#editor"), {
        language: "zh-cn", // 中文
       editorConfig:{
        
       }
      })
        .then((editor) => {
          editor.plugins.get("FileRepository").createUploadAdapter = (
            loader
          ) => {
            return new UploadAdapter(loader);
          };
          const toolbarContainer = document.querySelector("#toolbar-container");
          toolbarContainer.appendChild(editor.ui.view.toolbar.element);
          this.editor = editor; // 将编辑器保存起来，用来随时获取编辑器中的内容等，执行一些操作
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
<style lang="scss">
</style>

