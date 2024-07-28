<template>
  <v-app>
    <NavBar />
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import NavBar from './components/NavBar.vue';
  
  export default defineComponent({
    name: 'App',
    components: {
      NavBar
    },
  });
  </script>