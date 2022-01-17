<template>
  <div
    id="app"
    class="image-div"
    :style="{ backgroundImage: 'url(' + dxx + ')' }"
  ></div>
</template>

<script>
import axios from 'axios';
export default {
  name: "App",
  data() {
    return {
      api_endpoint: 'https://service-0wsl5m13-1256946954.cd.apigw.tencentcs.com/release/qndxx',
      title: "正在获取title...",
      dxx: '正在获取img...'
    };
  },
  mounted() {
    axios
      .get(this.api_endpoint)
      .then((resp) => {
        this.title = resp.data.title
        this.dxx = resp.data.dxx_img
      })
      .catch((err) => {
        console.log(err)
      })
  },
  metaInfo() {
    return {
      title: this.title,
      htmlAttrs: {
        lang: "en",
        amp: true,
      },
    };
  },
};
</script>

<style>
.image-div {
  position: absolute;
  width: 100%;
  height: 100%;
  background-size: 100% 100%;
}

body {
  margin: 0;
  padding: 0;
}
</style>
