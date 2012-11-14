package main

import (
	"bufio"
	l4g "code.google.com/p/log4go"
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

	l4g.AddFilter("stdout", l4g.INFO, l4g.NewConsoleLogWriter())

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
		num = num + 1
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

	queryTick := time.Tick(50 * time.Second)

	for {
		select {
		case <-queryTick:
			// trying the manual method from dht_test in case I have the format 
			// wrong
			fmt.Printf("TICK\n")
			for i := 0; i < num; i++ {
				l4g.Info("querying for infoHash: %x", shalist[i])
				go dht.PeersRequest(shalist[i], sendAnnouncements)
				time.Sleep(10000 * time.Millisecond)
			}
		}
	}
}

// drainresults loops, constantly reading any new peer information sent by the
// DHT node and just ignoring them. We don't care about those :-P.
func drainresults(n *dht.DHT) {
	for infoHashPeers := range n.PeersRequestResults {
		for ih, peers := range infoHashPeers {
			if len(peers) > 0 {
				fmt.Printf("peer found for infohash [%x] ", ih)
				for _, peer := range peers {
					fmt.Println(dht.DecodePeerAddress(peer))
				}
			}
		}
	}
	l4g.Info("finishing drainresults")
}
