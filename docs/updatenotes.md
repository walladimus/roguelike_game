# Update Notes - 2025-12-15

## Snapshot
- Backend HTTP + WebSocket services compile and expose the planned surface area (health/version/auth/notices/lobby/save-resume) but only the lobby/resume flows persist data, and gameplay logic remains stubbed.
- Frontend repo has many folders laid out (game/net/ui/data/store) yet almost every file besides a few Svelte menu mockups is empty or syntactically invalid, so there is no runnable client.
- Shared TypeScript contracts exist only as empty placeholders, while the docs folder is already rich with design/balance/roadmap context.

## Backend status
- `backend/cmd/server/main.go` boots the API on `PORT` (default 8081) with the mux router from `backend/internal/httpserver` and an in-memory `storage.Store`; go modules resolve and storage tests (`backend/internal/storage/db_test.go`) cover lobby/save flows.
- `backend/internal/httpserver/router.go` wires middleware, auth/notices services, JSON `save_resume`/`load_resume`, lobby CRUD, and `/ws/lobby` upgrades. Auth currently prefers Postgres when `ROGUE_DB_DSN` is provided but otherwise uses the in-memory store and opaque token strings (no refresh/JWT yet).
- `backend/internal/storage` has a complete in-memory implementation plus Postgres connection helpers, but friends/notices/lobby repository files are stubs waiting for concrete queries; no migration beyond user/password exists. Resume/lobby data never leaves memory across processes.
- `backend/internal/ws` added a multi-lobby hub (`LobbyManager`, `Hub`, `Client`) and snapshot helpers but still only echoes/broadcasts join events; `game` package files (round logic, respawn, balance, etc.) are empty scaffolds, so back-end gameplay simulation is not implemented.

## Frontend status
- Project scaffolds Vite + TS but `frontend/src/main.ts`, `app.d.ts`, most game/system/entity files, and shared UI components are blank files; they simply reserve paths for future code.
- Menu mockups (`frontend/src/ui/menumain.svelte`, `menu*.svelte`) demonstrate the intended flow but have syntax issues (e.g., missing `=` in `on:click` bindings, undefined `<Panel>`/`<Button>` components, inconsistent string interpolation) and rely on not-yet-written state stores.
- Networking helpers are mock-only: `frontend/src/net/api.ts` has malformed syntax (missing braces/try-catch structure) and only logs fake data, `frontend/src/net/ws.ts` hardcodes `ws://localhost:8081/ws/echo` and references an undefined `onMessage` callback. No build-ready entry point glues these together, so the FE cannot talk to the backend yet.
- Data/state/store directories (`frontend/src/data`, `game`, `state`, `store`) contain file shells with no exports, so gameplay systems/UI cannot share any source of truth.

## Shared layer
- Files under `shared/types/*.ts` and `shared/constants/*.ts` are empty placeholders. Nothing imports from `shared/` yet, so there is no enforced contract between the backend Go structs and the frontend TypeScript.

## Documentation
- `README.md` acts as a detailed GDD; `docs/` already covers roadmap, backend API surface, achievements, balance, social systems, and UI wireframes. `docs/changelog.md` records the latest backend work (router, lobbies, auth hashing, WebSocket origins).

## Open issues / next steps
1. Flesh out backend gameplay/state packages (`backend/internal/game/*`) and persist lobby/match data in Postgres so the service does more than lobby bookkeeping.
2. Finish the WebSocket lobby protocol (JOIN/LOBBYSTATE syncing with the HTTP store) and expose typed payloads under `shared/` for the UI.
3. Stand up the Vite entry (`frontend/src/main.ts`) with Svelte root components, fix syntax errors in menu components, and implement real data stores/net clients that hit the backend endpoints instead of mock logs.
4. Populate `shared/constants` and `shared/types` with the contracts implied by the docs (match state, notices, achievements) and reference them from both layers to avoid drift.
