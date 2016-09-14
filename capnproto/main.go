package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"

	"zombiezen.com/go/capnproto2"
)

var ovhToken string

func handleERR(err error) {
	if err != nil {
		panic(err)
	}
}

// Send msg to RA
func send(b []byte) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "laas.runabove.com:2204")
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	n, err := conn.Write(b)

	log.Println(err, " - ", n)
	return err
}

func init() {
	ovhToken = os.Getenv("OVH_TOKEN")
	if ovhToken == "" {
		log.Fatalln("OVH_TOKEN not set in ENV")
	}
}

func main() {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	record, err := NewRecord(seg)
	handleERR(err)

	record.SetTs(float64(time.Now().Unix()))
	handleERR(record.SetHostname("home"))
	record.SetFacility(5)
	record.SetSeverity(4)
	handleERR(record.SetAppname("MyApp"))
	handleERR(record.SetProcid("procid"))
	handleERR(record.SetMsgid("msgid"))
	handleERR(record.SetMsg("message"))
	handleERR(record.SetFullMsg("full message"))
	handleERR(record.SetSdId("ID"))

	// pairs
	pairList, err := NewPair_List(seg, 1)
	handleERR(err)

	pair, err := NewPair(seg)
	handleERR(err)
	handleERR(pair.SetKey("X-OVH-TOKEN"))
	handleERR(pair.Value().SetString(ovhToken))
	handleERR(pairList.Set(0, pair))
	handleERR(record.SetPairs(pairList))

	// serialize
	b := bytes.NewBuffer([]byte{})
	handleERR(capnp.NewEncoder(b).Encode(msg))

	handleERR(send(b.Bytes()))

}
