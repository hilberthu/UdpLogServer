<template>
  <div class="block">
    <span class="demonstration">邮件生效时间与过期时间:</span>
    <el-date-picker
      v-model="begin_end"
      type="datetimerange"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
    ></el-date-picker>
    <el-divider></el-divider>
    <div class="demo-input-suffix" style="left: 50%">
      邮件标题：
      <el-input v-model="mail_title" placeholder="请输入邮件标题"></el-input>
    </div>
    <el-divider></el-divider>
    <div class="demo-input-suffix">
      邮件正文：
      <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="mail_text"></el-input>
    </div>
    <el-divider></el-divider>
    <el-table :data="tableData" stripe>
      <el-table-column prop="Id" label="Id" width="180"></el-table-column>
      <el-table-column prop="Name" label="资源名字" width="180"></el-table-column>
      <el-table-column prop="Count" label="数量"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button @click="DelAttach(scope.row)" type="text" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button type="text" @click="dialogFormVisible = true">添加附件</el-button>
    <el-button type="text" @click="Send()">发送</el-button>

    <el-dialog title="添加附件" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="请选择附件" :label-width="formLabelWidth">
          <el-select v-model="form.region" placeholder="请选择附件">
            <el-option
              v-for="item in selectRoleList"
              :key="item.id"
              :lable="item.id"
              :value="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <div class="demo-input-suffix">
          数量：
          <el-input v-model="attach_count" placeholder="1"></el-input>
        </div>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="add_attach()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Utils from "@/api/utils";
let obj = [
  {
    id: "fishboost",
    name: "关前道具-飞机"
  },
  {
    id: "bombAndLineBoost",
    name: "关前道具-炸弹加排消"
  },
  {
    id: "ballboost",
    name: "关前道具-彩虹球"
  },
  {
    id: "100",
    name: "关内道具-锤子"
  },
  {
    id: "101",
    name: "关内道具-十字锤"
  },
  {
    id: "102",
    name: "关内道具-交换手"
  },
  {
    id: "211",
    name: "金币"
  },
  {
    id: "212",
    name: "体力"
  },
  {
    id: "213",
    name: "星星"
  },
  {
    id: "214",
    name: "装饰币"
  },
  {
    id: "215",
    name: "限时体力"
  },
  {
    id: "newhall_1571798786_0001",
    name: "大厅花坛"
  },
  {
    id: "newhall_1571798790_0001",
    name: "装饰画"
  }
];

function ConvertItemName2Id(name) {
  for (var index in obj) {
    if (obj[index].name == name) {
      return obj[index].id;
    }
  }
  return "";
}

function add_sysmail_callback(vueobj, response) {
  console.log(response);
  alert("发送成功");
}
export default {
  data() {
    return {
      tableData: [],
      dialogFormVisible: false,
      attach_count: 1,
      formLabelWidth: "120px",
      form: {
        region: ""
      },
      begin_end: [],
      mail_text: "",
      mail_title: ""
    };
  },

  computed: {
    selectRoleList() {
      //初始化下拉框动态数据
      return obj;
    }
  },
  methods: {
    add_attach() {
      console.log("form.region==", this.form.region);
      var id = ConvertItemName2Id(this.form.region);
      console.log("id==", id);
      console.log("count==", this.attach_count);
      this.dialogFormVisible = false;
      this.tableData.push({
        Id: id,
        Name: this.form.region,
        Count: this.attach_count,
        Index: this.tableData.length
      });
      console.log("tableData==", this.tableData);
    },
    DelAttach(row) {
      this.tableData.splice(row.index, 1);
      console.log("this.tableData==", this.tableData);
    },
    Send() {
      var SendData = {};
      SendData.BeginTime = Date.parse(this.begin_end[0]) / 1000;
      SendData.EndTime = Date.parse(this.begin_end[1]) / 1000;
      SendData.MailTitil = this.mail_title;
      SendData.Text = this.mail_text;
      SendData.AttachList = this.tableData;

      if (SendData.EndTime < SendData.BeginTime) {
        alert("邮件过期时间小于生效事件！");
        return;
      }
      if (SendData.MailTitil == "") {
        alert("邮件标题不能位空！");
        return;
      }
      if (SendData.Text == "") {
        alert("邮件内容不能位空！");
        return;
      }
      let strSendData = JSON.stringify(SendData);
      Utils.myajax("add_system_mail", strSendData, add_sysmail_callback, this);
      console.log("SendData");
    }
  }
};
</script>