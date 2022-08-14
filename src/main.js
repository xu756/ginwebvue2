import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "@/assets/css/global.min.css";
import "default-passive-events";
// import "@/assets/icon/iconfont.css";
import "./plugin/element.js";
Vue.config.productionTip = false;
//axios
import { get, post } from "./plugin/axios.js";
Vue.prototype.$get = get;
Vue.prototype.$post = post; // 将axios挂载到Vue的原型上，这样所有的组件都可以使用this.$http访问到axios
// js-md5 加密
import md5 from "js-md5";
Vue.prototype.$md5 = md5;
// moment 时间格式化
import moment from "moment";
Vue.prototype.$moment = moment;
import VueCookies from "vue-cookies";
Vue.use(VueCookies);
import CKEditor from "@ckeditor/ckeditor5-vue2";
Vue.use(CKEditor);
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
