import Button from '@mui/material/Button';
import main from '../../assets/main.png';

export default function Main() {
    return <>
    
    <div className="main">
        <div className="right">
            <img src={main} alt="" />
        </div>
        <div className="left">
            <h2>
            Welcome to our Crypto Trading Simulator!
            </h2>
            <Button id="register" variant="contained">Get Started</Button>            
        </div>
    </div>
    </>
}