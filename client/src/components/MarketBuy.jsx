import React from 'react'
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { styled } from '@mui/material/styles';

const MarketBuy = () => {
    const CssTextField = styled(TextField)({
        
        '& .MuiInput-underline:after': {
          borderBottomColor: '#EEEE',
        },
        '& .MuiOutlinedInput-root': {
            '& fieldset': {
                borderColor: '#EEEE',
              },
              '&:hover fieldset': {
                borderColor: '#EEEE',
              },
          '&.Mui-focused fieldset': {
            borderColor: "rgb(30, 158, 30)",
          },
        },
      });
  return (
    <div className="limit-buy">  
    <div className='market-innerdiv'>
        
  <CssTextField className='textfield'
            sx={{input : {color : "white"}}}
            type="number"
            margin="normal"
            required
            fullWidth
            placeholder='Total'
            id="price"
            name="price"
            autoComplete="name"
            />
    <Button 
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              >
              BUY
            </Button>
  </div>
                </div>
  )
}

export default MarketBuy