<template>
  <div>
    <div class="block">
      <span class="demonstration">有效开始时间</span>
      <el-date-picker
        v-model="value2"
        type="datetime"
        placeholder="选择日期时间"
        align="right"
        :picker-options="pickerOptions"
      ></el-date-picker>
    </div>
    <div class="block">
      <span class="demonstration">有效结束</span>
      <el-date-picker
        v-model="value2"
        type="datetime"
        placeholder="选择日期时间"
        align="right"
        :picker-options="pickerOptions"
      ></el-date-picker>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: tableDataTemp
    };
  },
  created() {},
  methods: {}
};
</script>