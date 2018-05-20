syntax = "proto2";


package scan;

// The scan service definition.
service Scaner {
  // Sends a greeting
  rpc FindByKey (FindRequest) returns (FindResult) {}
  rpc FindByIndex (FindRequestByIndex) returns (FindResult) {}
}


// The request message containing the start key and fetch num
message FindRequest {
  required string start = 1;
  required int32 num = 2;
}

// The request message containing the start index and fetch num
message FindRequestByIndex {
  required int32 start = 1;
  required int32 num = 2;
}

// The response message containing the errcode, data and the next fetch index
message FindResult {
  required int32 err = 1;
  repeated string kvpairs = 2;
  required int32 nextindex = 3;
}

