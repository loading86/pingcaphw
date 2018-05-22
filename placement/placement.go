package main


const {
	RegionMaxSize = 64 //64 MB
}

//Store define a server that contains a tikv server process
type Store struct{
	ID int
	RackID int
	DcID int
	Regions map[int]bool //used as set
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
	Stores map[int]Store
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

type ThreeRacksStrategy struct{
}

type AllSsdStrategy struct{	
}

type OneSsdStrategy struct{	
}

func (stg ThreeDcStrategy)GetNewRegion(cluster Cluster, place Placement) error{
	for 
}



func (stg ThreeDcStrategy)Satisfied(cluster Cluster, regin Region) bool {
	dcs := map[int]bool
	for _, r := range(regin.Replicas){
		store, ok := cluster.Stores[r]
		if !ok {
			return false
		}
		dcs[store.DcID] = true
	}
	if(len(dcs) == 3){
		return true
	}
	return false
}


func (stg ThreeRacksStrategy)Satisfied(cluster Cluster, regin Region) bool {
	racks := map[int]bool
	for _, r := range(regin.Replicas){
		store, ok := cluster.Stores[r]
		if !ok {
			return false
		}
		racks[store.RackID] = true
	}
	if(len(racks) == 3){
		return true
	}
	return false
}

func (stg AllSsdStrategy)Satisfied(stores []Store, regin Region){
	for _, r := range(regin.Replicas){
		store, ok := cluster.Stores[r]
		if !ok {
			return false
		}
		if(store.Labels["storagetype"] != "ssd"){
			return false
		}
	}
	return true
}

func (stg OneSsdStrategy)Satisfied(stores []Store, regin Region){
	ssdnum := 0
	for _, r := range(regin.Replicas){
		store, ok := cluster.Stores[r]
		if !ok {
			return false
		}
		if(store.Labels["storagetype"] == "ssd"){
			ssdnum += 1
		}
	}
	return ssdnum == 1
}

func Check(stores []Store, region Region, strategy Strategy) Region{

}
