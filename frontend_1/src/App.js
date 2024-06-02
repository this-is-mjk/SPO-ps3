import Input from "@mui/material/Input";
import Button from "@mui/material/Button";
import { useState } from "react";
import "./App.css";
import axios from "axios";

function App() {
  const [formData, setFormData] = useState({
    UserId: "",
    Password: "",
  });
  const [action, setAction] = useState("Login");
  const [url, setUrl] = useState("http://localhost:8080/login");
  // const [response, setResponse] = useState(null);

  const handelChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };
  const changeAction = () => {
    if (action === "Login") {
      setAction("SignUp");
      setUrl("http://localhost:8080/signup");
    } else {
      setAction("Login");
      setUrl("http://localhost:8080/login");
    }
  };
  const request = () => {
    axios
      .post(url, {
        UserId: formData.UserId,
        Password: formData.Password,
      })
      .then((response) => {
        console.log(response);
        window.alert(response.data);
      })
      .catch((error) => {
        console.log(error);
        window.alert(error.message);
      });
  };
  return (
    <div className="login-signup-page">
      <div className="dataForm">
        <h2>{action}</h2>
        <Input
          className="inputField-btn"
          type="text"
          name="UserId"
          value={formData.userId}
          onChange={handelChange}
          placeholder="User ID"
        />
        <Input
          className="inputField-btn"
          type="password"
          name="Password"
          value={formData.password}
          onChange={handelChange}
          placeholder="Password"
        />
        <Button
          variant="contained"
          color="success"
          onClick={request}
          style={{ margin: "10px", width: "70%" }}
        >
          {action}
        </Button>
        <Button
          className="inputField-btn"
          onClick={changeAction}
          style={{ margin: "10px", width: "70%" }}
        >
          {action === "Login" ? "Sign Up" : "Login"}
        </Button>
      </div>
    </div>
  );
}
export default App;
