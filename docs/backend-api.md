# Backend API Cheatsheet

Quick reference for the early phases described in the roadmap (Phase 1 lobby skeleton + health/version).

## HTTP
- `GET /health` → `{ "ok": true }`
- `GET /version` → `{ "version": "0.0.1-dev" }`
- `POST /lobby/create`  
  Body: `{ "username": "HostName" }`  
  Response: `{ "code": "ABC123", "members": [ { "username": "HostName", "ready": false } ] }`
- `POST /lobby/join`  
  Body: `{ "code": "ABC123", "username": "Player2" }`  
  Response: same shape as create, with added member.
- `POST /lobby/{code}/ready`  
  Body: `{ "username": "Player2", "ready": true }`  
  Response: lobby with updated ready flags.
- `GET /lobby/{code}`  
  Response: lobby snapshot `{ code, members }`.
- Debug: `GET /debug/lobbies` lists active WS hubs and player names (for local dev).

## WebSocket
- URL: `/ws/lobby?lobby=ABC123` (defaults to `default` if missing).
- On connect, send `{"type":"JOIN","data":{"username":"Player2"}}`.
- Broadcasts you will receive:
  - `LOBBYSTATE`: `{"type":"LOBBYSTATE","data":{"lobbyCode":"ABC123","members":[{"username":"HostName","ready":false},{"username":"Player2","ready":true}]}}`
  - Echoed chat or other messages sent by clients (CHAT, TURNSTATE placeholders).
- Lifecycle:
  - Server logs connect/disconnect.
  - When HTTP lobby create/join/ready succeeds, server pushes a fresh `LOBBYSTATE` to that lobby’s hub so UI stays in sync.

## Notes / TODO for Phase 2
- Auth endpoints (`/auth/register`, `/auth/login`) are stubbed: they accept username/password and return a placeholder token; replace with a real store when accounts land.
- Lobby storage is in-memory; replace with DB-backed store when migrations are ready.

## Current Implementation Gaps / Extras
- `POST /auth/register` and `POST /auth/login` are wired with hashed storage and random opaque tokens. Persistence is in-memory by default; if `ROGUE_DB_DSN` is set, Postgres is used (requires migration `0002_add_user_password.sql`). No JWT yet.
- WebSocket lobby snapshots (`LOBBYSTATE`) broadcast lobby membership (username + ready) from the HTTP store when create/join/ready succeed (delivered to connected clients in that lobby if any).
- Additional dev endpoints exist that are not part of this cheatsheet: `POST /api/save_resume` stores an opaque JSON blob and returns an `id`; `GET /api/load_resume?id=` returns that blob if present.
