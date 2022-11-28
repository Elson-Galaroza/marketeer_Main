import React, {useEffect, useState} from "react";
import { useHistory } from 'react-router-dom'

export const Login = (props) => {
    const [username, setUserName] = useState('');
    const [password, setPassword] = useState('');

     
    async function login(){
        console.warn(username, password)
        let item ={username, password}
        let result = await fetch("http://localhost/marketeer/login",{
        mode: "no-cors",    
        method: 'POST',
                headers:{
                    "Content-Type": "application/json",
                    "Accept": 'application/json'
                },
                body: JSON.stringify(item) 
        });
        result = await result.json();

    }

    return(
        <div  className = "auth-form-container">
             
        <div className="col-sm-6 offset-sm-3">
            <input type = "text" placeholder="username" onChange={(e)=>setUserName(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="password" onChange={(e)=>setPassword(e.target.value)} className = "form-control"></input>
        </div>
        <button   onClick={(login)}>login</button>
        {/* <button   onClick={(e)}>login</button> */}

            <button className="link-btn" onClick={() => props.onFormSwitch('register')}>Register here if you don't have an account</button>
        </div>
    )
}

export default Login 