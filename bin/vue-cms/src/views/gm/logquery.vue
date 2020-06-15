<template>
  <div>
    <div class="block">
      <span class="demonstration">起止时间</span>
      <el-date-picker
        v-model="value1"
        type="datetimerange"
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
      ></el-date-picker>
    </div>
    <div class="demo-input-suffix" style="position: relative;width: 200px;">
      玩家UID：
      <el-input v-model="uid"></el-input>
      <button
        type="button"
        class="el-button el-button--primary is-plain"
        style="position: relative;top: -40px;left:210px"
      >
        <span v-on:click="query_ok_button">确定</span>
      </button>
    </div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="uid" label="uid" width="50"></el-table-column>
      <el-table-column prop="time" label="时间" width="50"></el-table-column>
      <el-table-column prop="content" label="内容"></el-table-column>
    </el-table>
  </div>
</template>

<script>
import axios from "axios";
import Utils from "@/api/utils";
function query_callback(vueobj, response) {
  let json_obj = JSON.parse(response.data.data);
  console.log(json_obj);
  for (var key in json_obj) {
    let one_item = {};
    one_item.uid = json_obj[key]["Uid"];
    one_item.time = json_obj[key]["Time"];
    one_item.content = json_obj[key]["Content"];
    tableDataTemp.push(one_item);
    console.log("push", one_item.uid);
  }
  vueobj.tableData = tableDataTemp;
}
export default {
  data() {
    return {
      pickerOptions: {},
      value1: [new Date(), new Date()],
      uid: ""
    };
  },
  methods: {
    query_ok_button() {
      console.log(this.value1[0]);
      var startTime = Math.ceil(this.value1[0].getTime() / 1000);
      var EndTime = Math.ceil(this.value1[1].getTime() / 1000);
      var send = this.uid + "|" + startTime + "|" + EndTime;
      console.log("send==", send);
      if (send != null) {
        Utils.myajax("querylog", send, query_callback, this);
      }
    },
    inputTargetValue: function(e) {
      console.log(e.target.value);
      this.uid = e.target.value;
    }
  }
};
</script>