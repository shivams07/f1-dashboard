import React, { useEffect, useState } from "react";

function App() {
  const [telemetry, setTelemetry] = useState<any>(null);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setTelemetry(data);
    };

    return () => ws.close();
  }, []);

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Live Telemetry</h1>
      {telemetry ? (
        <div>
          <p><strong>Car:</strong> {telemetry.carId}</p>
          <p><strong>Speed:</strong> {telemetry.speed} km/h</p>
          <p><strong>Position:</strong> X: {telemetry.x}, Y: {telemetry.y}</p>
        </div>
      ) : (
        <p>Waiting for data...</p>
      )}
    </div>
  );
}

export default App;