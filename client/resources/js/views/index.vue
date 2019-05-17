<template>
  <div class="w-full">
    <div class="h-2 bg-red-500"></div>
    <div class="flex items-center justify-center h-screen">
      <div class="container mx-24 bg-white rounded ">
        <div class="px-12 py-6">
          <div class="flex justify-center mb-3">
            <img class="object-cover h-20 " src="/img/logo.png"
                 alt="">
          </div>
          <div class="text-center">
            <h1 class="font-bold text-3xl text-gray-500 my-3 w-full">Short Url Service</h1>
            <div class="w-full text-center">
              <form action="#">
                <div class="mx-5 mx-auto p-1 pr-0 flex items-center">
                  <input 
                         v-model="originUrl"
                         placeholder="https://www.wewee.com/"
                         class="flex w-full bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-red-400 mr-2">
                  <button @click.prevent="genHandler"
                          type="button"
                          class="shadow bg-red-400 hover:bg-red-300 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded">Generate</button>
                </div>
              </form>
              <QRCode :value=shortUrl class="mt-3" v-show=shortUrl></QRCode>
              <div class="mt-3" v-if=shortUrl>
                <h2 class="text-gray-700">{{shortUrl}}</h2>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "index-page",

  data() {
    return {
      originUrl: null,
      urlReg: /^((ht|f)tps?):\/\/([\w\-]+(\.[\w\-]+)*\/)*[\w\-]+(\.[\w\-]+)*\/?(\?([\w\-\.,@?^=%&:\/~\+#]*)+)?/,
      shortUrl:null,
    }
  },

  methods: {
    genHandler() {
      if (!this.isUrl(this.originUrl)) {
        alert("请输入正确url地址");
        return
      }
      this.$root.Bus.$emit('loading')
      axios.post("/api/short-urls",{url:this.originUrl}).then(({data}) => {
        this.shortUrl = data.data.url
        
      }).catch((err) => {
        console.error(err)
      });
      this.$root.Bus.$emit('loaded')
    },

    isUrl(url) {
      return this.urlReg.test(url);
    }
  }
};
</script>