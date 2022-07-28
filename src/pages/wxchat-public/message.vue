<template>
  <div class="quill-editor">
    <!-- 图片上传组件辅助-->
    <el-upload class="avatar-uploader" :action="uploadUrl" name="img" with-credentials :show-file-list="false" :on-success="uploadSuccess" :before-upload="beforeUpload">
    </el-upload>
    <!--富文本编辑器组件-->
    <quill-editor v-model="content" :content="content" :options="editorOption" @blur="onEditorBlur($event)" @focus="onEditorFocus($event)" @ready="onEditorReady($event)" ref="QuillEditor">
    </quill-editor>
  </div>
</template>

<script>
import { quillEditor } from 'vue-quill-editor'
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import 'quill/dist/quill.bubble.css';
import resizeImage from 'quill-image-resize-module' // 图片缩放组件引用
import { ImageDrop } from 'quill-image-drop-module'; // 图片拖动组件引用
Quill.register('modules/imageDrop', ImageDrop); // 注册
Quill.register('modules/resizeImage ', resizeImage ) // 注册

const toolbarOptions = [
  ['bold', 'italic', 'underline', 'strike'],        // 加粗，斜体，下划线，删除线
  ['blockquote', 'code-block'],                      //引用，代码块
  [{ 'header': 1 }, { 'header': 2 }],               // 几级标题
  [{ 'list': 'ordered' }, { 'list': 'bullet' }],     // 有序列表，无序列表
  [{ 'script': 'sub' }, { 'script': 'super' }],      // 下角标，上角标
  [{ 'indent': '-1' }, { 'indent': '+1' }],          // 缩进
  [{ 'direction': 'rtl' }],                         // 文字输入方向
  [{ 'size': ['small', false, 'large', 'huge'] }],  // 字体大小
  [{ 'header': [1, 2, 3, 4, 5, 6, false] }],// 标题
  [{ 'color': [] }, { 'background': [] }],          // 颜色选择
  [{ 'font': ['SimSun', 'SimHei', 'Microsoft-YaHei', 'KaiTi', 'FangSong', 'Arial'] }],// 字体
  [{ 'align': [] }], // 居中
  ['clean'],                                       // 清除样式,
  ['link', 'image'],  // 上传图片、上传视频
]
export default {
  components: {
    quillEditor
  },
  data () {
    return {
      name: 'register-modules-example',
      content: '',
      editorOption: {
        placeholder: '请在这里输入',
        theme: 'snow', //主题 snow/bubble
        modules: {
          history: {
            delay: 1000,
            maxStack: 50,
            userOnly: false
          },
          toolbar: {
            container: toolbarOptions,
            handlers: {
              image: function (value) {
                if (value) {
                  // 调用element的图片上传组件
                  document.querySelector('.avatar-uploader input').click()
                } else {
                  this.quill.format('image', false)
                }
              }
            }
          },
           imageDrop: true,      //图片拖拽
          imageResize: {          //放大缩小
            displayStyles: {
              backgroundColor: "black",
              border: "none",
              color: "white"
            },
            modules: ["Resize", "DisplaySize", "Toolbar"]
          },
        }
      },
      uploadUrl: `/upload/img` 	// 服务器上传地址
    }
  },
  methods: {
    // 失去焦点
    onEditorBlur (editor) { },
    // 获得焦点
    onEditorFocus (editor) { },
    // 开始
    onEditorReady (editor) { },
    // 值发生变化
    onEditorChange (editor) {
      this.content = editor.html;
    },
    beforeUpload (file) { 
        // 判断上传文件的类型
        const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
        if (!isJpgOrPng) {
            this.$message.error('上传图片只能是 JPG/PNG 格式!')
        }
        const isLt2M = file.size / 1024 / 1024 < 2
        if (!isLt2M) {
            this.$message.error('上传图片大小不能超过 2MB!')
        }
        return isJpgOrPng && isLt2M
        
    },
    uploadSuccess (res) {
      // 获取富文本组件实例
      let quill = this.$refs.QuillEditor.quill
      // 如果上传成功
      if (res) {
        console.log(res)
        // 获取光标所在位置
        let length = quill.getSelection().index;
        // 插入图片，res为服务器返回的图片链接地址
        quill.insertEmbed(length, 'image', res.data.url)
        // 调整光标到最后
        quill.setSelection(length + 1)
      } else {
        // 提示信息，需引入Message
        this.$message.error('图片插入失败！')
      }
    }
  }
}
</script>
