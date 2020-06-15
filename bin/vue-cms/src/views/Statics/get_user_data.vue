<template>
  <div class="block">
    <div class="demo-input-suffix" style="left: 50%">
      UID：
      <el-input v-model="uid" placeholder="请输入玩家UID"></el-input>
      <el-button type="text" @click="Send()">确定</el-button>
    </div>
    <el-divider></el-divider>
    <div class="demo-input-suffix">
      玩家数据：
      <el-input type="textarea" :rows="1024" placeholder="玩家数据：" v-model="user_data"></el-input>
    </div>
  </div>
</template>

<script>
import Utils from "@/api/utils";
let tableDataTemp = [];
function get_user_data_callback(vueobj, response) {
  let json_obj = JSON.parse(response.data.data);
  console.log(json_obj);
  vueobj.user_data = response.data.data;
}

export default {
  data() {
    return {
      uid: "",
      user_data: ""
    };
  },
  methods: {
    Send() {
      Utils.myajax("get_user_data", this.uid, get_user_data_callback, this);
    }
  }
};
</script>