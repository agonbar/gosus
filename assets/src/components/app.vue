<template>
  <div id="app" class="container">
    <nav id="navbar-example2" class="navbar navbar-light bg-light">
      <a class="navbar-brand" href="#">Gosus</a>
      <router-link to="/">Main</router-link>
      <router-link to="/thermals">Thermals</router-link>
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