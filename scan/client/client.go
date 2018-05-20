package main

import (
	"errors"
	"fmt"
	pb "../common/scan"
	"golang.org/x/net/context"
    "google.golang.org/grpc"
)


type ScanClient struct{
	rpcClient pb.ScanerClient 
}


type ScanIter struct{
	client ScanClient
	vals []string
	index int32
	nextindex int32
}

func (c ScanClient)Find(start string, num int32)(*ScanIter, error){
	var req pb.FindRequest
	req.Start = &start
	req.Num = &num 
	res, e := c.rpcClient.FindByKey(context.Background(), &req)
	if e != nil {
		// log.Println("Error call rpc method:", e)
		// http.Error(w, e.Error(), http.StatusInternalServerError)
		return nil, errors.New("cant find key")
	}
	fmt.Printf("res reply:%v\n", res)
	var iter ScanIter
	iter.client = c
	iter.vals = res.Kvpairs
	iter.index = 0
	iter.nextindex = *res.Nextindex
	return &iter, nil
}

func (iter *ScanIter)Next(num int32) error {
	if(iter.nextindex == -1){
		return errors.New("have no more data")
	}
	var req pb.FindRequestByIndex
	req.Start = &(iter.nextindex)
	req.Num = &num  
	res, e := iter.client.rpcClient.FindByIndex(context.Background(), &req)
	if e != nil {
		// log.Println("Error call rpc method:", e)
		// http.Error(w, e.Error(), http.StatusInternalServerError)
		return errors.New("cant find key")
	}
	fmt.Printf("res reply:%v\n", res)
	iter.vals = res.Kvpairs
	iter.index = 0
	iter.nextindex = *(res.Nextindex)
	return nil
}

func main(){
	// RPC calls.
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		// log.Println("Error dail rpc server:", e)
		// http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	} 
	client := ScanClient{pb.NewScanerClient(conn)}
	iter, err2 := client.Find("oTCJcwpZPy", 0)
	if(err2 != nil){
		fmt.Println("Find err")
		return
	}
	iter.Next(10)
	fmt.Printf("res iter:%v\n", iter)	
}