<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="LevelId" label="关卡ID" width="180"></el-table-column>
      <el-table-column prop="PlayTimes" label="玩的次数" width="180"></el-table-column>
      <el-table-column prop="PassTimes" label="通关次数" width="180"></el-table-column>
      <el-table-column prop="PassRation" label="通关率" width="180"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="Delete(scope.row)" type="text" size="small">重置</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button @click="refresh()" size="big">刷新</el-button>
  </div>
</template>

<script>
import Utils from "@/api/utils";
let tableDataTemp = [];
function get_pass_ratio_callback(vueobj, response) {
  let json_obj = JSON.parse(response.data.data);
  console.log(json_obj);
  for (var key in json_obj) {
    let one_item = {};
    one_item.LevelId = json_obj[key]["LevelId"];
    one_item.PlayTimes = json_obj[key]["PlayTimes"];
    one_item.PassTimes = json_obj[key]["PassTimes"];
    one_item.PassRation = json_obj[key]["PassRation"];
    tableDataTemp.push(one_item);
    console.log("push", one_item.uid);
  }
  tableDataTemp.sort(function(a, b) {
    return parseInt(a.LevelId) - parseInt(b.LevelId);
  });
  vueobj.tableData = tableDataTemp;
}

function del_pass_ratio_callback(vueobj, response) {
  tableDataTemp = [];
  Utils.myajax("get_pass_ratio", "", get_pass_ratio_callback, this);
  vueobj.tableData = tableDataTemp;
}

export default {
  data() {
    return {
      tableData: tableDataTemp,
      dialogTableVisible: false,
      dialogFormVisible: false,
      form: {
        tag: "",
        whitelist: "",
        templist: ""
      },
      formLabelWidth: "120px"
    };
  },
  created() {
    Utils.myajax("get_pass_ratio", "", get_pass_ratio_callback, this);
  },
  methods: {
    Delete(val) {
      var LevelId = val["LevelId"];
      Utils.myajax("del_pass_ratio", LevelId, del_pass_ratio_callback, this);
    },
    refresh() {
      console.log("refresh");
      tableDataTemp = [];
      Utils.myajax("get_pass_ratio", "", get_pass_ratio_callback, this);
      this.tableData = tableDataTemp;
    }
  }
};
</script>