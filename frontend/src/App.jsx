import { useState } from 'react'
import { useEffect } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
    const [error, setError] = useState('');
    const [satisfactions, setSatisfactions] = useState([]);

    useEffect(() => {
        const url = '/api/satisfactions';
        (async () => {
            try {
                const response = await fetch(url);
                const json = await response.json();

                setSatisfactions(json || []);
            } catch (err) {
                console.log('error', err);
                setError(err);
            }
        })();
    }, []);

  return (
    <div className="App">
      <div className="card">
        <h1>Hi hi world</h1>
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
          {!satisfactions && Loading()}
          <div>
              <div>
                  {satisfactions.map(Satisfaction)}
              </div>
          </div>
      </div>
    </div>
  )
}

function Satisfaction(satisfaction) {
    return (
        <li key={satisfaction.id}>
            {satisfaction.id}: {satisfaction.issueUrl}
        </li>
    );
}

export default App
