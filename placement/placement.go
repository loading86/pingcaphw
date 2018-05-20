package main


//Store define a server that contains a tikv server process
type Store struct{
	ID int
	RackID int
	DcID int
	Regions map[int]bool //used as set
	RegionNum int
	Labels map[string]string //key should be mem,cpu,storage and so on
}

//Rack define a rack whick have some stores
type Rack struct{
	ID int
	DcID int
	Stores map[int]Store
}

func (r *Rack)GetBestStore() *Store{
	
}

//Datacenter define a datacenter which have some racks
type Datacenter struct{
	ID int
	Racks map[int]Rack 
}

//Cluster define overview of all datacenters
type Cluster struct{
	Dcs map[int]Datacenter
}

var cluster Cluster

//Region specify stores which contains same data trunk
type Region struct{
	ID int 
	Replicas []int
}

type Placement struct{
	RegionLocation map[int]Region
	StoreRegions map[string][]int
}

type Strategy interface{
	GetNewRegion(stores []Store, place Placement) error
}

type ThreeDcStrategy struct{

}

func (stg ThreeDcStrategy)GetNewRegion(stores []Store, place Placement) error{
	for 
}






func Check(stores []Store, region Region, strategy Strategy) Region{

}