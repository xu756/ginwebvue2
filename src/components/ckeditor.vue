<template>
  <ckeditor
    :editor="editor"
    v-model="context"
    :config="editorConfig"
    @ready="onReady"
  ></ckeditor>
</template>
<script>
import DecoupledEditor from "@ckeditor/ckeditor5-build-decoupled-document";
import "@ckeditor/ckeditor5-build-decoupled-document/build/translations/zh-cn"; // 中文包
import UploadAdapter from "@/plugin/upload.js";
export default {
  data() {
    return {
      context: "",
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
    Data: {
      type: String,
      default: "<p>在这里输入文字....</p>",
    },
  },
  mounted() {
    this.context = this.Data;
  },
  watch: {
    context(val) {
      this.$emit("getdata", val);
    },
    Data(val) {
      this.context = val;
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