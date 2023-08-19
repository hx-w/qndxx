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
      title: '请使用[?id=20xx年第xx期]的形式访问',
      result: 'https://www.example.com'
    };
  },
  mounted() {
    axios
      .get(this.api_endpoint)
      .then((resp) => {
        console.log(this.$route.params)
        // show all route param
        console.log(this.$route)
        if (this.$route.query.params['id']) {
          this.title = '“青年大学习”' + this.$route.query.params['id']
        }
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
