<template>
  <div>
    <el-table :data="articles" style="width: 100%" stripe>
      <el-table-column prop="id" label="ID" width="50"> </el-table-column>
      <el-table-column prop="title" label="标题" width="600"> </el-table-column>
      <el-table-column prop="tag" label="标签">
        <template slot-scope="scope">
          <el-tag v-for="item in scope.row.tag" :key="item.id" effect="plain">{{
            item.name
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="category" label="分类"> </el-table-column>
      <el-table-column prop="CreatedAt" label="创建时间">
        <template slot-scope="scope">
          {{ moment(scope.row.CreatedAt).startOf("hour").fromNow() }}
        </template>
      </el-table-column>
      <el-table-column prop="UpdatedAt" label="最后更新时间">
        <template slot-scope="scope">
          {{
            scope.row.UpdatedAt == scope.row.CreatedAt
              ? "-"
              : moment(scope.row.UpdatedAt).startOf("hour").fromNow()
          }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="Edit(scope.row)"
            >编辑</el-button
          >
          <el-button type="text" size="small" @click="Delete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
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
      moment: this.$moment,
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
    // 删除功能
    Delete(row) {
      this.$confirm("此操作将永久删除该文章, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.$post("/delete/article", { id: row.id }).then(() => {
          this.getArticles();
        });
      });
    },
    // 编辑功能
    Edit(row) {
      this.$router.push({
        path: `/article/edit?article_title=${row.title}&article_id=${
          row.id
        }&user_token=${this.$cookies.get("token")}`,
      });
    },
  },
};
</script>

<style>
</style>