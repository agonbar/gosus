    <template>
  <div class="content">
    <h4 id="raw">PARSED DATA</h4>
    <div>
      <div class="row" v-for="(sensor) in c.data.value" :key="sensor.key">
        <div class="list-group" v-for="(measure, measureIndex) in sensor" :key="measure.key">
          <ul id="measures" class="list-group" v-if="measureIndex!='asus-isa-0000' && measure.Adapter">
            {{ measureIndex }}
            <li class="list-group-item" v-for="(item) in measure" :key="item.key">
              {{ item }}
            </li>
          </ul>
        </div>
      </div>
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
    async initialLoad() {
      this.interval = setInterval(() => counter.update(), 1000);
    },
    update: function() {
      counter.update();
    }
  }
};
</script>

    <style>
.list-group.inner li {
  border: none;
}
</style>