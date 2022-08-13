import Vue from "vue";
import VueRouter from "vue-router";
Vue.use(VueRouter);
import VueCookies from "vue-cookies";

const routes = [
  {
    path: "/",
    name: "index",
    component: () => import("@/pages/index/index.vue"),
    redirect: "/home",
    children: [
      {
        path: "/home",
        name: "home",
        meta: {
          title: "首页"
        },
        component: () => import("@/pages/home/home.vue")
      },
      {
        path: "/WxchatPublic",
        name: "WxchatPublic",
        meta: {
          title: "微信公众号"
        },
        component: () => import("@/pages/wxchat-public/index.vue")
      },
      {
        path: "/WxchatPublic/menu",
        name: "WxchatPublicMenu",
        meta: {
          title: "微信公众号菜单"
        },

        component: () => import("@/pages/wxchat-public/menu.vue")
      },
      {
        path: "/WxchatPublic/message",
        name: "WxchatPublicMessage",
        meta: {
          title: "微信公众号消息"
        },

        component: () => import("@/pages/wxchat-public/message.vue")
      },
      {
        path: "/admin/menu",
        name: "adminMenu",
        meta: {
          title: "面板菜单"
        },

        component: () => import("@/pages/admin/menu.vue")
      },
      {
        path: "/admin/user",
        name: "adminUser",
        meta: {
          title: "面板用户"
        },
        component: () => import("@/pages/admin/user.vue")
      }
    ]
  },
  {
    path: "/login",
    name: "login",
    meta: {
      title: "请登录"
    },
    component: () => import("@/pages/login/index.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});
// 全局后置钩子
router.afterEach(to => {
  // 设置title
  document.title = to.meta.title;
});
router.beforeEach((to, from, next) => {
  if (to.path === "/login") {
    next();
  } else {
    if (VueCookies.get("token")) {
      next();
    } else {
      next("/login");
    }
  }
});

export default router;
