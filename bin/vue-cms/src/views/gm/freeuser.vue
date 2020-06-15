<template>
  <div>
    <span>解封-请输入玩家uid</span>
    <div class="el-input">
      <input
        type="text"
        autocomplete="off"
        placeholder="请输入UID"
        id="free_uid"
        class="el-input__inner"
        style="width:300px"
        v-on:keyup="inputTargetValue"
      />
      <button
        type="button"
        class="el-button el-button--primary is-plain"
        style="position: relative;left: 0px;"
      >
        <span v-on:click="free_ok_button">确定</span>
      </button>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import Utils from '@/api/utils'
function free_callback(response) {
    console.log("response==", response);
    alert("解封成功");
}

export default {
  data() {
    return {
      uid: ""
    };
  },
  methods: {
    free_ok_button() {
      var send = "[" + this.uid + "]";
      console.log("send=", send);
      if (send != null) {
        Utils.myajax("free_user", send, free_callback,this);
      }
    },
    inputTargetValue: function(e) {
      console.log(e.target.value);
      this.uid = e.target.value;
    }
  }
};
</script>