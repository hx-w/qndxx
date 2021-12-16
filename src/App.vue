<template>
  <div id="app" class="image-div" :style="{backgroundImage: 'url(' + dxx + ')'}">
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      dxx: "https://h5.cyol.com/special/daxuexi/bsz1fk1u3w/images/end.jpg",
    };
  },
};
</script>

<style>
.image-div {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: 100% 100%;
}
</style>
