<template>
  <div>
    <el-table :data="articles" style="width: 100%">
      <el-table-column prop="id" label="ID" width="50"> </el-table-column>
      <el-table-column prop="title" label="标题" width="600"> </el-table-column>
      <el-table-column prop="tag" label="标签">
        <template slot-scope="scope">
          <el-tag v-for="item in scope.row.tag" :key="item.id"  effect="plain">{{
            item.name
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="category" label="分类"> </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :page-sizes="[10, 20, 50, 100]"
      :page-size="page.size"
      :current-page="page.current"
      layout="total, sizes, prev, pager, next, jumper"
      :total="page.total"
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  data() {
    return {
      articles: [],
      page: {
        total: 0, //数据总数
        size: 10,
        current: 1,
      },
    };
  },
  mounted() {
    this.getArticles();
  },
  methods: {
    getArticles() {
      this.$get("/get/articles", this.page).then(({ data }) => {
        this.articles = data.articles;
        this.page.total = data.total;
      });
    },
    // 分页功能
    handleSizeChange(val) {
      this.page.size = val;
      this.getArticles();
    },
    handleCurrentChange(val) {
      this.page.current = val;
      this.getArticles();
    },
  },
};
</script>

<style>
</style>