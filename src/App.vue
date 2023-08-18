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
      api_endpoint: 'https://9966df427a59473692f15d799442f2ce.apig.ap-southeast-3.huaweicloudapis.com/qndxx',
      title: "正在请求数据...",
      dxx: '正在获取图片...'
    };
  },
  mounted() {
    window.location.replace('https://baidu.com')
    // axios
    //   .get(this.api_endpoint)
    //   .then((resp) => {
    //     this.title = resp.data.title
    //     this.dxx = resp.data.dxx_img
    //   })
    //   .catch((err) => {
    //     console.log(err)
    //   })
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
