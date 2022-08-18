<template>
  <div>
    <el-form  :model="article" label-width="100px">
      <el-form-item label="文章标题：">
        <el-input
          v-model="article.title"
          placeholder="请输入文章标题"
        ></el-input>
      </el-form-item>
      <el-form-item label="文章分类：">
        <el-select v-model="article.category" placeholder="请选择文章分类">
          <el-option v-for="item in categories" :key="item.Id"  :value="item.Name" :label="item.Name"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="文章标签：">
        <el-tag>标签一</el-tag>
<el-tag type="success">标签二</el-tag>
<el-tag type="info">标签三</el-tag>
<el-tag type="warning">标签四</el-tag>
<el-tag type="danger">标签五</el-tag>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
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
    };
  },
  mounted() {
    this.$get("article", { id: this.article.id }).then(({ data }) => {
      this.article = data;
    });
    this.getTags();
  },
  methods: {
    //获取所有标签 分类
    getTags() {
      this.$post("/get/tags/category").then(({ data }) => {
        console.log(data);
        this.tags = data.tags;
        this.categories = data.categorys;
      });
    },
  },
};
</script>

<style>
</style>