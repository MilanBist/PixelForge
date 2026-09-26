import "../styles/Register.css"
import { useState } from "react";
import { data, Link } from 'react-router-dom';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import apiClient from "../utils/Base";

// check the username
const checkName = (username)=>{
    if (username === "" || username.length <2 || username.length >50){
        alert("Please enter the valid first name.");
        return false;
    }
    return true;
}

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

export default function Register({setCurrentStatus}){
    const navigate = useNavigate();

    // make the states for all of the given things
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const UpdateUsername = (evt)=>{
        setUsername(evt.target.value);
    }

    const UpdatePassword = (evt)=>{
        setPassword(evt.target.value);
    }
    const UpdateEmail = (evt)=>{
        setEmail(evt.target.value);
    }

    const handleSignupSubmit = (evt)=>{
        evt.preventDefault();
    
        if (String(username).length===0
        ||(email).length === 0 || String(password).length === 0){
            alert("Fill all of credentials.");
            return;
        }
    
        let checkUsername = checkName(username);
        if (checkUsername !== true || checkUsername !== true){
            // alert enter the name validly
            alert("Enter the valid name.");
            return;
        }
    
        let cemail = checkEmail(email);
        if (cemail === false){
            alert("Enter valid email.");
            return;
        }
    
        let cpass = checkPassword(password);
        if (cpass !== true){
            alert("Enter valid password.");
            return;
        }    


        const formdata = {
            userName: username,
            email: email,
            password: password,
        };

        console.log("Form data to send to the register page is: ", formdata);

        // send this to the frontend using the axios
        apiClient.post("/register", formdata).then((resp) =>{
            // if the response status is 202
            console.log("Register response is: ", resp.data);
            localStorage.setItem("accessToken", resp["data"]["data"]["access"]);
            localStorage.setItem("refreshToken", resp["data"]["data"]["refresh"])
            setCurrentStatus("Logged In");
            navigate("/");
        }).catch((err) =>{
            const responseStatus = err.response.status;
            switch (responseStatus){
                case 400:
                    alert("Wrong sending method.");
                case 401:
                    alert("Login credentials.");
                case 404:
                    alert("User doesn't exist. Please register.");
                case 409:
                    alert("User already exist. Please login.")
                case 500:
                    alert("Internal Server error.");
            }
        }).finally(()=>{
            console.log("Successfully register response.");
        })
    }

    return(
        <div id='form-body'>
            <div className="form-container">
                <h2>Create Account</h2>
                <form onSubmit={handleSignupSubmit}>

                    <label htmlFor="username">Username</label>
                    <input 
                        type="text" 
                        onChange={UpdateUsername} 
                        name="username" 
                        placeholder="Username"
                    />


                    <label htmlFor="email">Email</label>
                    <input 
                        type="email" 
                        onChange={UpdateEmail}
                        name="email" 
                        id="email" 
                        placeholder="Email"
                    />

                    <label htmlFor="password">Password</label>
                    <input 
                        type="password" 
                        name="password" 
                        id="password" 
                        onChange={UpdatePassword}
                        placeholder="Password"
                    />
                    <button type="submit" onSubmit={handleSignupSubmit}>Sumbit Form</button>
                    <div className='link-login'>
                        <h4>Already have an account?</h4> 
                        <Link to='/login'><p>Click Here!</p></Link>
                    </div>
                </form>
            </div>
        </div>
    )
}