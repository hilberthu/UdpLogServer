webpackJsonp([9],{f11k:function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var l=i("mvHQ"),a=i.n(l),o=i("Q5rI"),s=[];function r(t,e){var i=JSON.parse(e.data.data);for(var l in console.log(i),i){var a={};for(var o in a.tag=i[l].Tag,i[l].WhiteList)a.whitelist=a.whitelist+";"+i[l].WhiteList[o];for(var o in i[l].TempList)a.templist=a.templist+";"+i[l].TempList[o];s.push(a),console.log("push",a.uid)}t.tableData=s}function n(t,e){s=[],o.a.myajax("get_tag_list","",r,t)}function m(t,e){s=[],o.a.myajax("get_tag_list","",r,t)}var f={data:function(){return{tableData:s,dialogTableVisible:!1,dialogFormVisible:!1,form:{tag:"",whitelist:"",templist:""},formLabelWidth:"120px"}},created:function(){o.a.myajax("get_tag_list","",r,this)},methods:{DialogOpen:function(t){this.dialogFormVisible=!0,this.form.tag=t.tag,this.form.whitelist=t.whitelist,this.form.templist=t.templist,console.log("val=",t)},Add:function(){this.dialogFormVisible=!0,this.form.tag="",this.form.whitelist="",this.form.templist=""},ModifyTag:function(){console.log(this.form.tag,this.form.whitelist,this.form.templist);var t={};t.Tag=this.form.tag,t.WhiteList=this.form.whitelist.split(";"),t.Templist=this.form.templist.split(";"),o.a.myajax("modify_tag",a()(t),n,this),this.dialogFormVisible=!0},Delete:function(t){var e=t.tag;console.log("tag==",e),o.a.myajax("del_tag",e,m,this)},refresh:function(){console.log("refresh"),s=[],o.a.myajax("get_tag_list","",r,this),this.tableData=s}}},c={render:function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData,stripe:""}},[i("el-table-column",{attrs:{prop:"tag",label:"Tag",width:"180"}}),t._v(" "),i("el-table-column",{attrs:{prop:"whitelist",label:"白名单",width:"180"}}),t._v(" "),i("el-table-column",{attrs:{prop:"templist",label:"临时允许提交名单",width:"180"}}),t._v(" "),i("el-table-column",{attrs:{label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[i("el-button",{attrs:{type:"text",size:"small"},on:{click:function(i){return t.DialogOpen(e.row)}}},[t._v("修改")]),t._v(" "),i("el-button",{attrs:{type:"text",size:"small"},on:{click:function(i){return t.Delete(e.row)}}},[t._v("删除")])]}}])})],1),t._v(" "),i("el-button",{attrs:{size:"big"},on:{click:function(e){return t.refresh()}}},[t._v("刷新")]),t._v(" "),i("el-button",{attrs:{size:"big"},on:{click:function(e){return t.Add()}}},[t._v("添加")]),t._v(" "),i("el-dialog",{attrs:{title:"tag编辑",visible:t.dialogFormVisible},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[i("el-form",{attrs:{model:t.form}},[i("el-form-item",{attrs:{label:"tag","label-width":t.formLabelWidth}},[i("el-input",{attrs:{autocomplete:"off"},model:{value:t.form.tag,callback:function(e){t.$set(t.form,"tag",e)},expression:"form.tag"}})],1),t._v(" "),i("el-form-item",{attrs:{label:"白名单","label-width":t.formLabelWidth}},[i("el-input",{attrs:{autocomplete:"off"},model:{value:t.form.whitelist,callback:function(e){t.$set(t.form,"whitelist",e)},expression:"form.whitelist"}})],1),t._v(" "),i("el-form-item",{attrs:{label:"临时权限","label-width":t.formLabelWidth}},[i("el-input",{attrs:{autocomplete:"off"},model:{value:t.form.templist,callback:function(e){t.$set(t.form,"templist",e)},expression:"form.templist"}})],1)],1),t._v(" "),i("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[i("el-button",{on:{click:function(e){t.dialogFormVisible=!1}}},[t._v("取 消")]),t._v(" "),i("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.ModifyTag()}}},[t._v("确 定")])],1)],1)],1)},staticRenderFns:[]},u=i("VU/8")(f,c,!1,null,null,null);e.default=u.exports}});
//# sourceMappingURL=9.da01c4bd100b8f3dcf9a.js.map