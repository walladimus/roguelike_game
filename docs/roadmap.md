# Roadmap / Build Order

This is the order we will build the game so that it's playable as early as possible and scales without rewrite.

---

## Phase 0: Repo + Docs (YOU ARE HERE)
- Project folders created:
  - `frontend/`, `backend/`, `shared/`, `docs/`
- High-level design docs written:
  - `vision.md`
  - `modes.md`
  - `social-system.md`
  - `roadmap.md`
- Agreement on tech stack:
  - Frontend: browser client (TypeScript, Svelte or similar)
  - Backend: Go HTTP + WebSocket
  - DB: Postgres (or SQLite for dev)
- Shared types folder planned
- Visual treatment: commit to pre-rendered assets and a nostalgic 2000-2004 presentation so art/systems decisions align early.

Goal: Alignment. Everyone knows what game we’re making.

---

## Phase 1: Bare Minimum Running Build
Backend:
- `/health` (returns `{ ok: true }`)
- `/version` (returns build version)
- WebSocket `/ws/lobby?lobby=CODE` (multi-lobby hub). Clients send `{"type":"JOIN","data":{"username":"..."}}` on connect; server rebroadcasts chat and `LOBBYSTATE { lobbyCode, members: [{ username, ready }] }` to everyone in that lobby.

Frontend:
- Main Menu UI:
  - Game title
  - Buttons: Play / Achievements / Friends / Settings / Notices / Coffee
  - Play shows Create / Join menu (non-functional)
- Connect to `/health` just to prove backend reachability

Goal: You can open the game in a browser, see the menu, and know the server is alive.

---

## Phase 2: Lobbies & Friends Skeleton
Database:
- Add `users`, `friends`, `lobbies`, `lobby_members`, `player_stats`, `achievements`, `user_achievements`, `notices`, `player_requests`.

Backend:
- `POST /auth/register`, `POST /auth/login` (basic username+password; hashed, in-memory by default, Postgres when configured)
- `POST /lobby/create`
- `POST /lobby/join`
- `GET /lobby/:code` → returns current members + ready state
- WebSocket lobby channel:
  - join with `?lobby=CODE`, send JOIN with username; server broadcasts `LOBBYSTATE { lobbyCode, members: [{ username, ready }] }` on join/leave/ready plus chat notices when players join/leave.

Frontend:
- “Create Game” form actually calls backend
- “Join Game” form actually calls backend
- Lobby screen shows:
  - Players in lobby
  - Their ready status
  - Start button for host (disabled if not all ready)

Goal: Multiple clients can sit in the same lobby and see each other update live.

---

## Phase 3: Round/Phase Engine (No Combat Yet)
Backend:
- Match state machine:
  - round number
  - current phase: PvE / PvP / Boss
  - timer until next phase
- Broadcasts this over WebSocket to all lobby members as the “match loop”

Frontend:
- Simple HUD / overlay that shows:
  - “Round 3”
  - “Phase: PvP - Competitive”
  - Countdown timer
- This replaces boredom; players feel like “the game is running,” even though we’re not simulating combat yet.

Goal: We get the backbone of PvE→PvP→Boss flow and 5:4:1 logic working.

---

## Phase 4: Basic Combat Sandbox
Frontend:
- Implement a simple 2D arena view.
- Player can move a rectangle around using the pre-rendered nostalgic art kit (keep mechanics simple while the renderer sells the look).
- “Enemies” are dumb AI squares that move toward you.
- HP and damage numbers exist locally.

Backend:
- For now, trust client for position/hits (we'll fix cheating later).
- Sync player HP / deaths / XP.

Goal: You can actually “play” a primitive PvE round.

---

## Phase 5: Scaling / Economy / Boss
- Add Seller
- Add revive cost logic
- Add Boss damage contribution and loot split
- Add max-health penalty for Competitive PvP
- Add “True Ending” flag to achievements

Goal: Game loop is now recognisably the thing we pitched.

---

## Phase 6: Polish
- Map variety and secret Enchanter/Upgrader spawn logic
- LAN auto-discovery
- Cosmetics + coffee support perks
- Accessibility / performance tuning

---

## MVP Definition
We hit MVP when ALL of these are true:
- You can create a lobby
- Your friend can join with a code
- The phase cycle runs (PvE → PvP → Boss and repeats)
- You both see the same round/phase state
- You can move and take damage in PvE
- There’s a basic reward loop (XP, drops)
- Achievements can unlock and be viewed in the Achievements menu

