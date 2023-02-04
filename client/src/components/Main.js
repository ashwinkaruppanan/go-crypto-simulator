import React from 'react';
import logo from "../assets/logo.png";
import mainImg from "../assets/main.png";
import Button from '@mui/material/Button';

export default function Main(){
  return (<>

    {/* header */}
    <div className="navbar">
            <div className="left">
              <div className="logo">
                <img src={logo} alt="" />
              </div>
            </div>
            <div className="buttons">
                <Button className="login" variant="outlined">Log In</Button>
                <Button className="register" variant="contained">Register</Button>
            </div>
        </div>       

    {/* main */}
    <main className='main'>
        <div className='left'>
            <img src={mainImg} alt="img"/>
        </div>
        <div className='right'>
            <h2> Welcome to our Crypto Trading <br/>  Simulator!</h2>
            <Button className="register" variant="contained">Register</Button>
        </div>
    </main>


    {/* footer */}
    <div className="footer">
        <p>copyright &#169;	{new Date().getFullYear()}</p>        

    </div>
  </>
  );
}