import React, { useState } from 'react'

import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography'
import { styled } from '@mui/material/styles';
import {  useNavigate } from 'react-router-dom';
// import axios from 'axios';

const Login = (props) => {

  const [validation, setValidation] = useState("")
  const navigate = useNavigate();

  const loginUser = () => {
    
    let email = document.getElementById('email').value
    let password = document.getElementById('password').value

    fetch("http://localhost:8080/api/v1/login/", {
      method : 'POST',
      headers : {
        'Content-Type': 'application/json'
      },
      body : JSON.stringify ({
        "email" : email,
      "password" : password
      }),
      credentials : 'include'
  }).then(res => {
    if(res.status === 200){
          props.setAuthValue(true)
          navigate('/trade')
    }
  })
  .catch(err => {
    console.log(err);
    setValidation(err.response.data.error)
  })
  
  }

    const CssTextField = styled(TextField)({
        '& label.Mui-focused': {
          color: '#00ADB5',
        },
        '& .MuiInput-underline:after': {
          borderBottomColor: '#00ADB5',
        },
        '& .MuiOutlinedInput-root': {
          '&.Mui-focused fieldset': {
            borderColor: '#00ADB5',
          },
        },
      });
  return (
    <div className='loginForm'>
  <Typography className='text' variant="h4" >Log In</Typography>
            <CssTextField className='textfield'
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus/>
            <CssTextField className='textfield'
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              
              autoComplete="current-password"/>
            <Button 
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              onClick={loginUser}
            >
              Log In
            </Button>
            
            {validation !== "" && <Typography className='text' variant="h6" style={{color:'red'}}>{validation}</Typography>}
  </div>

  )
}

export default Login