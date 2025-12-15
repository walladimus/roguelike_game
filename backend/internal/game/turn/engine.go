package turn

import (
	"sync"
	"time"
)

var (
	mu   sync.Mutex
	demo *Match
)

func Init() *Match {
	mu.Lock()
	defer mu.Unlock()

	if demo != nil {
		return demo
	}

	now := time.Now().UTC()
	demo = &Match{
		ID:          "DEMO",
		Round:       1,
		Phase:       PhasePVE,
		ActiveIndex: 0,
		Players: []Player{
			{ID: "P1", Name: "Zain", HP: 100, Alive: true},
			{ID: "P2", Name: "Simon", HP: 100, Alive: true},
		},
		CreatedAt: now,
		UpdatedAt: now,
		pveLeft:   5,
	}
	return demo
}

func Advance() *Match {
	m := Init()

	mu.Lock()
	defer mu.Unlock()

	if n := len(m.Players); n > 0 {
		m.ActiveIndex = (m.ActiveIndex + 1) % n
	}

	switch m.Phase {
	case PhasePVE:
		m.pveLeft--
		if m.pveLeft <= 0 {
			m.Phase = PhasePVP
			m.pvpLeft = 4
		}
	case PhasePVP:
		m.pvpLeft--
		if m.pvpLeft <= 0 {
			m.Phase = PhaseBOSS
			m.bossLeft = 1
		}
	case PhaseBOSS:
		m.bossLeft--
		if m.bossLeft <= 0 {
			m.Round++
			m.Phase = PhasePVE
			m.pveLeft = 5
		}
	}

	m.UpdatedAt = time.Now().UTC()
	return m
}

func Snapshot() Match {
	_ = Init()

	mu.Lock()
	defer mu.Unlock()

	cp := *demo
	cp.Players = append([]Player(nil), demo.Players...)
	return cp
}
