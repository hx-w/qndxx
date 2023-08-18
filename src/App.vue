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
      api_endpoint: window.location.href + '/.netlify/functions/qndxx',
      title: "正在请求数据...",
      result: '正在获取图片...'
    };
  },
  mounted() {
    axios
      .get(this.api_endpoint)
      .then((resp) => {
        this.result = resp.data.result
        window.location.replace(this.result)
      })
      .catch((err) => {
        console.log(err)
        this.result = resp.data.result
        window.location.replace(this.result)
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
