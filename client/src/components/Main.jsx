import React from 'react'
import Header from './Header'
import Home from './Home'
import Footer from './Footer'
import {Routes, Route} from 'react-router-dom'
import Login from './Login'
import Register from './Register'

import Trade from './Trade'
import OpenHistory from './OpenHistory'

const Main = () => {
  return (
    <>
    
    <Routes>
      <Route path='/' element={<><Header/><Home /><Footer /></>}></Route>
      <Route path='/login' element={<><Header/><Login /><Footer /></>}></Route>
      <Route path='/register' element={<><Header/><Register /><Footer /></>}></Route>
      <Route path='/trade' element={<><Trade auth="true"/><OpenHistory /><Footer /></>}></Route>
    </Routes>
    
    
    </>
  )
}

export default Main