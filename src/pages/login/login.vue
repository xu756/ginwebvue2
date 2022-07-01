<template>
  <el-form ref="form_ref" :model="userform" :rules="form_rules">
    <el-form-item label="用户名" prop="username">
      <el-input v-model="userform.username"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input type="password" v-model="userform.password"></el-input>
    </el-form-item>
    <el-row :gutter="20" align="bottom">
      <el-col :span="15">
        <el-form-item label="验证码" prop="captcha">
          <el-input v-model="userform.captcha"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="5">
        <el-form-item label="验证码">
          <s-identify
            :identifyCode="code"
            @click.native="resetcode"
          ></s-identify>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-button type="primary" @click="submit" style="width: 100%"
          >提交</el-button
        >
      </el-col>
      <el-col :span="12">
        <el-button type="primary" @click="reset" style="width: 100%"
          >重置</el-button
        >
      </el-col>
    </el-row>
  </el-form>
</template>
<script>
import SIdentify from "@/components/SIdentify.vue";
export default {
  name: "Login",
  components: {
    SIdentify,
  },
  data() {
    return {
      // 表单数据
      userform: {
        username: "", // 用户名
        password: "", // 密码
        captcha: "", // 验证码
      },
      code: "",
      //验证规则
      form_rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 3, max: 5, message: "长度在3到5个字符", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 3, max: 5, message: "长度在3到5个字符", trigger: "blur" },
        ],
        captcha: [
          { required: true, message: "请输入验证码", trigger: "blur" },
          { min: 4, max: 4, message: "长度为4个字符", trigger: "blur" },
        ],
      },
    };
  },
  created() {
    this.resetcode();
  },
  methods: {
    // 提交表单
    submit() {
      this.$refs.form_ref.validate((valid) => {
        if (valid) {
          if (this.userform.captcha == this.code) {
            this.$store.dispatch("login", this.userform).then((res) => {
              if (res.code == 200) {
                this.$message({
                  message: "登录成功",
                  type: "success",
                });
                this.$router.push("/");
              } else {
                this.$message({
                  message: res.msg,
                  type: "error",
                });
              }
            });
          } else {
            this.$message({
              message: "验证码错误",
              type: "error",
            });
            // 重置验证码
            this.resetcode();
            //清空密码
            this.userform.password = "";
          }
        } else {
          this.$message({
            message: "请检查表单数据",
            type: "error",
          });
          return false;
        }
      });
    },
    //重置表单
    reset() {
      this.$refs.form_ref.resetFields();
    },
    // 重置验证码
    resetcode() {
      this.code = Math.random().toString().slice(2, 6); //生成验证码
      this.userform.captcha = ""; //清空验证码
    },
  },
};
</script>
<style lang="scss">
</style>;
