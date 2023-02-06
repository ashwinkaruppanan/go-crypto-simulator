import React from 'react'
import logo from "../assets/logo.png";
import Button from '@mui/material/Button';
import {Link }from 'react-router-dom'

const Header = () => {
  return (
    <div className="navbar">
          <div className="left">
            <div className="logo">
              <Link to='/'><img src={logo} alt="" /></Link>
              
            </div>
          </div>
          <div className="buttons">              
              <Button className="login" variant="outlined" component={Link} to="/login">Log In</Button>
              <Button className="register" variant="contained" component={Link} to="/register">Register</Button>
          </div>
      </div>  
  )
}

export default Header