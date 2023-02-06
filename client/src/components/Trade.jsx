import React from 'react';
import logo from '../assets/logo.png';
import Button from '@mui/material/Button';
import TradingViewWidget from './TradingView';
import LimitBuy from './LimitBuy';
import LimitSell from './LimitSell';
import { useState } from 'react';
import MarketBuy from './MarketBuy';
import MarketSell from './MarketSell';

const Trade = (props) => {

    const [active, setActive] = useState("limit")
    if(props.auth !== 'true'){
        return (
            <>
            <h1 style={{color: "white", textAlign: 'center', marginTop:'45vh'}}>Please Login</h1>
            </>
        )
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
            <h5>6754 $</h5>
            <h5>0.00008 BTC</h5>
            <Button className="login" variant="outlined" >Log Out</Button>
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