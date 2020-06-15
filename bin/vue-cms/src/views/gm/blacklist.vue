<template>
  <div>
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="uid"
      label="uid"
      width="180">
    </el-table-column>
    <el-table-column
      prop="account_name"
      label="账号名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="content"
      label="账号信息">
    </el-table-column>
    <el-table-column
      label="操作"
      >
      <template slot-scope="scope">
        <el-button @click="handleClick(scope.row)" type="text" size="small">解封</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-button @click="refresh()" size="big">刷新</el-button>
  </div>

</template>

<script>
  import Utils from '@/api/utils'
  let tableDataTemp = []
  function get_blacklist_callback(vueobj, response) {        
          let json_obj = JSON.parse(response.data.data)
          console.log(json_obj)
          for(var key in json_obj){
              let one_item = {}
              one_item.uid = key
              one_item.account_name = JSON.parse(json_obj[key])["AccountName"]
              one_item.content = json_obj[key]
              tableDataTemp.push(one_item)
              console.log("push",one_item.uid)
          }
          vueobj.tableData = tableDataTemp
          
    }
    function free_callback(response, vueobj) {
        console.log("response==", response);
        alert("解封成功");
        vueobj.tableData = []
        tableDataTemp = []
        Utils.myajax("get_black_list", "", get_blacklist_callback, vueobj)
       
    }

   export default {
    data() {
      return {
        tableData: tableDataTemp
      }
    },
    created(){
      Utils.myajax("get_black_list", "", get_blacklist_callback, this)
    },
    methods: {
      handleClick(val) {
      var send = "[" + val.uid + "]";
      console.log("send=", send);
      if (send != null) {
        Utils.myajax("free_user", send,free_callback, this);
      }
      this.tableData = tableDataTemp
      },
      refresh(){
      console.log("refresh")
      tableDataTemp = []
      Utils.myajax("get_black_list", "", get_blacklist_callback, this)
      this.tableData = tableDataTemp
    }
    },
    }
</script>