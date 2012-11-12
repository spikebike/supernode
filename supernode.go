package main

import (
	"bufio"
	l4g "code.google.com/p/log4go"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/nictuku/dht"
	"io"
	"os"
	"time"
)

func main() {
	var numTargetPeers int
	var sendAnnouncements bool
	var file *os.File
	var num int
	var shalist [64]string
	var err error
	var part []byte

	if file, err = os.Open("hops.log"); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	num = 0
	for {
		if part, _, err = reader.ReadLine(); err != nil {
			break
		}
		if err == io.EOF {
			err = nil
		}
		hasher := sha1.New()
		hasher.Write(part)
		shalist[num] = string(hasher.Sum(nil))
		//		fmt.Printf("%s %x\n", part, shalist[num])
		num = num + 1
	}
	for i := 0; i < num; i++ {
		fmt.Printf("%d %x\n", i, shalist[i])
	}

	sendAnnouncements = true
	numTargetPeers = 64

	port := 42345
	l4g.Error("used port %d", port)
	dht, err := dht.NewDHTNode(port, numTargetPeers, true)
	if err != nil {
		l4g.Error("DHT node creation error", err)
		return
	}

	go dht.DoDHT()
	go drainresults(dht)

	queryTick := time.Tick(10 * time.Second)
	quickTick := time.Tick(1 * time.Second)

	for {
		select {
		case <-queryTick:
			fmt.Printf("TICK\n");
			for i := 0; i < num; i++ {
				fmt.Printf("querying for infoHash: %x\n", shalist[i])
				go dht.PeersRequest(shalist[i], sendAnnouncements)
				<-quickTick
			}
		}
	}
}

func nodeId() []byte {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		l4g.Exitf("nodeId rand: %v", err)
	}
	return b
}

// drainresults loops, constantly reading any new peer information sent by the
// DHT node and just ignoring them. We don't care about those :-P.
func drainresults(n *dht.DHT) {
	infoHashPeers := <-n.PeersRequestResults
	for ih, peers := range infoHashPeers {
		if len(peers) > 0 {
			fmt.Printf("peer found for infohash [%x]\n", ih)
			for _, peer := range peers {
				fmt.Println(dht.DecodePeerAddress(peer))
			}
			return
		}
	}

}
