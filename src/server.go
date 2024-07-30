package main

import (
	"context"
	"net/http"
	"time"
	"log"
	"golang.org/x/time/rate"
	 pb "lubdub.nl/aruba-iot/src/proto"
	"google.golang.org/protobuf/proto"
	"io"
	"nhooyr.io/websocket"
	"encoding/binary"
	"encoding/hex"
	"bytes"
//	"github.com/prometheus/client_golang/prometheus/promhttp"
//	"github.com/prometheus/client_golang/prometheus/promauto"
//	"github.com/prometheus/client_golang/prometheus"

)

// echoServer is the WebSocket echo server implementation.
// It ensures the client speaks the echo subprotocol and
// only allows one message every 100ms with a 10 message burst.
type echoServer struct {
	// logf controls where logs are sent.
	logf func(f string, v ...interface{})
	//Gauge prometheus.GaugeVec
}


func handleData(c <-chan pb.BleData){
	for packet := range c {
//		log.Println(packet)
		switch packet.GetFrameType(){
			case pb.BleFrameType_adv_ind:
				decodeData(packet.GetData(),packet.GetMac())
		}
	}
}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"194.171.98.6"},
		CompressionMode: websocket.CompressionContextTakeover,
	})


	datachan := make(chan pb.BleData, 10)

	go handleData(datachan)

	if err != nil {
		s.logf("%v", err)
		return
	}
	defer c.CloseNow()
	defer close(datachan)

	l := rate.NewLimiter(rate.Every(time.Millisecond*10), 100)
	for {
		err = echo(r.Context(), c, l,datachan)
		if err != nil {
			s.logf("failed to echo with %v: %v", r.RemoteAddr, err)
			return
		}
	}
}

type aranet_mf_data struct {
	Flags uint8
	Patch uint8
	Minor uint8
	Major uint8
}

type aranet4 struct {
	CO2 uint16
	Temp uint16
	Pressure uint16
	Humidity uint8
	Battery uint8
	Status uint8
	Interval uint16
	Age uint16
}

func printStruct(readout aranet4){
	log.Println("readout:")
	log.Println("CO2:\t",readout.CO2)
	log.Println("Temp:\t",readout.Temp/20)
	log.Println("Pressure:\t",readout.Pressure/10)
	log.Println("Humidity:\t",readout.Humidity)
	log.Println("Satus:\t",readout.Status)
	log.Println("Battery:\t",readout.Battery)
	log.Println("Age:\t",readout.Age,"/",readout.Interval)
}


func decodeData(data []byte,mac []byte) error  {
	// decode header
	index := 0
	header_len := data[index]
	index += int(header_len) + 1
//	log.Println("index:", index)
	// decode data
//	data_len := int(data[index])
//	log.Println("data_len:", data_len)
	index += 12
	payload_data := bytes.NewReader(data[index:])
	mf_data := bytes.NewReader(data[7:])
//	log.Println(data[7:index],data[index:index+13],data[index+13:])
//	log.Println(data)
	decoded_data := aranet4{}
	mfd_data := aranet_mf_data{}
	binary.Read(payload_data,binary.LittleEndian, &decoded_data)
	binary.Read(mf_data,binary.LittleEndian, &mfd_data)
	log.Println(hex.EncodeToString(mac))
	printStruct(decoded_data)
	return nil
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(ctx context.Context, c *websocket.Conn, l *rate.Limiter,dchan chan<- pb.BleData) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := l.Wait(ctx)
	if err != nil {
		return err
	}
	iot_blob := &pb.Telemetry{}
	_, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r)
	if err := proto.Unmarshal(b,iot_blob); err != nil {
		log.Fatalln("Failed to parse mssage", err)
	}
	if ble_data := iot_blob.GetBleData(); ble_data != nil{
		for _, data := range ble_data {
			if data != nil {
				dchan <- *data
			}
		}
	}
	if err != nil {
		return err
	}

	return err
}
