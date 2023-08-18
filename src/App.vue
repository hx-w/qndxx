<template>
  <div
    id="app"
    class="image-div"
    :style="{ backgroundImage: 'url(' + result + ')' }"
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
        this.title = "暂时无法获取标题"
        this.result = resp.data.result
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
