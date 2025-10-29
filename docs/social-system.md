# Social, Friends, Lobbies, and Joining

## Accounts / Identity
- Players create a username.
- Username must be unique.
- Username is how other players add you.
- A lightweight login system links you to:
  - Your stats
  - Your achievements / endings
  - Your friends list
  - Your recent matches

## Friends System
Goals:
- Let people easily play together again.
- Let you build a little network of people you actually like.

Features:
- Send friend request by username.
- Accept / reject requests.
- See current friends.
- See "recent players" from the last few lobbies you were in and quickly add them.

Status states:
- `pending` (request sent, waiting)
- `accepted`
- `blocked` (future protection)

Friends list is mainly:
- party invites
- quick join
- social proof (“this person carried me last night”)

We DO NOT expose real emails or personal info in-game. Username only.

---

## Lobbies

### What is a Lobby
A lobby is the staging area before a run starts.
- Has a host.
- Has a code.
- Has mode settings.

When you click “Play” then “Create”:
- You choose:
  - Story Mode OR Rogue Mode
  - Max players
  - PvE mode (Solo or Team)
  - PvP mode (Standard or Competitive)
  - Round length (if Standard)
  - Difficulty / modifiers (later)

That creates a lobby.

When you click “Join”:
- You can enter a code to join an existing lobby.
- You can also scan for LAN games (local network broadcast).

### Lobby Codes
- Every lobby is given a short code like `F9KQ2`.
- Anyone with that code can join — no full account link needed for LAN testing.
- Once the lobby status becomes `in_progress` (the run actually starts), the code is invalidated and cannot be reused.
- This prevents code re-entry mid-run.

### Ready State
Each lobby member can set `is_ready = true`.
Host can start once all required players are ready.

### LAN Play
The host machine can run the backend server locally on the network, announce itself on LAN.
Other players on the same network can:
- auto-discover host
- or join with the short code

This keeps the “couch co-op / sleepover / school IT room” feel.

---

## Notices & Requests

### Notices
- In the main menu there’s a “Notices & Requests” button.
- “Notices” tab shows:
  - Patch notes
  - Balance tweaks
  - Bug fixes
  - Upcoming features
- Stored and served by the backend.

### Player Requests
- There’s also a “Requests” tab.
- Players can submit suggestions or complaints.
- These get stored for us to read later.
- This gives players a voice and gives us direction for updates.

---

## Buy Me A Coffee
- Button in the main menu opens a browser page for voluntary support (£1–£20).
- Gives unique perks (cosmetic / flex), but does NOT give unfair gameplay advantages.
- No Pay To Win. The core gameplay loop must remain skill + strategy based.

---

## Achievements / Endings View
- “Achievements” menu shows:
  - Achievements you’ve unlocked (full detail)
  - Locked ones as silhouettes / redacted text
- Story Mode endings:
  - You can rewatch unlocked endings here.
  - PvP/Rogue-exclusive unlocks are shown but not replayable as cutscenes.

This makes Story Mode feel like it matters beyond just one clear.