import axios from 'axios'
function myajax(action, data, callback, vueobj) {
    let strUrl = "/ajax?" + "action=" + action + "&data=" + data;
    return axios.get(strUrl).then(response => callback(vueobj, response))
}
export default{
    myajax
}