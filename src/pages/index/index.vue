<template>
  <el-container>
    <el-aside width="200px">
      <el-menu :default-active="defaultactive"  router background-color="#F6FBF4" active-text-color="#47B5FF">
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
                        padding-left: 10px;
                        float: left;
                      "
                    >
                      {{user.username }}
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
      <el-main>
        <el-tabs
          v-model="activeIndex"
          type="card"
          closable
          v-if="openTab.length"
          @tab-click="tabClick"
          @tab-remove="tabRemove"
        >
          <el-tab-pane
            :key="item.name"
            v-for="item in openTab"
            :label="item.title"
            :name="item.route"
          >
          </el-tab-pane>
        </el-tabs>
        <router-view></router-view
      ></el-main>
    </el-container>
  </el-container>
</template>
<script>
import screenfull from "screenfull";
export default {
  data() {
    return {
      defaultactive: "/home",
      menus: [],
      user: {},
    };
  },
  mounted() {
    this.isuer();
    this.defaultactive = this.$route.path;
    this.Getdefault();
    // 刷新时以当前路由做为tab加入tabs
    // 当前路由不是首页时，添加首页以及另一页到store里，并设置激活状态
    // 当当前路由是首页时，添加首页到store，并设置激活状态
    if (this.$route.path !== "/home") {
      this.$store.commit("add_tabs", {
        route: "/home",
        name: "home",
        title: "首页",
      });
      this.$store.commit("add_tabs", {
        route: this.$route.path,
        name: this.$route.name,
        title: this.$route.meta.title,
      });
      this.$store.commit("set_active_index", this.$route.path);
    } else {
      this.$store.commit("add_tabs", {
        route: "/home",
        name: "home",
        title: "首页",
      });
      this.$store.commit("set_active_index", "/home");
      // this.$router.push("/home");
    }
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
    //tab标签点击时，切换相应的路由
    tabClick(tab) {
      if (tab.name == this.defaultactive) {
        return;
      } else {
        this.$router.push({ path: this.activeIndex });
      }
    },
    //移除tab标签
    tabRemove(targetName) {
      //首页不删
      if (targetName == "/home") {
        return;
      }
      this.$store.commit("delete_tabs", targetName);
      if (this.activeIndex === targetName) {
        // 设置当前激活的路由
        if (this.openTab && this.openTab.length >= 1) {
          this.$store.commit(
            "set_active_index",
            this.openTab[this.openTab.length - 1].route
          );
          this.$router.push({ path: this.activeIndex });
        } else {
          this.$router.push({ path: "/home" });
        }
      }
    },
  },
  computed: {
    openTab() {
      return this.$store.state.openTab;
    },
    activeIndex: {
      get() {
        return this.$store.state.activeIndex;
      },
      set(val) {
        this.$store.commit("set_active_index", val);
      },
    },
  },
  watch: {
    $route(to, from) {
      //判断路由是否已经打开
      //已经打开的 ，将其置为active
      //未打开的，将其放入队列里
      this.defaultactive = to.path;
      let flag = false;
      for (let item of this.openTab) {
        if (item.name === to.name) {
          this.$store.commit("set_active_index", to.path);
          flag = true;
          break;
        }
      }

      if (!flag) {
        this.$store.commit("add_tabs", {
          route: to.path,
          name: to.name,
          title: to.meta.title,
        });
        this.$store.commit("set_active_index", to.path);
      }
    },
  },
};
</script>
 