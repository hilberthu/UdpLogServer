package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"utils"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

var G_DBConn *sqlite3.Conn

type ServerConf struct {
	HttpPort string
	UdpPort  string
}

type LogContent struct {
	UID     int64
	Level   string
	Tag     string
	Content string
}

var G_StConf ServerConf

func loadConf() {
	fi, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	//utils.Debugln("loadConf:", string(fd))
	err = json.Unmarshal(fd, &G_StConf)
	if err != nil {
		utils.Debugln("error:", err)
		//os.Exit(0)
	}
	fmt.Println(G_StConf)
	return
}

func FunctionMapCall(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func querylog(w http.ResponseWriter, r *http.Request, strData string) {
	retMap := GetRetMap()
	vStr := strings.Split(strData, "|")
	UID, _ := strconv.ParseInt(vStr[0], 10, 64)
	StartTime, _ := strconv.ParseInt(vStr[1], 10, 64)
	EndTime, _ := strconv.ParseInt(vStr[2], 10, 64)
	fmt.Println("param==", UID, StartTime, EndTime)

	stmt, err := G_DBConn.Prepare(`SELECT * FROM loginfo WHERE uid = ? and time >=? and  time<=?`, UID, StartTime, EndTime)
	var uid int64
	var level int
	var rcvTime int64
	var tag string
	var content string
	for {
		hasRow, err := stmt.Step()
		if !hasRow {
			break
		}
		err = stmt.Scan(&uid, &level, &rcvTime, &tag, &content)
		if err != nil {
			break
		}
		fmt.Println("uid", uid)
	}
	err = stmt.Scan(&uid, &level, &rcvTime, &tag, &content)
	if err != nil {
		checkErr(err)
	}
	fmt.Println("uid", uid)

	var nUids []int64
	json.Unmarshal([]byte(strData), &nUids)
	fmt.Fprint(w, RetMap2String(retMap))
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	utils.Debugln("ajaxHandler begin---------------")
	funcs := make(map[string]interface{})
	funcs["querylog"] = querylog
	strAction := r.FormValue("action")
	strData := r.FormValue("data")
	utils.Debugln("ajaxHandler---------------", strAction, strData)
	FunctionMapCall(funcs, strAction, w, r, strData)
}

func GetRetMap() map[string]string {
	var retDataMap map[string]string = make(map[string]string)
	retDataMap["ret"] = "0"
	retDataMap["data"] = ""
	return retDataMap
}

func RetMap2String(retMap map[string]string) string {
	v, _ := json.Marshal(retMap)
	return string(v)
}

func handleUpload(w http.ResponseWriter, request *http.Request) {
	BaseUploadPath := "./LogFiles"
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(w, "Read file error")
		log.Println("err==: " + err.Error())
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	log.Println("filename: " + fileHeader.Filename)
	vData := strings.Split(fileHeader.Filename, "_")
	if len(vData) < 7 {
		_, _ = io.WriteString(w, "file format error")
		return
	}
	destPath := BaseUploadPath + "/" + vData[0]
	os.Mkdir(destPath, os.ModeDir)
	//创建文件
	newFile, err := os.Create(destPath + "/" + fileHeader.Filename)
	if err != nil {
		_, _ = io.WriteString(w, "Create file error:"+destPath+"/"+fileHeader.Filename)
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(w, "Write file error")
		return
	}
	_, _ = io.WriteString(w, "Upload success")
}

func handleDownload(w http.ResponseWriter, request *http.Request) {
	BaseUploadPath := "./Logfiles"
	//文件上传只允许GET方法
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}
	//文件名
	filename := request.FormValue("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	log.Println("filename: " + filename)
	//打开文件
	file, err := os.Open(BaseUploadPath + "/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}

func getFilelist(path string) {
	for {
		var diff_time int64 = 24 * 3600 * 3
		now_time := time.Now().Unix() //当前时间，使用Unix时间戳
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			file_time := f.ModTime().Unix()
			if (now_time - file_time) > diff_time { //判断文件是否超过3天
				fmt.Printf("Delete file %v !\r\n", path)
				println(path)
				if path != "./LogFiles/" {
					os.RemoveAll(path)
				}

			} else {
				println(path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("filepath.Walk() returned %v\r\n", err)
		}
		time.Sleep(time.Second * 60)
	}

}
func StartWeb() error {
	//http.Handle("/", http.FileServer(http.Dir("./log")))
	//http://192.168.122.88:10021/?filename=conf.json
	http.HandleFunc("/ajax", ajaxHandler)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/download", handleDownload)
	fsh := http.FileServer(http.Dir("./dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fsh))
	fsh2 := http.FileServer(http.Dir("./LogFiles"))
	http.Handle("/LogFiles/", http.StripPrefix("/LogFiles/", fsh2))
	err := http.ListenAndServe(G_StConf.HttpPort, nil)
	if err != nil {
		fmt.Println("http listen " + G_StConf.HttpPort)
	} else {
		fmt.Println("http listen failed " + G_StConf.HttpPort)
	}
	return err

}

func DealWithUdpPkg() {
	//f, err1 := os.OpenFile("alian", os.O_APPEND, 0777)
	udpAddr, err := net.ResolveUDPAddr("udp", G_StConf.UdpPort)
	if err != nil {
		log.Fatalln("Error: ", err)
		//os.Exit(0)
	}

	// Build listining connections
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln("Error: ", err)
		//os.Exit(0)
	}
	defer conn.Close()
	// Interacting with one client at a time
	recvBuff := make([]byte, 1500)
	for {
		//log.Println("Ready to receive packets!")
		// Receiving a message
		rn, clientaddr, err := conn.ReadFromUDP(recvBuff)
		if rn == 1 {
			conn.WriteToUDP([]byte(recvBuff[0:1]), clientaddr)
		}
		if err != nil {
			utils.Errorln("Error:", err.Error())
			time.Sleep(time.Millisecond * 30)
			continue
		}

		utils.Debugln("DealWithUdpPkg", string(recvBuff))
		recvBuff2 := recvBuff[:rn]
		var logPkg LogContent
		err = json.Unmarshal(recvBuff2, &logPkg)
		if err != nil {
			utils.Errorln("Error:", err.Error())
			continue
		}

		stmt, err := G_DBConn.Prepare("INSERT INTO loginfo(uid, level, time,tag, content) values(?,?,?,?,?)")
		if err != nil {
			utils.Errorln("Error:", err.Error())
			continue
		}
		timeNow := time.Now().Unix()
		err = stmt.Exec(logPkg.UID, logPkg.Level, timeNow, logPkg.Tag, logPkg.Content)
		if err != nil {
			utils.Errorln("Error:", err.Error())
			continue
		}

	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//StartMysqlWriteTimer()
	utils.SetLevel(4)
	loadConf()
	var err error
	G_DBConn, err = sqlite3.Open("test.sqlite")
	if err != nil {
		checkErr(err)
	}
	defer G_DBConn.Close()

	// It's always a good idea to set a busy timeout
	G_DBConn.BusyTimeout(5 * time.Second)
	sql_table := `
    CREATE TABLE IF NOT EXISTS loginfo(
        uid INTEGER ,
        level INTEGER,
		time INTEGER,
		tag VARCHAR(32),
        content mediumblob NULL
    );`
	err = G_DBConn.Exec(sql_table)
	if err != nil {
		checkErr(err)
	}
	createIndex := `CREATE INDEX IndexUid ON loginfo (uid);`
	G_DBConn.Exec(createIndex)
	//stmt, err := db.Prepare("INSERT INTO loginfo(uid, level, time,tag, content) values(?,?,?,?,?)")
	//checkErr(err)
	//err = stmt.Exec(1023, 1, 20198888, "tag1", []byte("content"))
	checkErr(err)
	stmt, err := G_DBConn.Prepare(`SELECT * FROM loginfo WHERE uid = ?`, 1034)
	hasRow, err := stmt.Step()
	checkErr(err)
	if !hasRow {
		fmt.Println("don't have row")
	}
	var uid int64
	var level int
	var rcvTime int64
	var tag string
	var content string
	err = stmt.Scan(&uid, &level, &rcvTime, &tag, &content)
	if err != nil {
		checkErr(err)
	}

	go DealWithUdpPkg()
	time.Sleep(time.Second)
	udpconn, err := net.Dial("udp", "127.0.0.1:9016")
	defer udpconn.Close()
	if err != nil {
		fmt.Println("net.Dial error", err.Error())
	}
	var stLogContent LogContent
	stLogContent.UID = 1034
	stLogContent.Level = "debug"
	stLogContent.Tag = "fadf"
	stLogContent.Content = "testcontent"
	logBuff, _ := json.Marshal(stLogContent)
	//"uid|level|"
	udpconn.Write(logBuff)
	udpconn.Write([]byte("0"))
	var rdbuff []byte
	udpconn.Read(rdbuff)
	fmt.Println(string(rdbuff))
	go getFilelist("./LogFiles/")
	StartWeb()

}
