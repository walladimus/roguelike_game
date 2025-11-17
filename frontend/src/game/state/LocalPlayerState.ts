export type PlayerState = {
    hp: number;
    inventory: string[];
    position: { x: number; y: number };
};

let localPlayer: PlayerState = {
    hp: 100,
    inventory: [],
    position: { x: 2, y: 2 }
};

export function setPlayerState(state: PlayerState) {
    localPlayer = state;
}

export function getPlayerState(): PlayerState {
    return localPlayer;
}