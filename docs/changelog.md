# Changelog

## Unreleased

- Added contributor guide `AGENTS.md` outlining project structure, build/test commands, coding style, and PR expectations.
- Backend hardening:
  - WebSocket upgrader now checks `ALLOWED_ORIGINS` (defaults to localhost) and logs connect/join/leave/state events.
  - Lobby system extended with in-memory lobby storage (create/join/get/ready) and REST endpoints: `/lobby/create`, `/lobby/join`, `/lobby/{code}`, `/lobby/{code}/ready`.
  - HTTP router now exposes `/version` and applies request logging + panic recovery middleware.
  - Server boots with in-memory store by default.
- Auth endpoints (`/auth/register`, `/auth/login`) now hash passwords; use in-memory storage by default or Postgres when `ROGUE_DB_DSN` is set (requires migration `0002_add_user_password.sql`); tokens are opaque, not JWT yet.
- WebSocket `LOBBYSTATE` payload now includes lobby members with ready state from the HTTP lobby store; broadcasts stay in sync with HTTP create/join/ready.
- Roadmap doc updated: Phase 1 WebSocket now described as multi-lobby `/ws/lobby?lobby=CODE` with JOIN/LOBBYSTATE messaging; Phase 2 lobby channel notes revised; cleaned Play menu text.
