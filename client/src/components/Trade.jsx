import React from 'react';
import logo from '../assets/logo.png';
import Button from '@mui/material/Button';
import TradingViewWidget from './TradingView';
import LimitBuy from './LimitBuy';
import LimitSell from './LimitSell';
import { useState } from 'react';
import MarketBuy from './MarketBuy';
import MarketSell from './MarketSell';
import {  useNavigate } from 'react-router-dom';

import { useEffect } from 'react';

const logout = (navigate) => {
        fetch("http://localhost:8080/api/v1/logout/", {method : 'DELETE', credentials : 'include'})
        .then(res => console.log(res))
        .catch(err => console.log(err))
        navigate('/login')
}

const Trade = () => {

    const [active, setActive] = useState("limit")
    
    const [balance , setBalance] = useState({
        btc : 0,
        usd : 0
    })
    const [refCount , setRefCount] = useState(0)
    
    let navigate = useNavigate();


    useEffect(() => {

        fetch("http://localhost:8080/api/v1/balance/", 
        {method : 'GET' , credentials : 'include'})
        .then(res => res.json())
        .then(data => {
            console.log(data);
            setBalance(() => ({
                btc : data.bitcoin,
                usd : data.fiat
            }))
        })
        .catch(err => console.log(err))
    },[refCount]) 

    function setRef() {
        var newCount = refCount + 1
        setRefCount(newCount)
    }

    

  return (
    <>
    <div className="trade-nav">
        <div className="left">
        <div className="logo">
              <img src={logo} alt="" />              
            </div>
        </div>
        <div className="right">
            <h4>BALANCE </h4>
            <h5>{balance.usd} $</h5>
            <h5>{balance.btc} BTC</h5>
            <Button className="login" variant="outlined" onClick={() => logout(navigate)}>Log Out</Button>
            <button onClick={() => setRef()}>+</button>
        </div>
    </div>
    <div className="trade-view">
        <div className="left">
                <TradingViewWidget />
        </div>
        <div className="right">
                <div className="limit-market">
                    <p onClick={() => setActive("limit")} style={{color: active === 'limit' && "#00ADB5"}}>LIMIT</p>
                    <p onClick={() => setActive("market")} style={{color: active === 'market' && "#00ADB5"}}>MARKET</p>
                </div>
                {active === 'limit' && <><LimitBuy /><LimitSell /></>}
                {active === 'market' && <><MarketBuy /><MarketSell /></>} 
            </div>
        </div>  
    </>
  )
}

export default Trade