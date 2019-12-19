<template>
  <div id="app" class="container">
      <!-- bootstrap -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
    integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    
    <nav id="navbar-example2" class="navbar navbar-light bg-light">
      <div class="navbar-brand">Gosus</div>
      <router-link to="/thermals">Thermals</router-link>
      <router-link to="/fans">Fans</router-link>
    </nav>
    <br />
    <br />
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  data() {
    return {
    };
  },
  async mounted() {
    await this.initialLoad();
    console.log("created app.vue");
  },
  methods: {
    async initialLoad() {
    },
  }
};
</script>