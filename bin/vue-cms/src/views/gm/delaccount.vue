<template>
  <div>
    <span>删除账号-请输入玩家uid</span>
    <div class="el-input">
      <input
        type="text"
        autocomplete="off"
        placeholder="请输入UID"
        id="ban_uid"
        class="el-input__inner"
        style="width:300px"
        v-on:keyup="inputTargetValue"
      />
      <button
        type="button"
        class="el-button el-button--primary is-plain"
        style="position: relative;left: 0px;"
      >
        <span v-on:click="ban_ok_button">确定</span>
      </button>
    </div>
  </div>
</template>
<script>
import axios from "axios";
function del_callback(response, vueobj) {
  console.log("response==", response);
  alert("删除账号成功");
}

export default {
  data() {
    return {
      uid: ""
    };
  },
  methods: {
    ban_ok_button() {
      var send = this.uid;
      console.log("send=", send);
      if (send != null) {
        this.myajax("del_user", send, del_callback, this);
      }
    },
    inputTargetValue: function(e) {
      console.log(e.target.value);
      this.uid = e.target.value;
    }
  }
};
</script>