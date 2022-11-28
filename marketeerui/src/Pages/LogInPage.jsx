 
import React, {useRef, useState, useEffect} from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export const LogIn = (props) => {
let navigate = useNavigate();
const [username, setUserName] = useState('');
const [password, setPassword] = useState('');
 
 let displayData
    const logInUrl = "http://localhost/marketeer/login"
    const systemCheck = "http://localhost:/marketeer/systemcheck"
    async function login(){
        console.warn(username, password)
        console.log(username, password)
        let item ={username, password}
        let result = await fetch(logInUrl,{
        mode: "no-cors",    
        method: 'POST',
                headers:{
                    "Content-Type": "application/json",
                    "Accept": 'application/json'
                },
                body: JSON.stringify(item) 
        });
        

        // fetch(systemCheck,{ mode: "no-cors", }).then(response=>response.json())
        // .then(responseData => {
        //     displayData = responseData.map(function(todo){
        //         return(
        //             <p key = {todo.code}></p>
        //         )
        //     })
        //     //console.log(responseData)
        // })
    }

    // function componentDidMount() {
    //     // GET request using fetch with set headers
    //     const headers = { 'Content-Type': 'application/json' }
    //     fetch(systemCheck, { mode: "no-cors",   headers })
    //         .then(response => response.json())
    //         .then(data => this.setState({ totalReactPackages: data.code }));
    // }
   

    // const URL = "http://localhost/marketeer/login";
    // const [postResult, setPostResult] = useState(null);
  
    // const fortmatResponse = (res) => {
    //   return JSON.stringify(res, null, 2);
    // }
    
    // async function login() {
     
    //     let postData ={username, password}
    //     console.warn(postData)
    //   try {
    //     const res = await fetch(URL, {
    //         mode: "no-cors",
    //       method: "post",
    //       headers: {
    //         "Content-Type": "application/json",
    //         "Accept": 'application/json'
    //       },
    //       body: JSON.stringify(postData),
    //     });
  
    //     if (!res.ok) {
    //       const message = `An error has occured: ${res.status} - ${res.statusText}`;
    //       throw new Error(message);
    //     }
  
    //     const data = await res.json();
    //     console.log(data)
       
    //     const result = {
    //       status: res.status + "-" + res.statusText,
    //       headers: {
    //         "Content-Type": res.headers.get("Content-Type"),
    //         "Content-Length": res.headers.get("Content-Length"),
    //       },
    //       data: data,
    //     };
  
    //     setPostResult(fortmatResponse(result));
    //   } catch (err) {
    //     setPostResult(err.message);
    //   }
    //   console.log(postResult)
    // }
 

  

    return(
        <div>hello
            <div className="col-sm-6 offset-sm-3">
            <input type = "text" placeholder="username" onChange={(e)=>setUserName(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="password" onChange={(e)=>setPassword(e.target.value)} className = "form-control"></input>
            </div>

            <button   onClick={() => {
                
              //login(); 
              login();
                //componentDidMount();
               
                // navigate("/mainmenu")
            }}>login</button>
        </div>
    )
}


 


export default LogIn