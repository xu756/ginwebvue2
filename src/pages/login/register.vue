<template>
  <el-form
    ref="user_ref"
    :model="user"
    :rules="user_rules"
    label-width="65px"
    label-position="left"
  >
    <el-form-item label="用户名" prop="username">
      <el-input v-model="user.username"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="user.password"></el-input>
    </el-form-item>
    <el-row>
      <el-col :span="16">
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="user.email"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="6" :offset="1">
        <el-button type="primary" @click="setcode" v-if="sendemail == 'false'"
          >发送验证码</el-button
        >
        <div v-else style="line-height: 40px">
          {{ emailtime + " s后重新发" }}
        </div>
      </el-col>
    </el-row>
  </el-form>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {
      user: {
        username: "admin",
        password: "123456",
        email: "756334744@qq.com",
      },
      user_rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 3, max: 5, message: "长度在3到5个字符", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 12, message: "长度在3到5个字符", trigger: "blur" },
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
          { required: true, message: "请输入验证码", trigger: "blur" },
          { min: 4, max: 4, message: "长度在3到5个字符", trigger: "blur" },
        ],
      },
      sendemail: "false",
      emailtime: 5,
    };
  },
  methods: {
    setcode() {
      this.$refs.user_ref.validate((valid) => {
        if (valid) {
          this.sendemail = "true";
          var T = setInterval((ID) => {
            this.emailtime--;
            if (this.emailtime == 0) {
              //清除定时器
              clearInterval(T);
              this.sendemail = "false";
              this.emailtime = 5;
            }
          }, 1000);
          this.$post("register1", {
            username: this.user.username,
            password: this.user.password,
            email: this.user.email,
          }).then((res) => {
            console.log(res);
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
  },
};
</script>
<style lang="scss">
</style>