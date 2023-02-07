import React from 'react'
import Header from './Header'
import Home from './Home'
import Footer from './Footer'
import {Routes, Route} from 'react-router-dom'
import Login from './Login'
import Register from './Register'
import AuthTrade from './AuthTrade'
import { useState } from 'react'




const Main = () => {

  const [authValue, setAuthValue] = useState(false)

  
  return (
    
    <Routes>
      <Route path='/' element={<><Header/><Home /><Footer /></>}></Route>
      <Route path='/login' element={<><Header/><Login setAuthValue={(val) => setAuthValue(val)}/><Footer /></>}></Route>
      <Route path='/register' element={<><Header/><Register /><Footer /></>}></Route>
      {authValue ? <Route path='/trade' element={<AuthTrade />} ></Route> : <Route path='/trade' element={<><Header />
            <h1 style={{color: "white", textAlign: 'center', marginTop:'45vh'}}>Please Login</h1>
            </>} ></Route>}
    </Routes>
  )
}

export default Main