package packets

import (
	"encoding/binary"
	"encoding/json"
	"math"
)

// PlayerUpdate the Player update packet
type PlayerUpdate struct {
	PlayerID int                 `json:"player_id"`
	Control  PlayerUpdateControl `json:"control"`
	Pulley   PlayerUpdatePulley  `json:"pulley"`
}

// PlayerUpdateControl the player control direction
type PlayerUpdateControl struct {
	UP        bool `json:"up"`
	DOWN      bool `json:"down"`
	LEFT      bool `json:"left"`
	RIGHT     bool `json:"right"`
	UseItem   bool `json:"use_item"`
	Direction int  `json:"direction"`
}

// PlayerUpdatePulley the player especial events
type PlayerUpdatePulley struct {
	Enabled             bool `json:"enabled"`
	Direction           int  `json:"direction"`
	UpdateVelocity      bool `json:"update_velocity"`
	VortexStealthActive bool `json:"vortex_stealth_active"`
	GravityDirection    int  `json:"gravity_direction"`
	ShieldRaised        bool `json:"shield_raised"`
}

// GetPlayerUpdate read the buffer packet and return the Map struct
func GetPlayerUpdate(index int, buffer []byte) PlayerUpdate {
	playerUpdate := make(map[string]interface{}, 10)

	var mByte byte

	i := index + 0

	mByte = buffer[i]

	playerID := int(mByte)
	playerUpdate["player_id"] = playerID

	i++

	mByte = buffer[i]

	controlUp := mByte >> 0
	controlDown := mByte >> 1
	controlLeft := mByte >> 2
	controlRight := mByte >> 3
	controlJump := mByte >> 4
	controlUseItem := mByte >> 5
	controlDirection := mByte >> 6

	playerUpdate["control"] = map[string]interface{}{
		"up":        controlUp == 1,
		"down":      controlDown == 1,
		"left":      controlLeft == 1,
		"right":     controlRight == 1,
		"jump":      controlJump == 1,
		"use_item":  controlUseItem == 1,
		"direction": controlDirection,
	}

	i++

	mByte = buffer[i]

	pulleyEnabled := mByte >> 0
	pulleyDirection := mByte >> 1
	pulleyUpdateVelocity := mByte >> 2
	pulleyVortexStealthActive := mByte >> 3
	pulleyGravityDirection := mByte >> 4
	pulleyShieldRaised := mByte >> 5

	playerUpdate["pulley"] = map[string]interface{}{
		"enabled":               pulleyEnabled == 1,
		"direction":             int(pulleyDirection),
		"update_velocity":       int(pulleyUpdateVelocity),
		"vortex_stealth_active": pulleyVortexStealthActive == 1,
		"gravity_direction":     int(pulleyGravityDirection),
		"shield_raised":         pulleyShieldRaised == 1,
	}

	i++

	mByte = buffer[i]

	miscHoveringUp := mByte >> 0
	miscVoidVaultEnabled := mByte >> 1
	miscSitting := mByte >> 2
	miscDownedDD2Event := mByte >> 3
	miscIsPettingAnimal := mByte >> 4
	miscIsPettingSmallAnimal := mByte >> 5
	miscUsedPotionOfReturn := mByte >> 6
	miscHoveringDown := mByte >> 7

	playerUpdate["misc"] = map[string]interface{}{
		"hovering_up":             miscHoveringUp == 1,
		"void_vault_enabled":      miscVoidVaultEnabled == 1,
		"sitting":                 miscSitting == 1,
		"downed_dd2_event":        miscDownedDD2Event == 1,
		"is_petting_animal":       miscIsPettingAnimal == 1,
		"is_petting_small_animal": miscIsPettingSmallAnimal == 1,
		"used_potion_of_return":   miscUsedPotionOfReturn == 1,
		"hovering_down":           miscHoveringDown == 1,
	}

	i++

	mByte = buffer[i]

	isSleepingInfo := mByte >> 0

	playerUpdate["is_sleeping"] = isSleepingInfo == 1

	i++

	mByte = buffer[i]

	selectedItem := mByte >> 0

	playerUpdate["selected_item"] = int(selectedItem)

	i++

	var mBytes = buffer[i : i+4]

	positionX := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["position_x"] = positionX

	i += 4

	mBytes = buffer[i : i+4]

	positionY := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["position_y"] = positionY

	i += 4

	mBytes = buffer[i : i+4]

	velocityX := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["velocity_x"] = velocityX

	i += 4

	mBytes = buffer[i : i+4]

	velocityY := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["velocity_y"] = velocityY

	i += 4

	mBytes = buffer[i : i+4]

	originalPositionX := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["original_position_x"] = originalPositionX

	i += 4

	mBytes = buffer[i : i+4]

	originalPositionY := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["original_position_y"] = originalPositionY

	i += 4

	mBytes = buffer[i : i+4]

	homePositionX := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["home_position_x"] = homePositionX

	i += 4

	mBytes = buffer[i : i+4]

	homePositionY := math.Float32frombits(binary.LittleEndian.Uint32(mBytes))

	playerUpdate["home_position_y"] = homePositionY

	var mPlayerUpdate PlayerUpdate

	data, _ := json.Marshal(playerUpdate)
	json.Unmarshal(data, &mPlayerUpdate)

	return mPlayerUpdate
}
