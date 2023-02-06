import React from 'react'

const Orders = (props) => {
    var a = new Date(props.openedAt * 1000);
  var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
  var year = a.getFullYear();
  var month = months[a.getMonth()];
  var date = a.getDate();
  var hour = a.getHours();
  var min = a.getMinutes();
  var sec = a.getSeconds();
  var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min + ':' + sec ;
  
  return (
        <tr>
            <td>{time}</td>
            <td>{props.pair}</td>
            <td>{props.type}</td>
            <td>{props.side}</td>
            <td>{props.price}</td>
            <td>{props.amount}</td>
            <td>{props.total}</td>
            <td><button style={{cursor:"pointer"}}>X</button></td>
        </tr>
    
  )
}

export default Orders