<script setup lang="ts">
document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <!-- Page -->
  <div class="view">
    <router-view />
  </div>
</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  background-color: #282a36;
  overflow: hidden;
}

.view {
  position: absolute;
  top: 0px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}
</style>
