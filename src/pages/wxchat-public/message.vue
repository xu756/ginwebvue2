<template>
  <div>
    <ckeditor
      :editor="editor"
      v-model="editorData"
      :config="editorConfig"
      @ready="onReady"
    ></ckeditor>
  </div>
</template>
<script>
import DecoupledEditor from "@ckeditor/ckeditor5-build-decoupled-document";
import "@ckeditor/ckeditor5-build-decoupled-document/build/translations/zh-cn"; // 中文包
import UploadAdapter from "@/plugin/upload.js";
export default {
  data() {
    return {
      editor: DecoupledEditor, // 编辑器实例
      editorData: "", // 编辑器中的内容
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
  methods: {
    onReady(editor) {
      // Insert the toolbar before the editable area.
      editor.ui
        .getEditableElement()
        .parentElement.insertBefore(
          editor.ui.view.toolbar.element,
          editor.ui.getEditableElement()
        );
      editor.plugins.get("FileRepository").createUploadAdapter = (loader) => {
        return new UploadAdapter(loader);
        
      };
    },
  },
};
</script>