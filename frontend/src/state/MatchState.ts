export type GameState = {
  map: string[][];
  player: { x: number; y: number };
  turn: number;
};

let currentState: GameState = {
  map: [],
  player: { x: 0, y: 0 },
  turn: 0
};

export function setGameState(state: GameState) {
  currentState = state;
}

export function getGameState(): GameState {
  return currentState;
}