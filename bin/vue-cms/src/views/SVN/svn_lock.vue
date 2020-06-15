<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="tag" label="Tag" width="180"></el-table-column>
      <el-table-column prop="whitelist" label="白名单" width="180"></el-table-column>
      <el-table-column prop="templist" label="临时允许提交名单" width="180"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="DialogOpen(scope.row)" type="text" size="small">修改</el-button>
          <el-button @click="Delete(scope.row)" type="text" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button @click="refresh()" size="big">刷新</el-button>
    <el-button @click="Add()" size="big">添加</el-button>

    <el-dialog title="tag编辑" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="tag" :label-width="formLabelWidth">
          <el-input v-model="form.tag" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="白名单" :label-width="formLabelWidth">
          <el-input v-model="form.whitelist" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="临时权限" :label-width="formLabelWidth">
          <el-input v-model="form.templist" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="ModifyTag()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Utils from "@/api/utils";
let tableDataTemp = [];
function get_tag_list_callback(vueobj, response) {
  let json_obj = JSON.parse(response.data.data);
  console.log(json_obj);
  for (var key in json_obj) {
    let one_item = {};
    one_item.tag = json_obj[key]["Tag"];
    for (var oneUser in json_obj[key]["WhiteList"]) {
      one_item.whitelist =
        one_item.whitelist + ";" + json_obj[key]["WhiteList"][oneUser];
    }
    for (var oneUser in json_obj[key]["TempList"]) {
      one_item.templist =
        one_item.templist + ";" + json_obj[key]["TempList"][oneUser];
    }
    tableDataTemp.push(one_item);
    console.log("push", one_item.uid);
  }
  vueobj.tableData = tableDataTemp;
}
function modify_tag_callback(vueobj, response) {
  tableDataTemp = [];
  Utils.myajax("get_tag_list", "", get_tag_list_callback, vueobj);
}

function delete_callback(vueobj, response) {
  tableDataTemp = [];
  Utils.myajax("get_tag_list", "", get_tag_list_callback, vueobj);
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
    Utils.myajax("get_tag_list", "", get_tag_list_callback, this);
  },
  methods: {
    DialogOpen(val) {
      this.dialogFormVisible = true;
      this.form.tag = val["tag"];
      this.form.whitelist = val["whitelist"];
      this.form.templist = val["templist"];
      console.log("val=", val);
    },
    Add() {
      this.dialogFormVisible = true;
      this.form.tag = "";
      this.form.whitelist = "";
      this.form.templist = "";
    },
    ModifyTag() {
      console.log(this.form.tag, this.form.whitelist, this.form.templist);
      var Send = {};
      Send.Tag = this.form.tag;
      Send.WhiteList = this.form.whitelist.split(";");
      Send.Templist = this.form.templist.split(";");
      Utils.myajax(
        "modify_tag",
        JSON.stringify(Send),
        modify_tag_callback,
        this
      );
      this.dialogFormVisible = true;
    },
    Delete(val) {
      var tag = val["tag"];
      console.log("tag==", tag);
      Utils.myajax("del_tag", tag, delete_callback, this);
    },
    refresh() {
      console.log("refresh");
      tableDataTemp = [];
      Utils.myajax("get_tag_list", "", get_tag_list_callback, this);
      this.tableData = tableDataTemp;
    }
  }
};
</script>