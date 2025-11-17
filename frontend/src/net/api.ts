export async function getServerVersion(): Promise<string | null>
try 
const res = await fetch("http://localhost:8081/version");
  if (!res.ok) 
    throw new Error("Server responded with an error");
const data = await res.json();
return data.version;
 catch (error) 
  console.error("Failed to fetch server version:", error);
  return null;

export async function sendMove(direction: string) {
  // Simulate backend response
  console.log('Mock sendMove called with direction: ${direction}');

  // Example mock map update
  const mockMap = [
    ["wall", "wall", "wall", "wall", "wall"],
    ["wall", "floor", "floor", "floor", "wall"],
    ["wall", "floor", "floor", "player", "wall"],
    ["wall", "floor", "floor", "floor", "wall"],
    ["wall", "wall", "wall", "wall", "wall"]
  ];

  const mockPlayer = { x: 3, y: 2 };
  const mockTurn = 1;

  return { map: mockMap, player: mockPlayer, turn: mockTurn };
}

export async function fetchGameState() {
  console.log('Mock fetchGameState called');

  const mockMap = [
    ["wall", "wall", "wall", "wall", "wall"],
    ["wall", "floor", "floor", "floor", "wall"],
    ["wall", "floor", "player", "floor", "wall"],
    ["wall", "floor", "floor", "floor", "wall"],
    ["wall", "wall", "wall", "wall", "wall"]
  ];

  const mockPlayer = { x: 2, y: 2 };
  const mockTurn = 0;

  return { map: mockMap, player: mockPlayer, turn: mockTurn };
}