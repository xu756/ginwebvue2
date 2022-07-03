<template>
  <el-form
    ref="user_ref"
    :model="user"
    :rules="user_rules"
    label-width="65px"
    label-position="left"
  >
    <el-form-item label="用户名" prop="username">
      <el-input v-model="user.username" :disabled="from"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="user.password" :disabled="from" type="password"></el-input>
    </el-form-item>
    <el-row>
      <el-col :span="16">
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="user.email"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" :offset="1">
        <el-button type="primary" @click="setcode" v-if="sendemail"
          >发送验证码</el-button
        >
        <div style="line-height: 40px" v-else>
          {{ emailtime + " s后重新发" }}
        </div>
      </el-col>
    </el-row>
    <el-form-item label="验证码" prop="code">
      <el-input v-model="user.code"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click.native="register" style="width: 100%"
        >注册</el-button
      >
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {
      user: {
        username: "",
        password: "",
        email: "",
        code: "",
        randStr: "",
      },
      user_rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 3, max: 5, message: "长度在3到5个字符", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 12, message: "长度在6到12个字符", trigger: "blur" },
        ],
        email: [
          { required: true, message: "请输入邮箱", trigger: "blur" },
          {
            validator: (rule, value, callback) => {
              if (value === "") {
                callback();
              } else {
                const reg =
                  /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/;
                if (reg.test(value)) {
                  callback();
                } else {
                  callback(new Error("请输入正确的邮箱"));
                }
              }
            },
          },
        ],
        code: [
          {
            validator: (rule, value, callback) => {
              //判断 user.randStr 是否为空
              if (this.user.randStr === "") {
                callback();
              } else {
                // 必填项
                if (value === "") {
                  callback(new Error("请输入验证码"));
                }
              }
            },
          },
        ],
      },
      from: false,
      sendemail: true,
      emailtime: 5,
    };
  },
  methods: {
    //1.发送邮箱验证码
    setcode() {
      this.$refs.user_ref.validate((valid) => {
        if (valid) {
          this.sendemail = false;
          var T = setInterval(() => {
            this.emailtime--;
            if (this.emailtime == 0) {
              //清除定时器
              clearInterval(T);
              this.sendemail = true;
              this.emailtime = 5;
            }
          }, 1000);
          this.$post("register1", this.user).then((res) => {
            if (res.type == "error") {
              this.$message.error(res.msg);
              this.$refs.user_ref.resetFields();
              this.from = false;
            } else {
              this.from = true;
              this.user.randStr = res.randStr;
            }
          });
        } else {
          this.$message({
            message: "请检查表单数据",
            type: "error",
          });
          return false;
        }
      });
    },
    //2.注册
    register() {
      if (this.from) {
        this.$post("register2", this.user).then((res) => {
          console.log(res);
        });
      } else {
        this.$message({
          message: "请先发送验证码",
          type: "error",
        });
      }
    },
  },
};
</script>