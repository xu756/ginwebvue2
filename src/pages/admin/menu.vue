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
          <svg class="icon" aria-hidden="true">
            <use :xlink:href="'#' + scope.row.icon"></use>
          </svg>
        </template>
      </el-table-column>
      <el-table-column label="编辑">
        <template slot-scope="scope">
          <el-button type="text" size="mini" @click="Edit(scope.row)">
            编辑
          </el-button>

          <el-button type="text" size="mini" @click="Delete(scope.row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose"
    >
      <span slot="title">{{ "修改" + edit.name + "菜单项" }}</span>
      <el-form :model="edit" label-position="left" ref="editRef">
        <el-row>
          <el-col :span="9" :push="2"
            ><el-form-item label="ID" label-width="50px">
              <el-input
                v-model="edit.id"
                disabled
                placeholder="不可修改"
              ></el-input> </el-form-item
          ></el-col>
          <el-col :span="8" :push="4"
            ><el-form-item label="路径" label-width="50px">
              <el-input
                v-model="edit.path"
                disabled
                placeholder="不可修改"
              ></el-input> </el-form-item
          ></el-col>
        </el-row>
        <el-row>
          <el-col :span="9" :push="2"
            ><el-form-item label="名称" label-width="50px">
              <el-input
                v-model="edit.name"
                placeholder="请输入菜单名称"
              ></el-input> </el-form-item
          ></el-col>
          <el-col :span="8" :push="4"
            ><el-form-item label="图标" label-width="50px">
              <el-input
                v-model="edit.icon"
                placeholder="请输入图标Symbol"
              ></el-input> </el-form-item
          ></el-col>
        </el-row>
        <el-row>
          <el-col :span="5" :push="15">
            <el-button type="primary" @click="EditMenu">提交</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      menus: [],
      dialogVisible: false, //对话框是否显示
      edit: {
        id: 0,
        name: "",
        path: "",
        icon: "",
      },
    };
  },
  mounted() {
    this.getMenus();
  },
  methods: {
    getMenus() {
      this.$get("/get/menu").then(({ data }) => {
        this.menus = data;
      });
    },
    // 编辑
    Edit(row) {
      this.dialogVisible = true;
      this.$nextTick(() => {
        this.edit = row;
      });
    },
    // 删除
    Delete(row) {
      console.log(row);
    },
    // 关闭
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((e) => {
          done();
          this.$refs.editRef.resetFields();
          console.log(this.edit);
        })
        .catch((_) => {});
    },
    EditMenu() {
      console.log(this.edit.name);
      this.dialogVisible = false;
    },
  },
};
</script>
