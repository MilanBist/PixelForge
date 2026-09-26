import axios from "axios";
import "../styles/Login.css"
import { useState } from "react";
import { Link, Navigate } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import apiClient from "../utils/Base";



// check for the email
const checkEmail = (email)=>{
    let newEmail = String(email).trim();
        // use the regex to verify the email
    const re  = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    const result = re.test(email);

    if (result !== true){
        return false;
    }
    return true;

}

// check for the password
const checkPassword = (password)=>{
    const re = /^[a-zA-Z0-9!@#$%^&*]{6,16}$/;
    const result = re.test(password);

    if (result !== true){
        return false;
    }
    return true;
}

export default function Login({setCurrentStatus}){
    const navigate = useNavigate();
    // make the states for all of the given things
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const UpdatePassword = (evt)=>{
        setPassword(evt.target.value);
    }
    const UpdateEmail = (evt)=>{
        setEmail(evt.target.value);
    }    

    const handleLoginSubmit = (evt)=>{
        evt.preventDefault();
       

        console.log("Button is being clicked.")
        // check for email and password
        if (String(email).length ==0){
            alert("Enter the email.");
            return
        }
        let cmail = checkEmail(email);
        if (String(password).length === 0){
            alert("Enter the password.");
            return;
        }

        let cpass = checkPassword(password);

        if (cmail === false){
            alert("Wrong email.");
            return;
        }

        if (cpass === false){
            alert("Wrong password.");
            return;
        }


    const bodyMap = {
        email: email,
        password: password
    };
''
    // to be sent
    console.log("For the login the body to sent: ", bodyMap)


    // if both are correct then call the handler for the login
    apiClient.post("/login",bodyMap).then((resp)=>{
            localStorage.setItem("accessToken", resp.data["data"]["access"]);
            localStorage.setItem("refreshToken", resp.data["data"]["refresh"]);
            setCurrentStatus("Logged In");
            navigate("/");
        }).catch((err) => {
            setError(err);
            const responseStatus = err.response.status;
            switch (responseStatus){
                case 400:
                    alert("Wrong sending method.");
                case 401:
                    alert("Login credentials.");
                case 404:
                    alert("User doesn't exist. Please register.");
                case 500:
                    alert("Internal Server error.");
            }
        }).finally(()=>{
            console.log("Finished login response to backend.");
        })
    
    }
       return(
        <div id="form-body">
            <div className="form-container">
                <form action="" onSubmit={handleLoginSubmit}>                
                    <label htmlFor="email">Email</label>
                    <input 
                        type="email" 
                        name="email" 
                        id="email" 
                        placeholder="johndoe@gmail.com"
                        onChange={UpdateEmail}
                    />

                    <label htmlFor="password">Password</label>
                    <input 
                        type="password" 
                        name="password" 
                        id="password" 
                        placeholder="password"
                        onChange={UpdatePassword}
                    />
                    <button type="submit">Submit Form</button>
                    <div className='link-login'>
                        <h4>Don't have an account?</h4> 
                        <Link to='/register'><p>Click Here!</p></Link>
                    </div>
                </form>
            </div>
        </div>
    )
}