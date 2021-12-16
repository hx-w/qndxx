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
      title: "“青年大学习”第十二季第十一期",
      dxx: "https://h5.cyol.com/special/daxuexi/bsz1fk1u3w/images/end.jpg",
    };
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
