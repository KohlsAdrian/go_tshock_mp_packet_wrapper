package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"kohlsadrian.com/tshock/packets"
)

var count = 0

var lastPacket map[string]interface{}

// Packet data
type Packet struct {
	PlayerID     int
	ID           int
	Index        int
	Length       int
	PacketBuffer []byte
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleData).Methods("POST")
	router.HandleFunc("/lastPacket", handleLastPacket).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:10000", router))
}

func handleLastPacket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lastPacket)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err == nil {
		responseString := string(bodyBytes)
		vars := strings.Split(responseString, "&")

		packetPlayerID := strings.Split(vars[0], "whoAmI=")[1]
		packetID := strings.Split(vars[1], "packetId=")[1]
		packetIndex := strings.Split(vars[2], "index=")[1]
		packetLength := strings.Split(vars[3], "length=")[1]
		packetBufferVar := strings.Split(vars[4], "packetData=")[1]

		// Decode packet data utf8 to terraria bytes
		playerID, errPlayerID := strconv.Atoi(packetPlayerID)
		id, errID := strconv.Atoi(packetID)
		index, errINDEX := strconv.Atoi(packetIndex)
		length, errLENGTH := strconv.Atoi(packetLength)
		packetDataBytes, err := hex.DecodeString(packetBufferVar)

		if err != nil {
			fmt.Printf("error decode - %d\n", id)
		}

		if errPlayerID == nil && errID == nil && errINDEX == nil && errLENGTH == nil {
			packet := Packet{
				PlayerID:     playerID,
				ID:           id,
				Index:        index,
				Length:       length,
				PacketBuffer: packetDataBytes,
			}
			handlePacket(packet)
		}
	}
}

func handlePacket(packet Packet) {
	count++
	ReadBinary(packet)
}

/**
Terraria Multiplayer Packet
*/

// ReadBinary read packet data to Go struct
func ReadBinary(packet Packet) {
	index := packet.Index
	//length := packet.Length
	buffer := packet.PacketBuffer
	packetID := packet.ID

	mMAP := make(map[string]interface{}, 10)

	if packetID != 16 {

	}

	switch packetID {
	case 3: // Set User Slot
		{

			break
		}
	case 13: //Player Update
		{
			mMAP["player_update"] = packets.GetPlayerUpdate(index, buffer)
			break
		}
	case 16: //Player HP
		{

			break
		}
	case 28: //NPC Strike
		{

			break
		}
	}
	json, _ := json.MarshalIndent(mMAP, "", "\t")

	if len(json) > 2 {
		lastPacket = mMAP
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Packets read: %d\n", count)
		fmt.Printf("Last packet read: %s\n", string(json))
	}

}
