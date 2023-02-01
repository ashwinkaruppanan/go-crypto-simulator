import Button from '@mui/material/Button';
import logo from '../../assets/logo.png';


function Header() {
    return <>
        <div className="header">
            <div className="left">
            <div className="logo">
                <img src={logo} alt="" />
            </div>
            </div>
            <div className="right">
                <Button id="login" variant="outlined">Log In</Button>
                <Button id="register" variant="contained">Register</Button>
            </div>
        </div>        
    </>
}
export default Header;