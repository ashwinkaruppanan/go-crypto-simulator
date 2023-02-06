import React from 'react'
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { styled } from '@mui/material/styles';

const MarketSell = () => {
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
            borderColor: "rgb(255, 31, 31)",
          },
        },
    });
  return (
    <div className="limit-sell">  
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
              SELL
            </Button>

                </div>
  </div>   
  )
}

export default MarketSell