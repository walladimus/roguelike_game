# Achievements & Endings

_Last updated: v0.0.1 design phase_

---

## System Overview
Achievements provide persistent goals across runs, rewarding both skill and exploration.  
They’re stored server-side and tied to your account.  
There are two broad categories:

1. **Gameplay Achievements** — based on combat, loot, and survival.
2. **Story Achievements (Endings)** — unlocked only in Story Mode.

Achievements can award:
- Titles (cosmetic, displayed in lobby)
- Icons or emblems (cosmetic)
- Minor stat bonuses (non-competitive modes only)
- Story cutscene replays (for endings)

---

## Gameplay Achievements

| Code | Name | Condition | Reward |
|:-----|:------|:-----------|:--------|
| FIRST_BLOOD | First Blood | Kill your first enemy. | None |
| SCRAP_COLLECTOR | Scrap Collector | Collect 50 loot items in a single run. | +2% loot rarity chance |
| HARD_EARNED | Hard-Earned | Reach level 10 in one run. | Title: "Veteran" |
| TEAM_PLAYER | Team Player | Heal or revive 3 teammates in a single run. | Emblem: Helping Hand |
| ALONE_WOLF | Lone Wolf | Complete a full PvE round solo without taking damage. | Cosmetic badge |
| KILLSTREAK | Killstreak! | Eliminate 10 enemies within 30 seconds. | +1% movement speed (Rogue only) |
| ARENA_MASTER | Arena Master | Win 10 or more PvP rounds in a row. | Title: "Arena Master" |
| CLUTCH_OR_BUCKET | Clutch or Bucket | Win a PvP round with <10 HP remaining. | Emote: /clutch |
| MERCILESS | Merciless | Eliminate every other player in a Competitive match. | Title: "The Merciless" |
| SHOPPING_SPREE | Shopping Spree | Buy something from every Seller tier in one run. | Cosmetic token |
| ENCHANTED_RUN | Fully Enchanted | Find and use every Enchanter/Upgrader in a run. | Unlocks secret ending path. |
| LAST_MAN | Last Man Laughing | Beat a Boss after every other teammate has died. | Title: "Laatste man die lacht" |

---


## Story Achievements / Endings

Each Story Mode run follows a character arc determined by player actions, choices, and performance.  
There are **10 total endings**, ranging from heroic to tragic.  
Some endings can overlap with secret achievement paths.

| Code | Name | Unlock Condition | Replayable? | Notes |
|:-----|:------|:-----------------|:-------------|:------|
| ENDING_A | **Redemption** | Complete Story Mode defeating the final Boss with *zero team deaths*. | ✅ | Canon “good” ending. |
| ENDING_B | **Sacrifice** | Sacrifice yourself to revive all teammates at the final Boss. | ✅ | Heroic ending; your character is memorialised in the credits. (first 10,000) |
| ENDING_C | **Corruption** | Betray or kill a teammate to survive the final PvP round. | ✅ | Dark ending; unlocks special dialogue in later runs. |
| ENDING_D | **Isolation** | Complete the story entirely solo — no allies, no summons, no revives. | ✅ | Lone-wolf route; unique intro quote next run. |
| ENDING_E | **Descent** | Fail the True Boss encounter after reaching the “Deep Core” stage. | ✅ | Tragic ending; unlocks corrupted visual effects. |
| ENDING_F | **Salvation** | Complete the game by saving every NPC and optional ally. | ✅ | Rare “light” ending; affects UI colour scheme next run. |
| ENDING_G | **Rebellion** | Defeat the final Boss *and* the “watcher” AI that oversees the arena. | ✅ | Meta ending; breaks the 4th wall slightly. |
| ENDING_H | **Illusion** | Accept the Boss’s deal mid-fight (refuse to continue). | ✅ | Non-standard ending; gives special cosmetic title “Dreamer”. |
| ENDING_I | **Oblivion** | Die permanently in the True Boss phase after refusing revival. | ✅ | Hidden nihilistic ending; unlocks “ghost” effect on your player profile. |
| ENDING_TRUE | **True Ascension** | Survive **every round**, defeat **every boss**, and take **no damage for the entire game**. Must be done in Story Mode on Hard. | ✅ | The “True Ending.” Grants golden nameplate, title “Ascendant,” and permanent aura effect. |

---

### 🔐 Notes
- Each ending is stored in the `user_achievements` table as a story-only achievement (`story_only = TRUE`).
- Players can **replay unlocked endings** from the Achievements menu.
- Secret endings (e.g., Rebellion, Illusion) can be hinted through in-game lore or environmen

---

## Hidden / Secret Achievements
The following are intentionally unlisted in-game until earned.

| Code | Name | Hint | Reward |
|:-----|:------|:------|:--------|
| SECRET_DEVIL_DEAL | ??? | “Sometimes gold isn’t enough.” | Hidden cosmetic skin |
| SECRET_UPGRADER_PATH | ??? | “Find them all.” | Unlocks secret ending trigger |
| SECRET_PET_COMPANION | ??? | “A friend in the chaos.” | Adds optional cosmetic companion |
| SECRET_LORE_KEEPER | ??? | “Some books shouldn’t be read twice.” | Unlocks lore codex feature |

---

## Future Expansion Hooks
- Seasonal achievements (timed events)
- Cross-run meta-achievements (“Complete 10 runs total”)
- Competitive ladder achievements (“Top 100 placement”)
- Social achievements (“Play with 20 unique friends”)
