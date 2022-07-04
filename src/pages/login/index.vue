<template>
  <div id="Maxbox">
    <div id="bg"></div>
    <div id="Mainbox">
      <div class="form_img">
        <img src="../../assets/images/bg1.png" />
      </div>
      <div class="form_box">
        <!-- 切换 -->
        <div class="form_tab">
          <div
            class="tab_item"
            @click="tab('login')"
            :class="{ active: this.show == 'login' }"
          >
            登录
          </div>
          <div
            class="tab_item"
            @click="tab('register')"
            :class="{ active: this.show == 'register' }"
          >
            注册
          </div>
        </div>
        <div class="form_container">
          <login-box
            key="1"
            v-if="show == 'login'"
            class="animate__animated animate__fast animate__fadeInLeft"
          ></login-box>
          <register-box
            key="2"
            v-if="show == 'register'"
            class="animate__animated animate__flipInX"
          ></register-box>
        </div>
      </div>
    </div>
  </div>
</template> 
<script>
import LoginBox from "./login.vue";
import RegisterBox from "./register.vue";
export default {
  name: "Login",
  data() {
    return {
      show: "login",
    };
  },
  // 登录页面的组件
  components: {
    LoginBox: LoginBox, // 登录页面的组件
    RegisterBox: RegisterBox, // 注册页面的组件
  },
  methods: {
    tab(type) {
      this.show = type;
    },
  }
};
</script>
<style lang="scss">
#Maxbox {
  width: 100%;
  height: 100%;
  text-align: center;
  color: #000;
  // 背景
  #bg {
    width: 100%;
    height: 100%;
    position: absolute;
    z-index: -10;
    background: url("../../assets/images/bg0.png") no-repeat center center /
      cover;
  }
  #Mainbox {
    width: 70%;
    height: 60%;
    padding: 3%;
    opacity: 0.7;
    background: #ffffff;
    border-radius: 20px;
    box-shadow: 0px 22px 49px 0px rgba(16, 66, 165, 0.6);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    // backdrop-filter: blur(50px);
    .form_img {
      width: 40%;
      height: 100%;
      float: left;
      img {
        width: 80%;
        height: 90%;
        margin: 0 10%;
      }
    }
    .form_box {
      width: 32%;
      height: 80%;
      float: right;
      background-image: linear-gradient(62deg, #8ec5fc 0%, #e0c3fc 100%);
      padding: 2%;
      border-radius: 20px;
      .form_tab {
        width: 100%;
        height: 50px;
        .active {
          color: #00a0e9 !important;
          border-bottom: 1px solid #000 !important;
        }
        .tab_item {
          width: 50%;
          height: 100%;
          border-bottom: 1px solid #e5e5e5;
          float: left;
          text-align: center;
          line-height: 50px;
          color: #000;
          font-size: 20px;
          font-weight: bold;
          cursor: pointer;
        }
      }
      .form_container {
        margin-top: 20px;
        width: 90%;
        margin-left: 5%;
      }
    }
  }
}
</style>