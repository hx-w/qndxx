<template>
  <div
    id="app"
    class="image-div"
    :style="{ backgroundImage: 'url(' + result + ')' }"
  >
    <div v-if="show">
      <p>请输入当前期号，形如"2023年第21期"</p>
      <input v-model="user_input">
      <button v-on:click="submit">确认</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "App",
  data() {
    const url = window.location.href
    return {
      api_endpoint: url.split('?')[0] + '/.netlify/functions/qndxx',
      title: '请在弹窗中输入标题',
      result: 'https://www.example.com',
      show: true
    };
  },
  mounted() {
    axios
      .get(this.api_endpoint)
      .then((resp) => {
        if (this.$route.query.id) {
          this.title = '“青年大学习”' + this.$route.query.id
          this.show = false
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
  methods: {
    submit: function () {
      this.title = '"青年大学习"' + this.user_input
      this.show = false
    }
  }
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
