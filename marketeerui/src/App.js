import React, {useState} from "react";
import './App.css';
 
import MainMenuPage from './Pages/MainMenuPage';
import UserProfile from './Pages/UserProfile';
import LogIn from './Pages/LogInPage';
import Error from './Pages/ErrorPage';
import RegisterUser from './Pages/RegistrationPage';
 
import{ BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
 
function App() {
  // const [currentForm, setCurrentForm] = useState('login');
  // const toggleForm = (formName) =>{
  //   setCurrentForm(formName);


    return (
      
    <Router>
      <nav>
        <Link to="login"> Log in Page </Link>
        <br></br>
        <Link to="register">register </Link>
        <br></br>
        <Link to="mainmenu"> Main Menu </Link>
        <br></br>
        <Link to="userprofile"> User Profile </Link>
      </nav>


      <Routes> 
        {/* user credential pages */}
        <Route path="/login" element = {<LogIn/>}/>
        
        {/* in-app pages */}
        <Route path="/mainmenu" element = {<MainMenuPage/>}/>
        <Route path="/profile" element = {<UserProfile/>}/> 
        <Route path="/register" element = {<RegisterUser/>}/> 
        <Route path="*" element = {<Error/>}/> 

      </Routes>
    </Router>
    );
    

      // return <div>
      //   HELLO
      // </div>

  }


  // return (
  //   <div className="App">
  //     {
  //       currentForm === "login" ? <Login onFormSwitch={toggleForm}/> : <RegisterUser onFormSwitch={toggleForm} />
  //     }
      
  //   </div>
  // );
//}

export default App;
