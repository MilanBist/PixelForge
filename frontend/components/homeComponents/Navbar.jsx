import logo from "../../src/assets/logo.png"
import { NavLink } from "react-router-dom";
function Navbar({currentStatus, color = "red"}) {
  if (currentStatus === "Logged In"){
    color = "green";
  }

  console.log("From nav bar",currentStatus);

  return (
    <header className="navbar">
      <div className="brand">
        <div className="brand-icon">
            <img src={logo} alt="Pixel forge" />
        </div>

        <div>
          <h2>PixelForge</h2>
          <span>Generate with us</span>
        </div>
      </div>

      <nav className="nav-links">
        <NavLink to="/">Home</NavLink>
        <NavLink to="/output">Output</NavLink>
        <NavLink to="/history">History</NavLink>
        <NavLink to="/uploadedFiles">Uploaded Files</NavLink>
      </nav>

      <nav className="auth-links">
        <NavLink to="/login">Login</NavLink>
        <p to="/status" style={{color:color}}>{currentStatus}</p>
      </nav>

    </header>
  );
}

export default Navbar;