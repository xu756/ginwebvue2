<template>
  <el-container>
    <el-aside width="200px">
      <el-menu :default-active="defaultactive" unique-opened router>
        <div v-for="item in menus" :key="item.name">
          <el-menu-item v-if="item.subMenu == null" :index="item.path"
            ><svg class="icon" aria-hidden="true">
              <use :xlink:href="'#' + item.icon"></use>
            </svg>
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
              <svg class="icon" aria-hidden="true">
                <use :xlink:href="'#' + subitem.icon"></use>
              </svg>
              <span>{{ subitem.name }}</span>
            </el-menu-item>
          </el-submenu>
        </div>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header>
        <el-row>
          <el-col :span="8">左边</el-col>
          <el-col :span="5" :push="11">
            <el-row
              style="
                padding: 12px 0;
                line-height: 36px;
                display: flex;
                justify-content: end;
              "
            >
              <el-col :span="3">
                <el-tooltip
                  class="item"
                  effect="light"
                  content="搜索"
                  placement="bottom"
                >
                  <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-sousuo"></use>
                  </svg>
                </el-tooltip>
              </el-col>

              <el-col :span="3" @click.native="fullscreen">
                <el-tooltip
                  class="item"
                  effect="light"
                  content="全屏"
                  placement="bottom"
                >
                  <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-quanping"></use>
                  </svg>
                </el-tooltip>
              </el-col>
              <el-col
                :span="2"
                @click.native="go('https://github.com/xu756/ginwebvue2')"
              >
                <el-tooltip
                  class="item"
                  effect="light"
                  content="源代码地址"
                  placement="bottom"
                >
                  <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-github"></use>
                  </svg>
                </el-tooltip>
              </el-col>
              <el-col :span="12"
                ><el-dropdown>
                  <span class="el-dropdown-link">
                    <el-avatar
                      :size="30"
                      style="float: left"
                      :src="
                        user.portrait
                          ? user.portrait
                          : 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
                      "
                    ></el-avatar>
                    <div
                      style="
                        line-height: 36px;
                        font-size: 16px;
                        font-weight: bold;
                        float: left;
                      "
                    >
                      {{ user.username }}
                    </div>
                    <i class="el-icon-arrow-down el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>
                      <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-yonghu"></use>
                      </svg>
                      <span @click="go('/admin/menu')">个人中心</span>
                    </el-dropdown-item>
                    <el-dropdown-item>
                      <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-tuichu"></use>
                      </svg>
                      <span @click="logout">退出登录</span>
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown></el-col
              >
            </el-row>
          </el-col>
        </el-row>
      </el-header>
      <el-main><router-view></router-view></el-main>
    </el-container>
  </el-container>
</template>
<script>
import screenfull from "screenfull";
export default {
  data() {
    return {
      defaultactive: "/",
      menus: [],
      user: {},
    };
  },
  mounted() {
    this.isuer();
    this.defaultactive = this.$route.path;
    this.Getdefault();
  },
  methods: {
    isuer() {
      this.$post("IsLogin");
    },
    Getdefault() {
      this.$post("default")
        .then((result) => {
          this.menus = result.data.Menu;
          this.user = result.data.User;
        })
        .catch((err) => {
          this.$message.error(err.message);
        });
    },
    // 退出登录
    logout() {
      this.$post("logout")
        .then((result) => {
          this.$cookies.remove("token");
          this.$router.push("/login");
        })
        .catch((err) => {
          this.$message.error(err.message);
        });
    },
    // 全屏
    fullscreen() {
      screenfull.toggle();
    },
    // 跳转页面
    go(path) {
      // 判断有没有http
      if (path.indexOf("http") == -1) {
        this.$router.push(path);
      } else {
        window.open(path);
      }
    },
  },
};
</script>
 