<template>
  <div id="app">
    <div id="nav">
        <img src="disk.svg" style="margin-top: -10px; height: 45px; "/>
        <router-link to="/">    MTest-Disk</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 0px;
}
#nav {
  padding: 10px;
}
#nav a {
  font-weight: bold;
  color: #465546;
  font-size: large;
}
</style>
