<template>
  <ckeditor
    :editor="editor"
    v-model="Data"
    :config="editorConfig"
    @ready="onReady"
  ></ckeditor>
</template>
<script>
import DecoupledEditor from "@ckeditor/ckeditor5-build-decoupled-document";
import "@ckeditor/ckeditor5-build-decoupled-document/build/translations/zh-cn"; // 中文包
import UploadAdapter from "@/plugin/upload.js";
export default {
  props: {
    Data: String,
  },
  data() {
    return {
      editor: DecoupledEditor, // 编辑器实例
      editorConfig: {
        language: "zh-cn", // 中文
        toolbar: {
          items: [
            "heading",
            "|",
            "fontfamily",
            "fontsize",
            "|",
            "alignment",
            "|",
            "fontColor",
            "fontBackgroundColor",
            "|",
            "bold",
            "italic",
            "strikethrough",
            "underline",
            "|",
            "link",
            "|",
            "outdent",
            "indent",
            "|",
            "bulletedList",
            "numberedList",
            "|",
            "insertTable",
            "|",
            "uploadImage",
            "blockQuote",
            "|",
            "undo",
            "redo",
          ],
        },
      },
    };
  },
  props: {
    editorData: {
      type: String,
      default: "<p>在这里输入文字....</p>",
    },
  },
  mounted() {
    this.Data = this.editorData;
  },
  watch: {
    Data(val) {
      // 当Data变化时，触发该函数
      this.Data = val;
      this.$emit("getdata", val);
    },
  },
  methods: {
    onReady(editor) {
      editor.ui
        .getEditableElement()
        .parentElement.insertBefore(
          editor.ui.view.toolbar.element,
          editor.ui.getEditableElement()
        );
      editor.plugins.get("FileRepository").createUploadAdapter = (loader) => {
        return new UploadAdapter(loader);
      };
      // 添加插件
    },
  },
};
</script>