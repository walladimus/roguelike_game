import { getGameState, setGameState } from '$lib/state/MatchState';

export function tryMove(direction: 'up' | 'down' | 'left' | 'right') {
  const state = getGameState();
  const { map, player, turn } = state;

  const dx = direction === 'left' ? -1 : direction === 'right' ? 1 : 0;
  const dy = direction === 'up' ? -1 : direction === 'down' ? 1 : 0;

  const newX = player.x + dx;
  const newY = player.y + dy;

  if (newY < 0 || newY >= map.length || newX < 0 || newX >= map[0].length) {
    return false;
  }

  const targetTile = map[newY][newX];

  if (targetTile === 'wall') {
    return false;
  }

  const newMap = map.map((row, y) =>
    row.map((cell, x) => {
      if (x === player.x && y === player.y) return 'floor';
      if (x === newX && y === newY) return 'player';
      return cell;
    })
  );

  setGameState({
    map: newMap,
    player: { x: newX, y: newY },
    turn: turn + 1
  });

  return true;
}