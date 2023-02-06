import React from 'react'
import { Link } from 'react-router-dom';
import mainImg from "../assets/main.png";
import Button from '@mui/material/Button';


const Home = () => {
  return (
    <main className='main'>
        <div className='left'>
            <img src={mainImg} alt="img"/>
        </div>
        <div className='right'>
            <h2> Welcome to our Crypto Trading <br/>  Simulator!</h2>
            <Button className="register" variant="contained" component={Link} to="/register">Register</Button>
        </div>
        </main>
  )
}

export default Home