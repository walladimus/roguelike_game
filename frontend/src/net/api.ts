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