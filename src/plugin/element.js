// 按需引入 Element 模块

import Vue from 'vue';
import 'element-ui/lib/theme-chalk/base.css';
import { Button, Message, Form, FormItem, Input,Col,Row,Card} from 'element-ui';
Vue.use(Button);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.use(Col);
Vue.use(Row);
Vue.use(Card);

Vue.prototype.$message = Message;
