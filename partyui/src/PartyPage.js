import logo from './logo.svg';
import './App.css';
import { Link } from "react-router-dom";

function PartyPage() {
  return (
    <div className="App">
      <Link to="/">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </header>
      </Link>
    </div>
  );
}

export default PartyPage;
