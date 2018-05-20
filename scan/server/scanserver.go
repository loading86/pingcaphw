package main

import (
	"sort"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net"
	"errors"
	pb "../common/scan"
	"golang.org/x/net/context"
    "google.golang.org/grpc"
)

const filename string = "tools/keyvals.map"

var keyvals = map[string]string{}

func readFile(filename string) (map[string]string, error) {
	var keyvals = map[string]string{}
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }
 
    if err := json.Unmarshal(bytes, &keyvals); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }
 
    return keyvals, nil
}

type keyVal struct{
	K string
	V string
}

type readOnlyArray struct{
	SerializedData []string 
	Indexes map[string]int32
}

func (readonly *readOnlyArray)Init(keyvals map[string]string){
	num := len(keyvals)
	fmt.Println(num)
	readonly.SerializedData = make([]string, num)
	readonly.Indexes = make(map[string]int32, num)
	keys := make([]string, 0, num)
	//var keys []string
	for k, v := range(keyvals){
		fmt.Printf("k %s, v:%s, len:%d\n", k, v, len(keys))
		keys = append(keys, k)
	}
	fmt.Println(len(keys))
	sort.Strings(keys)
	var kv keyVal
	for index, val := range(keys){
		readonly.Indexes[val] = int32(index)
		kv.K = val
		kv.V = keyvals[val]
		serialjson, _ := json.Marshal(kv)
		//fmt.Println(serialjson)
		readonly.SerializedData[index] = string(serialjson)
	}
	for _, val := range(readonly.SerializedData){
		fmt.Println(val)
	}
	for k, v := range(readonly.Indexes){
		fmt.Printf("k %s, v:%d\n", k, v)
	}
}

func (readonly readOnlyArray)GetNValsByKey(key string, num int32)([]string, int32, error){
	index, ok := readonly.Indexes[key]
	if !ok{
		return nil, 0, errors.New("not found key")
	}
	return readonly.GetNValsByIndex(index, num)
} 

func (readonly readOnlyArray)GetNValsByIndex(index int32, num int32)([]string, int32, error){
	var ret []string
	for i := int32(0); i < num; i++{
		if(i + index < int32(len(readonly.SerializedData))){
			ret = append(ret, readonly.SerializedData[i + index])
		}else{
			break;
		}
	}
	if(index + num < int32(len(readonly.SerializedData))){
		return ret, index + num, nil		
	}
	return ret, -1, nil
} 

//ServerHandler jsonrpc handler
type ServerHandler struct{
	readonly readOnlyArray
}

//FindByKey return vals after key
func (sh *ServerHandler)FindByKey(ctx context.Context, req *pb.FindRequest) (*pb.FindResult, error) {
	data, next, err := sh.readonly.GetNValsByKey(*req.Start, *req.Num)
	var res pb.FindResult
	var errcode int32
	if(err != nil){
		errcode = 1
		res.Err = &errcode
	}else{
		res.Err = &errcode
		res.Kvpairs = data
	}
	res.Nextindex = &next
	return &res, nil
}

//FindByIndex return vals after start
func (sh *ServerHandler)FindByIndex(ctx context.Context, req *pb.FindRequestByIndex) (*pb.FindResult, error) {
	data, next, err := sh.readonly.GetNValsByIndex(*req.Start, *req.Num)
	var res pb.FindResult
	var errcode int32
	if(err != nil){
		errcode = 1
		res.Err = &errcode
	}else{
		res.Err = &errcode
		res.Kvpairs = data
	}
	res.Nextindex = &next
	return &res, nil
}

func main(){
	keyvals, err := readFile(filename)
	if(err != nil){
		fmt.Printf("cant read json from %s", filename)
		return
	}
	fmt.Printf("keyvals %v", keyvals)
	var readonly readOnlyArray
	readonly.Init(keyvals)
	server := ServerHandler{readonly}
	lis, errl := net.Listen("tcp", "127.0.0.1:8888")
    if errl != nil {
        fmt.Println("failed to listen")
    }
    s := grpc.NewServer()
    pb.RegisterScanerServer(s, &server)
    s.Serve(lis)
}