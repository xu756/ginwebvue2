<template>
  <div id="Menu">
    <h1>菜单设置</h1>
    <el-table
      :data="menus"
      style="width: 100%"
      row-key="id"
      border
      stripe
      :tree-props="{ children: 'subMenu' }"
    >
      <el-table-column prop="id" label="id" width="80"> </el-table-column>
      <el-table-column
        prop="name"
        label="菜单名称"
        width="180"
      ></el-table-column>
      <el-table-column prop="path" label="路径" width="300"></el-table-column>
      <el-table-column prop="icon" label="图标" width="100" align="center">
        <template slot-scope="scope">
          <i :class="'iconfont ' + scope.row.icon"></i>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      menus: [],
    };
  },
  mounted() {
    this.getMenus();
  },
  methods: {
    getMenus() {
      this.$get("/get/menu").then(({ data }) => {
        console.log(data);
        this.menus = data;
      });
    },
  },
};
</script>

<style lang="scss">
</style>