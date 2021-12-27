<template>
  <div
    id="app"
    class="image-div"
    :style="{ backgroundImage: 'url(' + dxx + ')' }"
  ></div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      title: "“青年大学习”第十二季第十三期",
      dxx: 'https://h5.cyol.com/special/daxuexi/bf1avt2y7q/images/end.jpg'
    };
  },
  mounted() {
    // var urlPattern = /<li>.*<a href="(https.*?\.html)"/g
    this.$jsonp('http://news.cyol.com/gb/channels/vrGlAKDl/index.html')
      .then((resp) => {
        console.log(resp.data)
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
