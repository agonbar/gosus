    <template>
  <div class="content">
    <h4 id="raw">RAW DATA</h4>
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
      // c: {"data": { "value": {"chips":{"amdgpu-pci-0a00":{"Adapter":"PCI adapter","fan1":"823 RPM","power1":"50.15 W (cap = 183.00 W)","temp1":"+48.0°C (crit = +91.0°C, hyst = -273.1°C)","vddgfx":"1.20 V"},"asus-isa-0000":{"Adapter":"ISA adapter","cpu_fan":"0 RPM"},"it8665-isa-0290":{"+3.3V":"3.33 V","3VSB":"3.36 V (min = +4.73 V, max = +2.29 V)","Adapter":"ISA adapter","Vbat":"3.21 V","fan1":"557 RPM (min = 17 RPM)","fan2":"552 RPM (min = 12 RPM)","fan3":"564 RPM (min = 16 RPM)","fan5":"0 RPM (min = -1 RPM) ALARM","fan6":"0 RPM (min = -1 RPM)","in0":"1.33 V (min = +1.79 V, max = +2.42 V)","in1":"2.50 V (min = +1.32 V, max = +2.29 V)","in2":"2.01 V (min = +0.15 V, max = +1.64 V)","in3":"1.95 V (min = +2.64 V, max = +0.95 V)","in4":"33.00 mV (min = +1.05 V, max = +1.66 V)","in5":"33.00 mV (min = +1.95 V, max = +0.35 V)","in6":"33.00 mV (min = +0.77 V, max = +2.59 V)","intrusion0":"ALARM","temp1":"+48.0°C (low = +17.0°C, high = -65.0°C)","temp2":"+32.0°C (low = -126.0°C, high = +125.0°C) sensor = thermistor","temp3":"+38.0°C (low = -84.0°C, high = +109.0°C) sensor = thermistor","temp4":"+38.0°C (low = +46.0°C, high = -8.0°C) sensor = thermistor","temp5":"+38.0°C (low = -14.0°C, high = +27.0°C) sensor = thermistor","temp6":"+38.0°C (low = -45.0°C, high = -90.0°C) sensor = thermistor"},"k10temp-pci-00c3":{"Adapter":"PCI adapter","Tctl":"+68.1°C","Tdie":"+48.1°C (high = +70.0°C)"}}}}}
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