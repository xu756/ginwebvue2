<template>
  <el-container>
    <el-aside width="200px">
      <el-menu :default-active="defaultactive" unique-opened router>
        <el-menu-item index="/home">
          <i class="iconfont icon-home1"></i>
          <span>首页</span>
        </el-menu-item>
        <el-submenu index="/WxchatPublic">
          <template #title>
            <i class="iconfont icon-gongzhonghaoguanli"></i>
            <span>公众号</span>
          </template>
          <el-menu-item index="/WxchatPublic/menu">
            <i class="iconfont icon-caidan"></i>
            <span>自定义菜单</span>
          </el-menu-item>
          <el-menu-item index="/WxchatPublic/message">
            <i class="iconfont icon-xiaoxitongzhi"></i>
            <span>消息回复</span>
          </el-menu-item>
        </el-submenu>
        <el-menu-item index="/user/info">
          <i class="iconfont icon-yonghu"></i>
          <span>个人信息</span>
        </el-menu-item>
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
  name: "Login",
  data() {
    return {
      defaultactive: "/",
      loading: true,
    };
  },
  mounted() {
    this.isuer();
    this.defaultactive = this.$route.path;
  },
  methods: {
    isuer() {
      this.$post("IsLogin").then((res) => {
        if (!res.is) {
          this.$router.push("/login");
        }
      });
    },
  },
};
</script>
