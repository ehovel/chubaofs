package master

import (
	"time"
	"testing"
	"fmt"
	"github.com/chubaofs/chubaofs/proto"
)

func TestDataNode(t *testing.T) {
	// /dataNode/add and /dataNode/response processed by mock data server
	addr := "127.0.0.1:9096"
	addDataServer(addr,DefaultRackName)
	server.cluster.checkDataNodeHeartbeat()
	time.Sleep(5 * time.Second)
	getDataNodeInfo(addr, t)
	decommissionDataNode(addr, t)
	_, err := server.cluster.dataNode(addr)
	if err == nil {
		t.Errorf("decommission datanode [%v] failed", addr)
	}
	server.cluster.dataNodes.Delete(addr)
}

func getDataNodeInfo(addr string, t *testing.T) {
	reqUrl := fmt.Sprintf("%v%v?addr=%v", hostAddr, proto.GetDataNode, addr)
	fmt.Println(reqUrl)
	process(reqUrl, t)
}

func decommissionDataNode(addr string, t *testing.T) {
	reqUrl := fmt.Sprintf("%v%v?addr=%v", hostAddr, proto.DecommissionDataNode, addr)
	fmt.Println(reqUrl)
	process(reqUrl, t)
}
