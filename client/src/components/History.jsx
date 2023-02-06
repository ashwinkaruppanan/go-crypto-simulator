import React from 'react'

const History = (props) => {
    var a = new Date(props.openedAt * 1000);
    var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
    var year = a.getFullYear();
    var month = months[a.getMonth()];
    var date = a.getDate();
    var hour = a.getHours();
    var min = a.getMinutes();
    var sec = a.getSeconds();
    var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min + ':' + sec ;
    
    var b = new Date(props.executedAt * 1000);
    var hmonths = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
    var hyear = b.getFullYear();
    var hmonth = hmonths[b.getMonth()];
    var hdate = b.getDate();
    var hhour = b.getHours();
    var hmin = b.getMinutes();
    var hsec = b.getSeconds();
    var htime = hdate + ' ' + hmonth + ' ' + hyear + ' ' + hhour + ':' + hmin + ':' + hsec ;
    
  return (
    <tr>
        <td>{time}</td>
        <td>{htime}</td>
        <td>{props.pair}</td>
        <td>{props.type}</td>
        <td>{props.side}</td>
        <td>{props.price}</td>
        <td>{props.amount}</td>
        <td>{props.total}</td>
    </tr>
  )
}

export default History