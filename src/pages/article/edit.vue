<template>
  <div>
    <el-form :model="article" label-width="100px">
      <el-form-item label="文章标题：">
        <el-input
          v-model="article.title"
          placeholder="请输入文章标题"
        ></el-input>
      </el-form-item>
      <el-form-item label="文章分类：">
        <el-select v-model="article.category" placeholder="请选择文章分类">
          <el-option
            v-for="item in categories"
            :key="item.Id"
            :value="item.Name"
            :label="item.Name"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="文章标签：">
        <el-tag
          v-for="tag in article.tag"
          :key="tag.name"
          :disable-transitions="false"
          closable
          @close="handleClose(tag)"
        >
          {{ tag.name }}
        </el-tag>
        <el-input
          class="input-new-tag"
          v-if="inputVisible"
          v-model="inputValue"
          ref="saveTagInput"
          size="small"
          @keyup.enter.native="handleInputConfirm"
          @blur="handleInputConfirm"
        >
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput"
          >+ New Tag</el-button
        >
      </el-form-item>
      <el-form-item label="内容：">
        <ckeditor @getdata="getData" :Data="article.content"></ckeditor>
      </el-form-item>
      <el-form-item label="提交">
        <el-button @click="submit">提交</el-button>
      </el-form-item>
    </el-form>
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
      article: {
        id: 1,
        title: "",
        content: "",
        category: "",
        author: "",
        tag: [],
      },
      tags: [], // 标签列表
      categories: [], // 分类列表
      inputVisible: false, // 正在输入
      inputValue: "", // 新标签
    };
  },
  created() {
    this.$get("article", { id: this.article.id }).then(({ data }) => {
      this.article = data;
    });
    this.getTags();
  },
  methods: {
    //获取所有标签 分类
    getTags() {
      this.$post("/get/tags/category").then(({ data }) => {
        this.tags = data.tags;
        this.categories = data.categorys;
      });
    },
    // 移除标签
    handleClose(tag) {
      this.article.tag.splice(this.article.tag.indexOf(tag), 1);
    },
    //  添加标签
    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    //回车添加标签
    handleInputConfirm() {
      let inputValue = {};
      inputValue.name = this.inputValue;
      if (inputValue.name) {
        this.article.tag.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    //
    getData(data) {
      this.article.content = data;
    },
    // 提交
    submit() {
      console.log(this.article.content);
    },
  },
};
</script>