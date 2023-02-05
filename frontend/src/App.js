import "./App.css";
import React, { createContext, useReducer } from 'react';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/Home";
import Login from "./components/Login";
import Satisfactions from "./components/Satisfactions";
import { initialState, reducer } from "./store";


export const AuthContext = createContext();

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);

  return (
      <AuthContext.Provider
          value={{
            state,
            dispatch
          }}
      >
        <Router>
          <Routes>
            <Route path="/" element={<Home/>}/>
              <Route path="/auth" element={<Login/>}/>
              <Route path="/auth/callback" element={<Login/>}/>
              <Route path="/satisfactions" element={<Satisfactions/>}/>
          </Routes>
        </Router>
      </AuthContext.Provider>
  );
}

export default App;