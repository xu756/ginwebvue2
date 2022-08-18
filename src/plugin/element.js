// 按需引入 Element 模块

import Vue from "vue";
import "element-ui/lib/theme-chalk/base.css";
import {
  Button,
  Message,
  Form,
  FormItem,
  Input,
  Col,
  Row,
  Container,
  Header,
  Aside,
  Main,
  Menu,
  MenuItem,
  MenuItemGroup,
  Submenu,
  Table,
  TableColumn,
  Upload,
  Tabs,
  TabPane,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Tooltip,
  Avatar,
  Dialog,
  MessageBox,
  Pagination,
  Tag,
  Select,
  Option,
} from "element-ui";
Vue.use(Button);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.use(Col);
Vue.use(Row);
Vue.use(Container);
Vue.use(Header);
Vue.use(Aside);
Vue.use(Main);
Vue.use(Menu);
Vue.use(MenuItem);
Vue.use(MenuItemGroup);
Vue.use(Submenu);
Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Upload);
Vue.use(Tabs);
Vue.use(TabPane);
Vue.use(Dropdown);
Vue.use(DropdownMenu);
Vue.use(DropdownItem);
Vue.use(Tooltip);
Vue.use(Avatar);
Vue.use(Dialog);
Vue.use(Pagination);
Vue.use(Tag);
Vue.use(Select);
Vue.use(Option);
Vue.prototype.$message = Message;
Vue.prototype.$confirm = MessageBox.confirm;
