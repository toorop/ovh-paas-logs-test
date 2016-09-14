package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/satori/go.uuid"

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
	/*tcpAddr, err := net.ResolveTCPAddr("tcp", "laas.runabove.com:2204")
	if err != nil {
		return err
	}*/

	// TLS
	conf := &tls.Config{}
	conn, err := tls.Dial("tcp", "laas.runabove.com:12204", conf)
	// no TLS
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	_, err = conn.Write(b)

	//log.Println(err, " - ", n)
	return err
}

func GetRandString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func sendRandomLog() (err error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return
	}

	record, err := NewRootRecord(seg)
	if err != nil {
		return
	}

	record.SetTs(float64(time.Now().Unix()))
	err = record.SetHostname("home")
	if err != nil {
		return
	}
	record.SetFacility(5)
	record.SetSeverity(4)
	err = record.SetAppname("MyApp")
	if err != nil {
		return
	}
	err = record.SetProcid("procid")
	if err != nil {
		return
	}
	err = record.SetMsgid(uuid.NewV4().String())
	if err != nil {
		return
	}
	err = record.SetMsg("message: " + GetRandString(10))
	if err != nil {
		return
	}
	err = record.SetFullMsg("full message: " + GetRandString(50))
	if err != nil {
		return
	}
	err = record.SetSdId(uuid.NewV4().String())
	if err != nil {
		return
	}
	// pairs 4 X-OVH-TOKEN
	pairList, err := NewPair_List(seg, 1)
	if err != nil {
		return
	}

	pair, err := NewPair(seg)
	if err != nil {
		return
	}

	err = pair.SetKey("X-OVH-TOKEN")
	if err != nil {
		return
	}
	err = pair.Value().SetString(ovhToken)
	if err != nil {
		return
	}
	err = pairList.Set(0, pair)
	if err != nil {
		return
	}
	err = record.SetPairs(pairList)
	if err != nil {
		return
	}

	// encode
	b := bytes.NewBuffer([]byte{})
	if err = capnp.NewEncoder(b).Encode(msg); err != nil {
		return err
	}

	return send(b.Bytes())

}

func init() {
	ovhToken = os.Getenv("OVH_TOKEN")
	if ovhToken == "" {
		log.Fatalln("OVH_TOKEN not set in ENV")
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := sendRandomLog(); err != nil {
				log.Println(i, err)
			}
		}()
	}
	wg.Wait()

	// test decode
	/*decMsg, err := capnp.NewDecoder(b).Decode()
	handleERR(err)
	record2, err := ReadRootRecord(decMsg)
	handleERR(err)
	pairs, err := record2.Pairs()
	handleERR(err)
	ovhToken, err := pairs.At(0).Value().String()
	handleERR(err)
	log.Println(string(ovhToken))
	*/
}
