import axios from 'axios'
export default class UploadAdapter {
	constructor(loader) {
	  this.loader = loader
	}
	upload() {
	  return this.loader.file.then(file => new Promise((resolve, reject) => {
		const data = new FormData()
		data.append('img', file)
		axios.request({
		  url: `/upload/img`,// 上传文件的接口地址，实际请填写完整地址
		  method: 'post',
		  data,
		  headers: {
			'Content-Type': 'multipart/form-data'
		  }
		}).then(res => {
		  if (res.data.type === 'secess') {
			const url = res.data.url // 后台返回的上传成功后的图片地址
			resolve({
			  default: url
			})
		  }
		}).catch(error => {
		  reject(error)
		})
	  }))
	}
	abort() {
	}
  }
  