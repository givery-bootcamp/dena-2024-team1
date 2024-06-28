import { useEffect, useState } from "react";

export function Event() {
  const [events, setEvents] = useState<{ data: string }[]>([]);

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:9000/events");

    eventSource.onmessage = function(event) {
      setEvents(prevEvents => [...prevEvents, event.data]);
    };

    eventSource.onerror = function(event) {
      console.error("EventSource failed:", event);
      eventSource.close();
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return (
    <div>
      <header>
        <h1>Server-Sent Events Example</h1>
        <div>
          {events.map((event, index) => (
            <div key={index}>{JSON.stringify(event)}</div>
          ))}
        </div>
      </header>
    </div>
  );
}
