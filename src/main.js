import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './plugin/element.js';
Vue.config.productionTip = false;
//axios
import request from './plugin/axios.js';
Vue.prototype.$http = request;
new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount('#app');
