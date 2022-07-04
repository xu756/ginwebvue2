import Vue from 'vue';
import VueRouter from 'vue-router';
Vue.use(VueRouter);
import VueCookies from 'vue-cookies'

const routes = [
	{
		path: '/',
		name: 'index',
		component: () => import('@/pages/index/index.vue')
	},
	{
		path: '/login',
		name: 'login',
		component: () => import('@/pages/login/index.vue')
	}
];

const router = new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes
});
router.beforeEach((to, from, next) => {
	
	if (to.path === '/login') {
	  next()
	} else {
	  if (VueCookies.get('token')) {
		next()
	  } else {
		next('/login')
	  }
	}
  })

export default router;
