import React, {useState} from "react";

export const Register = (props) => {
     

    const [firstName, setFirstName] = useState('');
    const [lastName, setlastName] = useState('');
    const [eMail, setEmail] = useState('');
    const [contactNumber, setContactNumber] = useState('');
    const [birthDate, setBirthDate] = useState('');
    const [address, setAddress] = useState('');
    const [username, setUserName] = useState('');
    const [password, setPassword] = useState('');

    async function register(){
        console.warn(firstName, lastName, eMail, contactNumber, birthDate, address, username, password)
        let item ={firstName, lastName, eMail, contactNumber, birthDate, address, username, password}
        let result = await fetch("http://localhost/marketeer/registerUser",{
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
        <div lassName = "auth-form-container">
            <form className="register-form"  >
                 

            <input type = "text" placeholder="First Name" onChange={(e)=>setFirstName(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="Last Name" onChange={(e)=>setlastName(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "email" placeholder="E-Mail" onChange={(e)=>setEmail(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="Contact Number" onChange={(e)=>setContactNumber(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="Birth Date" onChange={(e)=>setBirthDate(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="Address" onChange={(e)=>setAddress(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="username" onChange={(e)=>setUserName(e.target.value)} className = "form-control"></input>
            <br />
            <input type = "text" placeholder="password" onChange={(e)=>setPassword(e.target.value)} className = "form-control"></input>
            <br />

        <button  onClick={(register)}>login</button> 


            </form>
           
        </div>
    )
}

export default Register 