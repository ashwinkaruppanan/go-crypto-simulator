import React from 'react'

import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography'
import { styled } from '@mui/material/styles';

const Login = () => {
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
            >
              Log In
            </Button>
  </div>

  )
}

export default Login