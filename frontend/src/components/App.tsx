import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

function App() {
  const [count, setCount] = useState(0);
  const [error, setError] = useState("");
  const [satisfactions, setSatisfactions] = useState([]);

  //todo redirect to login if not logged in

  const isAdmin = true; //todo set state from API/session

  useEffect(() => {
    const url = "/api/satisfactions";
    (async () => {
      try {
        const response = await fetch(url);
        const json = await response.json();

        setSatisfactions(json || []);
      } catch (err) {
        console.log("error", err);
        setError(err);
      }
    })();
  }, []);

  return (
    <div className="App">
      <div className="card">
        <h1>Hi hi world</h1>
        {error && <p>There was an error: {error}</p>}
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        {!satisfactions && Loading()}
        <div>
          <div>{satisfactions.map(Satisfaction)}</div>
        </div>
      </div>
      <div>
        You probably want to <a href={"/api/oauth/redirect"}>Log In</a>
      </div>
      {isAdmin && (
        <div>
          Congrats on being an admin! You can view the results.
          <Link to={"/results"}>Results</Link>
        </div>
      )}
    </div>
  );
}

function Satisfaction(satisfaction) {
  return (
    <li key={satisfaction.id}>
      {satisfaction.id}: {satisfaction.issueUrl}
    </li>
  );
}

export default App;
