const { defineConfig } = require('@vue/cli-service');
module.exports = defineConfig({
	transpileDependencies: true,
	// 跨域支持
	devServer: {
		proxy: {
			'/api/vue2/user': {
				target: 'http://localhost:7000/api/vue2/user/',
				changeOrigin: true, // 允许跨域
				pathRewrite: {
					'^/api/vue2/user': ''
				}
			}
		}
	}
});
