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
        <el-form-item label="验证码">
          <el-input v-model="userform.captcha"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="5">
        <el-form-item label="验证码">
          <s-identify identifyCode="1254"></s-identify>
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
      },
    };
  },
  methods: {
    // 提交表单
    submit() {
      this.$refs.form_ref.validate((valid) => {
        if (valid) {
          console.log(this.userform);
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
  },
};
</script>
<style lang="scss">
</style>;
