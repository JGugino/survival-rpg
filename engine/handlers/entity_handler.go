package handlers

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityType int32
type FacingDirection int32

const (
	PLAYER EntityType = 0

	UP    FacingDirection = 0
	DOWN  FacingDirection = 1
	LEFT  FacingDirection = 2
	RIGHT FacingDirection = 3
)

type EntityData struct {
	Position rl.Vector2
	Width    int
	Height   int
	Speed    float32
	Moving   bool
	Facing   FacingDirection
}

type Entity struct {
	Type     EntityType
	Metadata EntityData
}

type EntityHandler struct {
	entities      []Entity
	playerCreated bool
}

func (entity *EntityHandler) InitHandler() {
	entity.entities = make([]Entity, 0)
	entity.playerCreated = false
}

func (entity *EntityHandler) UpdateEntities() {
	for e := range entity.entities {
		fmt.Println(e)
	}
}

func (entity *EntityHandler) AddEntity(entityType EntityType, metadata EntityData) (created bool, err error) {
	//Check if player already exists and returns error
	if entityType == PLAYER {
		if entity.playerCreated {
			return false, errors.New("player-exists")
		}
	}

	//Add new entity to entities
	entity.entities = append(entity.entities, Entity{Type: entityType, Metadata: metadata})

	//If the add entity was a PLAYER set playerCreated flag to true
	if entityType == PLAYER {
		entity.playerCreated = true
	}

	return true, nil
}
