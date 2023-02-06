import React, { useState } from 'react'
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography'
import { styled } from '@mui/material/styles';
import axios from 'axios';


const Register = () => {
  
  const [validation, setValidation ] = useState("")
  
  const registerUser = () => {
      let userName =  document.getElementById('name').value
      let email = document.getElementById('email').value
      let password = document.getElementById('password').value
  
      var pattern = /^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$/; 
      
      if(userName === ""){
        setValidation("Invalid Name")
        return
      }

      if(email.match(pattern) === null ) {
          setValidation("Invalid Email")
          return
      }

      if(password.length < 8) {
        setValidation("Password must contain 8 characters")
        return
      }

      axios.post("http://localhost:8080/api/v1/signup/", {
        "name" : userName,
        "email" : email,
        "password" : password
    })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
      
      
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
    <div className='loginForm' id='loginForm'>
  <Typography className='text' variant="h4" >Register</Typography>
  <CssTextField className='textfield'
              margin="normal"
              required
              fullWidth
              id="name"
              label="Name"
              name="name"
              autoComplete="name"
              autoFocus/>
            <CssTextField className='textfield'
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              type='email'
              name="email"
              autoComplete="email"
              />
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
              onClick={() => registerUser()}>
              Register
            </Button>
            {validation !== "" && <Typography className='text' variant="h6" style={{color:'red'}}>{validation}</Typography>}
  </div>
  )
}

export default Register