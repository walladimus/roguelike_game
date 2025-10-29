# Balance Notes and Scaling Philosophy

_Last updated: v0.0.1 design phase_

---

## Core Balancing Goals

1. **Fair escalation** — The player should always feel like they can win *if* they play smart, even though it gets harder every round.  
2. **Positive tension** — Resource scarcity and time pressure create meaningful choices (heal, upgrade, or save gold).  
3. **Consistency across hardware** — All balancing must work on low-performance devices; no physics-heavy scaling or visual gimmicks that affect gameplay.  
4. **Replay variability** — Loot and encounter variation ensure that no two runs feel identical, but difficulty should rise in predictable bands.

---

## Difficulty Curves

### PvE Difficulty Curve
| Round | Enemy HP Multiplier | Enemy Damage Multiplier | Elite Chance | Notes |
|:------|:--------------------|:------------------------|:-------------|:------|
| 1–2   | ×1.0                | ×1.0                    | 0%           | Tutorial-level enemies. |
| 3–5   | ×1.2                | ×1.1                    | 10%          | Light scaling begins. |
| 6–8   | ×1.5                | ×1.25                   | 15%          | Pressure phase; Seller appears once here. |
| 9–12  | ×2.0                | ×1.4                    | 20%          | Late-mid ramp; Elite drops include rare loot. |
| 13+   | ×2.5+               | ×1.6+                   | 30%+         | Designed to kill complacent players. |

*Every fifth round is eligible to trigger a Boss phase.*

### PvE Enemy Respawn Rules (Team mode)
- Base respawns per round: 2.
- +1 respawn per 5 rounds (max 6 respawns).
- Each respawn wave increases enemy HP by +5%.

### PvP Damage Scaling
| Player Count | Base Max HP | Damage Taken Multiplier | Kill Reward (Gold) | Notes |
|:--------------|:------------|:------------------------|:-------------------|:------|
| 2–3 players   | 100 HP      | ×1.0                    | +20% base          | Standard duel / skirmish. |
| 4–6 players   | 100 HP      | ×1.2                    | +10% base          | Faster deaths; more chaos. |
| 7+ players    | 100 HP      | ×1.3                    | +5% base           | Encourages quick fights. |

Competitive PvP:
- Lowest player each round loses **10% of total max HP** permanently.
- Minimum max HP = 40.

Standard PvP:
- Top 3 gain tiered rewards:
  - 1st: +25% EXP, +1 Lootcrate chance
  - 2nd: +10% EXP
  - 3rd: +5% EXP

---

## Boss Difficulty Scaling

| Boss # | HP Multiplier | Damage Multiplier | Loot Quality | Notes |
|:-------|:---------------|:------------------|:--------------|:------|
| 1      | ×1.0           | ×1.0              | Common–Rare   | Introductory boss. |
| 2      | ×1.3           | ×1.2              | Rare–Epic     | First test of coordination. |
| 3      | ×1.7           | ×1.4              | Rare–Epic+    | Survivors only; revives costly. |
| 4      | ×2.0           | ×1.6              | Epic–Legendary| Near-end Boss. |
| 5+     | ×2.5+          | ×2.0+             | Legendary+    | “True Ending” bosses; possibly multi-stage. |

**Boss Loot Distribution Rule:**  
- Damage share determines loot percentage.  
- Minimum guarantee: all surviving players get at least one drop.  
- Example (4-player co-op):
  - Player A: 40% damage → 40% loot pool  
  - Player B: 30% → 30%  
  - Player C: 20% → 20%  
  - Player D: 10% → 10%

**Revive / Buyback Costs**
| Attempt # | Base Cost (% of total gold) | Additional Effect |
|:-----------|:----------------------------|:------------------|
| 1          | 40%                         | Slot machine 75% revive chance |
| 2          | 60%                         | Slot machine 60% revive chance |
| 3          | 80%                         | Slot machine 45% revive chance |
| 4+         | 100%+                       | Team wipe if failed |

If the **whole team wipes**, revival requires majority approval and pooled gold.

---

## Economy Balance

### Base Gold
- Base gold per PvE round: **100 ± 25% RNG**
- Bonus gold from kills:  
  - Normal enemy: +10  
  - Elite: +30  
  - Boss minion: +50  

### Seller Pricing
| Item Type | Base Cost | Notes |
|:-----------|:-----------|:------|
| Healing Potion | 150 | Restores 50 HP |
| Armor Upgrade | 300 | +10 max HP (one use per run) |
| Weapon (Rare) | 400–700 | Scales with drop rarity |
| Weapon (Epic) | 800–1000 | Often cheaper at Boss Seller |
| Revive Token | Variable | 60% team gold, round-based scaling |

Healing should always feel expensive but achievable; forcing risk/reward choices.

---

## XP / Leveling Curve

| Level | Total XP Needed | Unlock |
|:------|:----------------|:-------|
| 1 → 2 | 100 | +5% base damage |
| 2 → 3 | 300 | +1 ability slot |
| 3 → 4 | 600 | +10% HP regen speed |
| 4 → 5 | 1000 | +5% movement speed |
| 5 → 6 | 1500 | +1 passive ability |
| 6 → 7 | 2100 | +5% loot rarity chance |
| 7 → 8 | 2800 | +1 ultimate ability |
| 8 → 9 | 3600 | +10% XP gain |
| 9 → 10 | 4500 | Unlocks cosmetic title “Veteran” |

Each run can soft-cap at level 10; additional XP contributes to leaderboard totals.

---

## Loot Rarity Distribution

| Rarity | Drop Chance (PvE) | Drop Chance (Boss) | Base Multiplier | Color |
|:--------|:------------------|:-------------------|:----------------|:------|
| Common | 60% | 25% | ×1.0 | Gray |
| Uncommon | 25% | 25% | ×1.15 | Green |
| Rare | 10% | 25% | ×1.3 | Blue |
| Epic | 4% | 20% | ×1.5 | Purple |
| Legendary | 1% | 5% | ×1.8 | Gold |

Team mode adds a +5% bonus chance to roll **Rare+** loot.

---

## Design Philosophy Summary
- **Never RNG-lock progress.** Every round should give *some* gain.  
- **Keep risk meaningful.** Healing and revival must always cost more than players want to pay.  
- **Solo = hard but fair.** It’s not a punishment; just fewer rewards.  
- **Team = rich but risky.** More players = more chaos = higher variance.  
- **Boss = gatekeeper.** Surviving a boss fight should feel like a reward, not just a checkpoint.  