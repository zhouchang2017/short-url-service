<template>
  <div>
    <div id="qrcode"></div>
  </div>
</template>

<script>
import QRCode from "qrcodejs2";
export default {
  name: "qr-code",

  props:{
      value:{
          type:String
      }
  },

  data(){
      return {
          url:null,
          qrcode:null
      }
  },

  watch:{
      value(value,oldValue){
          this.qrcode.clear()
          this.qrcode.makeCode(value)
      }
  },

  methods: {
    init() {
      this.qrcode = new QRCode("qrcode", {
        text: this.url,
        width: 128,
        height: 128,
        colorDark: "#4A5568",
        colorLight: "#ffffff",
        correctLevel: QRCode.CorrectLevel.H
      });
    }
  },

  mounted() {
    this.init();
  }
};
</script>

<style lang="less">
#qrcode {
  display: inline-block;
  img {
    background-color: #fff;
    padding: 6px;
  }
}
</style>
