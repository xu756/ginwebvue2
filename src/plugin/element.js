// 按需引入 Element 模块

import Vue from 'vue';
import { Button, Message} from 'element-ui';
Vue.use(Button);
Vue.prototype.$message = Message;
