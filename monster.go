package main

import (
	"math/rand"
)

type Monster struct {
	*Player
}

type MonsterModifier int

const (
	MonsterModifierNormal MonsterModifier = 1
	MonsterModifierThough                 = 2
	MonsterModifierBoss                   = 5
)

func (mm MonsterModifier) String() string {
	switch mm {

	case MonsterModifierNormal:
		return "Normal"
	case MonsterModifierThough:
		return "Though"
	case MonsterModifierBoss:
		return "Boss"
	}

	return "???"
}

var MonsterTypes = []*MonsterType{
	&MonsterType{
		Name:     "Blob",
		LvlStart: 0,
		LvlEnd:   5,
	},
	&MonsterType{
		Name:     "Bushes",
		LvlStart: 0,
		LvlEnd:   5,
	},
	&MonsterType{
		Name:     "Bird",
		LvlStart: 5,
		LvlEnd:   10,
	},
	&MonsterType{
		Name:     "Cat",
		LvlStart: 5,
		LvlEnd:   10,
	},
	&MonsterType{
		Name:     "Dog",
		LvlStart: 10,
		LvlEnd:   15,
	},
	&MonsterType{
		Name:     "Deer",
		LvlStart: 10,
		LvlEnd:   15,
	},
}

type MonsterType struct {
	Name     string
	LvlStart int
	LvlEnd   int
}

func GetMonster(level int) *Monster {
	monsterType := RandomMonsterType(level)
	if monsterType == nil {
		return nil
	}

	modifier := RandomMonsterModifier()

	monster := &Monster{
		Player: &Player{
			Name:  modifier.String() + " " + monsterType.Name,
			XP:    GetXPForLevel((level - 2) + int(modifier)),
			Money: 2 + int(modifier)*2,
		},
	}

	for monster.AvailableAttributePoints() > 0 {
		rng := rand.Float32()

		var attribute AttributeType
		if rng < 0.33 {
			attribute = AttributeStrength
		} else if rng < 0.66 {
			attribute = AttributeAgility
		} else {
			attribute = AttributeStamina
		}
		monster.Attributes.Modify(attribute, 1)
	}

	return monster
}

func RandomMonsterType(level int) *MonsterType {
	pool := make([]*MonsterType, 0)

	for _, mt := range MonsterTypes {
		if level >= mt.LvlStart && level < mt.LvlEnd {
			pool = append(pool, mt)
		}
	}
	if len(pool) < 1 {
		return nil
	}

	return pool[rand.Intn(len(pool))]
}

func RandomMonsterModifier() MonsterModifier {
	num := rand.Float32()

	if num < 0.6 {
		return MonsterModifierNormal
	} else if num < 0.9 {
		return MonsterModifierThough
	} else {
		return MonsterModifierBoss
	}
	return MonsterModifierNormal
}
