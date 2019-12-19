    <template>
  <div class="content">
    <h4 id="raw">RAW DATA</h4>
    <div>
      <div class="counter">{{c.data.value}}</div>
      <button v-on:click="incr">Incr</button>
    </div>
  </div>
</template>

    <script>
export default {
  data() {
    return {
      c: counter,
    };
  },
  async mounted() {
    await this.initialLoad();
  },
  methods: {
    async initialLoad() {},
    incr: function() {
      counter.add("mod");
    }
  }
};
</script>

    <style>
.list-group.inner li {
  border: none;
}
</style>