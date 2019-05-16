<template>
  <div class="w-full">
    <div class="h-2 bg-red-500"></div>
    <div class="flex items-center justify-center h-screen">
      <div class="container mx-24 bg-white rounded shadow-lg">
        <div class="px-12 py-6">
          <div class="flex justify-center">
            <img src="https://www.wewee.com/_nuxt/img/logo.a473a35.png"
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
                         class="flex w-full appearance-none rounded shadow p-3 text-grey-dark mr-2 focus:outline-none">
                  <button @click.prevent="genHandler"
                          type="submit"
                          class="appearance-none bg-red-400 text-white text-base font-semibold tracking-wide uppercase p-3 rounded shadow hover:bg-indigo-light">Generate</button>
                </div>
              </form>
              <div class="mt-3" v-if=shortUrl>
                <h2>{{shortUrl}}</h2>
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
      shortUrl:null
    };
  },

  methods: {
    genHandler() {
      if (!this.isUrl(this.originUrl)) {
        alert("请输入正确url地址");
        return
      }
      axios.post("/api/short-urls",{url:this.originUrl}).then(({data}) => {
        this.shortUrl = data.data.url
        
      }).catch((err) => {
        console.error(err)
      });
    },

    isUrl(url) {
      return this.urlReg.test(url);
    }
  }
};
</script>