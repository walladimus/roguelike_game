# Game Modes and Round Structure

## Round Cycle
The game always advances in repeating rounds with a phase pattern:
- **PvE**
- **PvP**
- **Boss**

The overall distribution is **5 : 4 : 1 (PvE : PvP : Boss)**.
- For every 10 total rounds, expect ~5 PvE, ~4 PvP, ~1 Boss.
- Boss fights increase in difficulty each time they appear.
- The final stage of any run is always:
  1. A Boss encounter
  2. Followed immediately by a PvP resolution in that same boss arena

This means late-game PvP happens inside dangerous terrain.

---

## PvE Phase

### Core Goals
- Farm EXP, gold, loot.
- Build loadout (weapons, abilities, upgrades).
- Survive scaling enemies.

### Rewards
- EXP (for levelling and unlocking abilities)
- Gold (for Seller / revives)
- Weapon drops / loot crates

### Enemy Behaviour
- Enemies spawn in the map and pursue players.
- Stronger enemies appear as rounds go on.
- Later rounds: mini-elite enemies can drop rare loot.

### PvE Variants
#### Solo PvE
- No other players in the stage.
- Lower enemy health.
- Enemies do **not** respawn.
- You earn **half EXP** and lower weapon drop rate / rarity.
- Risk: if you fall behind in gear because you played safe solo, you may get destroyed in PvP later.

#### Team PvE
- Multiple players coexist in the same PvE stage.
- Friendly fire OFF (you cannot kill other players here).
- Normal EXP gain.
- Increased rare weapon drop chance.
- Enemies **do respawn** up to a set number of waves per round, and that cap increases over time.
- Enemies have more health to push light cooperation.
- Competitive tension: teammates are also rivals later, so people race to secure kills / loot.
- Time-based: after the timer runs out, that PvE round ends whether all enemies are dead or not.

---

## PvP Phase

### Goal
Fight each other. Get rewarded based on survival, placement, and performance.

There are two PvP sub-modes.

### 1. Standard PvP
- The match runs for a fixed number of cycles.
- Players fight until they’re eliminated.
- NPCs can spawn if <3 players remain to keep pressure up (they hunt nearest player).
- At the end:
  - Winner is whoever lasted the longest overall.
  - Top 3 get better loot rewards.
  - Loot crates may appear during the final ~10% of cycles, with a ~10% spawn chance per cycle.
- Reward scales with how long you stayed alive.

### 2. Competitive PvP
- Higher-pressure mode.
- Players gain points based on placement/ranking.
- Higher rank = better loot.
- Lowest-ranked player each round loses some of their **max** health (base 100).
  - This reduces their future survivability permanently.
- If a player is about to be eliminated early, the game may:
  - Give them a recovery upgrade + ability to rejoin competitive health (comeback mechanic).
  - Allow other players to trigger a “save” ability to keep them alive (social pressure moment).
- The longer you survive, the more EXP you earn.
- Higher chance of loot crates overall in Competitive.
- Matches can last indefinitely (last-person-standing style).

---

## Boss Phase

### Core Loop
- Boss appears every ~5 rounds (following the 5:4:1 ratio).
- Boss difficulty scales upward every time.
- Final endgame is always forced to include a Boss Phase.

### Rules
- Boss may be:
  - Solo encounter per player, OR
  - Shared encounter for the whole group (co-op vs raid boss style)
  - The boss format is defined per boss type.

- Rewards:
  - Loot is split based on contribution (damage dealt).
    - In a 2-player boss fight, top damage dealer gets ~75% of the loot.
    - For larger groups, loot share scales down based on % of damage dealt.

- Wipe penalty:
  - If you die during a boss fight, you might come back using a slot machine / gamble mechanic that costs gold.
    - This is also used in Competitive PvP.
  - If the **whole team wipes**, the team is offered:
    - Option A: Pay a revive cost that scales with team gold pool (e.g. ~60% of combined gold). The cost cannot go down compared to previous revive prices.
    - Option B: End the run (new run required).
  - If majority of players refuse to pay, the run ends.

- Endgame scaling:
  - Bosses get harder over time.
  - Later bosses are meant to be nearly unfair unless you’re well-upgraded.
  - “True Ending” content assumes absurdly scaled bosses at the end, not just “tankier”, but with new mechanics.

---

## Seller

### What it is
- The Seller is a shop that can appear at the start of some rounds.
- Rough appearance ratio: ~2 : 7 : 1 (Seller : None : BossSeller).
  - Most rounds have nothing.
  - Sometimes you get a normal Seller.
  - Rarely you get a Boss Seller with better prices.

### Rules
- Sells weapons, healing, consumables, upgrades in exchange for gold.
- Prices scale with weapon rarity / strength.
- Healing is possible but intentionally expensive.
- Cannot be accessed while in combat.

### Boss Seller
- Higher chance of rolling high-tier loot.
- Sometimes discounted prices after a boss fight.

---

## Enchanter / Upgrader

### What it is
- Secret NPC / station.
- Spawns at a hidden location in the map (different each round).
- Never appears in Boss arenas.

### Purpose
- Enchantments: add modifiers to weapons.
- Upgrades: improve base stats (damage, fire rate, etc.)
- Exploration reward: encourages players to know the map and hunt for high-value upgrade spots.
- Secret ending: finding all Enchanters/Upgraders across a single run unlocks a special ending path.

---

## Difficulty Scaling Summary

- **PvE**
  - Enemy HP, damage, and wave count increase across rounds.
  - In Team PvE, respawns per round also scale up.
- **PvP**
  - Pressure escalates with NPC assists and max-health penalties (Competitive).
- **Boss**
  - Boss stats and mechanics escalate.
  - Revives get more expensive.
  - Loot remains the main way to “keep up.”