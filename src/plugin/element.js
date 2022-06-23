// 按需引入 Element 模块

import Vue from 'vue';
import { Button, Message, Form, FormItem, Input,} from 'element-ui';
Vue.use(Button);
Vue.use(Form);
Vue.use(FormItem);

Vue.prototype.$message = Message;
