<template>
  <el-container>
    <el-aside width="200px">
      <el-menu :default-active="defaultactive" unique-opened router>
        <div v-for="item in menus" :key="item.name">
          <el-menu-item v-if="item.subMenu == null" :index="item.path"
            ><i :class="'iconfont ' + item.icon"></i>
            <span>{{ item.name }}</span></el-menu-item
          >
          <el-submenu :index="item.path" v-else>
            <template slot="title">
              <i :class="'iconfont ' + item.icon"></i>
              <span>{{ item.name }}</span>
            </template>
            <el-menu-item
              v-for="subitem in item.subMenu"
              :key="subitem.id"
              :index="subitem.path"
            >
              <i :class="'iconfont ' + subitem.icon"></i>
              <span>{{ subitem.name }}</span>
            </el-menu-item>
          </el-submenu>
        </div>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>Header</el-header>
      <el-main><router-view></router-view></el-main>
    </el-container>
  </el-container>
</template>
<script>
export default {
  data() {
    return {
      defaultactive: "/",
      menus: [],
    };
  },
  mounted() {
    // this.isuer();
    this.defaultactive = this.$route.path;
    this.Getdefault();
  },
  methods: {
    isuer() {
      this.$post("IsLogin");
    },
    Getdefault() {
      this.$post("default").then((result) => {
        this.menus = result.data.Menu;
      });
    },
  },
};
</script>
 