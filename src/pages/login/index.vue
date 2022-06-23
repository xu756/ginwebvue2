<template>
  <div id="Maxbox">
    <div id="bg"></div>
    <div id="Mainbox">
      <div class="login_img">
        <img src="../../assets/images/bg1.png" alt="" />
      </div>
      <div class="login_box"></div>
    </div>
  </div>
</template>
<script>
export default {
  name: "Login",
  data() {
    return {
     
    };
  },
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
    width: 1220px;
    height: 550px;
    padding: 122px;
    opacity: 0.7;
    background: #ffffff;
    border-radius: 20px;
    box-shadow: 0px 22px 49px 0px rgba(16, 66, 165, 0.6);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    .login_img {
      width: 40%;
      height: 100%;
      float: left;
      img {
        width: 80%;
        height: 90%;
        margin: 0 10%;
      }
    }
    .login_box {
      width: 40%;
      height: 100%;
      float: right;
      background: #400404;
    }
  }
}
</style>