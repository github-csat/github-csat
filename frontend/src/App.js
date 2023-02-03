import { useEffect, useState } from "react";
import './App.css';

const Loading = () => (
    <div
        className="App"
        style={{ padding: "100px", fontWeight: "lighter", fontSize: "36px" }}
    >
      Loading...
    </div>
);


function App() {
  const [error, setError] = useState("");
  const [satisfactions, setSatisfactions] = useState([]);
  useEffect(() => {
      const url = "/api/satisfactions";
      (async () => {
          try {
              const response = await fetch(url);
              const json = await response.json();
              setSatisfactions(json);
          } catch (err) {
              console.log("error", err);
              setError(err)
          }
      })();

  }, []);

  return (
      <div className="App">
        {!satisfactions && Loading()}
        <div>
          <div>
            <h1>Title Placeholder</h1>
            <section className="container">
                <ul>
                  {satisfactions.map(Satisfaction)}
                </ul>
            </section>
          </div>
        </div>
      </div>
  );
}

function Satisfaction(satisfaction) {
  return (
          <li key={satisfaction.id}>{satisfaction.id}: {satisfaction.issueUrl}</li>
  )

}


export default App;
