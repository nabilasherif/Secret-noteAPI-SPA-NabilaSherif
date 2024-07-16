<template>
  <img alt="Note logo" src="./assets/notelogo.png">
  <HelloWorld msg="Welcome to Secret Note"/>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  }
}
</script>

<style>
img {
  width: 100px;
}
#app {
  
  font-family: "Lucida Console", "Courier New", monospace;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
