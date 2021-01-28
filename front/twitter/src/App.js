import "./App.css";
import Signup from "./Components/Signup";
import Login from "./Components/Login";
import Home from "./Components/Home";
import Profile from "./Components/Profile";

import * as React from "react";

export const AuthContext = React.createContext();

const initialState = {
  isAuthenticated: false,
  username: null,
  email: null,
  image: null,
  token: null,
};

const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("username", JSON.stringify(action.payload.username));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return {
        isAuthenticated: true,
        username: action.payload.username,
        email: action.payload.email,
        image: action.payload.image,
        token: action.payload.token,
      };
    case "SIGNUP":
      localStorage.setItem("username", JSON.stringify(action.payload.username));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return {
        isAuthenticated: true,
        username: action.payload.username,
        email: action.payload.email,
        image: action.payload.image,
        token: action.payload.token,
      };
    case "LOGOUT":
      localStorage.clear();
      return {
        ...state,
        isAuthenticated: false,
      };
    default:
      return state;
  }
};

function App() {
  const [state, dispatch] = React.useReducer(reducer, initialState);
  return (
    <AuthContext.Provider
      value={{
        state,
        dispatch,
      }}
    >
      {/* <Route path="/" component={Login} /> */}
      <div className="App">{!state.isAuthenticated ? <Login /> : <Home />}</div>
    </AuthContext.Provider>
  );
}

export default App;
