# Game Vision

## Working Title
Roguelike Arena (TBD — name to be finalised)

## High Concept
A fast, replayable, browser-based roguelike where players fight through escalating PvE waves, brutal PvP rounds, and high-stakes boss encounters in a repeating cycle. You can play solo or with friends. You gain loot, abilities, upgrades, and score. You unlock achievements and endings. You either survive — or you get wiped and start again.

Runs are short, intense, and social. Every run feels competitive and matters.

## Core Loop
1. **PvE Phase**
   - Fight AI enemies to gain EXP, weapons, gold, and abilities.
   - Enemies scale in difficulty across rounds.
   - Team or Solo variant.

2. **PvP Phase**
   - Fight other players.
   - Compete for position, survival, and loot rewards.
   - Two PvP modes:
     - Standard (round-limited survival / top 3 rewarded)
     - Competitive (ladder-style pressure with max health penalties)

3. **Boss Phase**
   - High-difficulty boss encounter.
   - Sometimes cooperative, sometimes individual.
   - High reward, high wipe risk.
   - Death here might force buybacks or end runs.

The phases repeat in a ratio of **5 : 4 : 1 (PvE : PvP : Boss)**.  
The final stage in any run always ends with a Boss encounter, followed by a PvP resolution in that same arena.

## Win / Lose State
- You “win” a run by:
  - Surviving to the end of the defined Story mode path, and defeating the final boss.
  - Outlasting everyone else in Rogue Competitive.
  - Achieving the True Ending condition.
- You “lose” a run when:
  - You die and cannot be revived (PvE/PvP).
  - Your team wipes on a boss and refuses (or cannot afford) the gold cost to come back.
  - You are eliminated in Competitive and do not recover.

## Player Fantasy
- “I’m getting stronger every round.”
- “I can clutch carry a team or betray them later.”
- “If I survive one more fight I get insane loot.”
- “We barely beat that boss and I was the MVP.”

## Player Motivation
- **Short-term:** better loot, more abilities, surviving the next round.
- **Mid-term:** achievements, progression of stats, unlock cosmetic bragging rights.
- **Long-term:** story endings, rare titles, reputation in Competitive.

## Design Pillars
1. **Replayable pressure**
   - Every round matters. Health and resources carry forward.
   - You always feel one mistake from disaster.

2. **Social tension**
   - Team in PvE, enemies in PvP.
   - You might heal someone now who'll try to kill you later.

3. **Fair clarity**
   - Drops, revive prices, boss targeting, etc. should feel harsh but explainable.
   - The player should always understand why they died or lost HP/loot.

4. **Low friction**
   - Browser-based.
   - Runs on weak hardware.
   - Easy to invite friends via code or LAN.

## Technical Goals
- Browser-based client.
- Lightweight rendering (2D for now).
- Real-time sync via WebSocket.
- Can run locally (LAN) and online.
- Data model supports accounts, lobbies, achievements, run history.

## Modes
Two top-level modes:
- **Story Mode**
  - Fixed-length run.
  - Scripted bosses, character arcs, defined endings.
  - Unlocks story achievements and lets you rewatch endings.
- **Rogue Mode**
  - Sandbox / repeatable.
  - Can be Standard (round-limited) or Competitive (last one alive).
  - Higher replay value.
  - Enables “true” endurance achievements and brag titles.

## Endgame / True Ending
- There is a “True Ending” that is intentionally extremely hard.
- Bosses keep scaling late-run.
- Final arena is forced Boss → PvP showdown.
- “True Ending” gives a unique, very rare achievement and title.
