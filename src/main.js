import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './plugin/element.js';
Vue.config.productionTip = false;
//axios
import request from './plugin/axios.js';
Vue.prototype.$http = request;	// 将axios挂载到Vue的原型上，这样所有的组件都可以使用this.$http访问到axios
// js-md5
import md5 from 'js-md5';
Vue.prototype.$md5 = md5;
new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount('#app');
