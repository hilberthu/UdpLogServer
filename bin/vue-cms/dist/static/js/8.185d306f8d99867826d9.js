webpackJsonp([8],{"4Txg":function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var l=a("Q5rI"),s=[];function r(t,e){var a=JSON.parse(e.data.data);for(var l in console.log(a),a){var r={};r.LevelId=a[l].LevelId,r.PlayTimes=a[l].PlayTimes,r.PassTimes=a[l].PassTimes,r.PassRation=a[l].PassRation,s.push(r),console.log("push",r.uid)}s.sort(function(t,e){return parseInt(t.LevelId)-parseInt(e.LevelId)}),t.tableData=s}function i(t,e){s=[],l.a.myajax("get_pass_ratio","",r,this),t.tableData=s}var n={data:function(){return{tableData:s,dialogTableVisible:!1,dialogFormVisible:!1,form:{tag:"",whitelist:"",templist:""},formLabelWidth:"120px"}},created:function(){l.a.myajax("get_pass_ratio","",r,this)},methods:{Delete:function(t){var e=t.LevelId;l.a.myajax("del_pass_ratio",e,i,this)},refresh:function(){console.log("refresh"),s=[],l.a.myajax("get_pass_ratio","",r,this),this.tableData=s}}},o={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData,stripe:""}},[a("el-table-column",{attrs:{prop:"LevelId",label:"关卡ID",width:"180"}}),t._v(" "),a("el-table-column",{attrs:{prop:"PlayTimes",label:"玩的次数",width:"180"}}),t._v(" "),a("el-table-column",{attrs:{prop:"PassTimes",label:"通关次数",width:"180"}}),t._v(" "),a("el-table-column",{attrs:{prop:"PassRation",label:"通关率",width:"180"}}),t._v(" "),a("el-table-column",{attrs:{label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",size:"small"},on:{click:function(a){return t.Delete(e.row)}}},[t._v("重置")])]}}])})],1),t._v(" "),a("el-button",{attrs:{size:"big"},on:{click:function(e){return t.refresh()}}},[t._v("刷新")])],1)},staticRenderFns:[]},u=a("VU/8")(n,o,!1,null,null,null);e.default=u.exports}});
//# sourceMappingURL=8.185d306f8d99867826d9.js.map