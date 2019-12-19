    <template>
  <div class="content">
    <h4 id="raw">PARSED DATA</h4>
    <div>
      <div class="counter">{{c.data.value}}</div>
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