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
        title: "首页",
        component: () => import("@/pages/home/home.vue")
      },
      {
        path: "/WxchatPublic",
        name: "WxchatPublic",
        title:"微信公众号",
        component: () => import("@/pages/wxchat-public/index.vue")
      },
      {
        path: "/WxchatPublic/menu",
        name: "WxchatPublicMenu",
        title:"微信公众号菜单",
        component: () => import("@/pages/wxchat-public/menu.vue")
      },
      {
        path: "/WxchatPublic/message",
        name: "WxchatPublicMessage",
        title:"微信公众号消息",
        component: () => import("@/pages/wxchat-public/message.vue")
      },
      {
        path: "/admin/menu",
        name: "adminMenu",
        title:"面板菜单",
        component: () => import("@/pages/admin/menu.vue")
      },
      {
        path: "/admin/user",
        name: "adminUser",
        title:"面板用户",
        component: () => import("@/pages/admin/user.vue")
      }
    ]
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/pages/login/index.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
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
