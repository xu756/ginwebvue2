// const webpack = require('webpack')


module.exports = {
  // configureWebpack: {
  //   plugins: [
  //     new webpack.ProvidePlugin({
  //       'window.Quill': 'quill/dist/quill.js',
  //       'Quill': 'quill/dist/quill.js'
  //     })
  //   ]
  // },
  transpileDependencies: true,
  // 跨域支持
  devServer: {
    proxy: {
      "/api": {
        target: "https://cs.xu756.top/api/",
        changeOrigin: true, // 允许跨域
        pathRewrite: {
          "^/api": ""
        }
      }
    }
  }
}
