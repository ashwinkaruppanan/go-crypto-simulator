import React from 'react'
import Orders from './Orders'
import History from './History'
import hisData from '../temp/history.json'
import data from '../temp/open.json'

const orderData = data.map(order => <Orders 
    openedAt={order.opened_at}
    pair={order.pair}
    type={order.type}
    side={order.side}
    price={order.price}
    amount={order.amount}
    total={order.total}
    cancel={order.OrderID}
/>)

const historyData = hisData.map(order => <History 
  openedAt ={order.opened_at}
  executedAt = {order.executed_at}
  pair={order.pair}
  type={order.type}
  side={order.side}
  price={order.price}
  amount={order.amount}
  total={order.total}
  />)

const OpenHistory = () => {
  return (
    <div className="OpenHistory">
      <div className="openorders">
      <p>OPEN ORDERS</p>
      <div className="scrollinner">        
      <table>
        <th>OPENED AT</th>
        <th>PAIR</th>
        <th>TYPE</th>
        <th>SIDE</th>
        <th>PRICE</th>
        <th>AMOUNT</th>
        <th>TOTAL</th>
        <th>CANCEL</th>
        {orderData}
      </table>
      </div>
      </div>
      <div className="history">
        <p>HISTORY</p>
        <div className="scrollinner">
        <table>
        <th>OPENED AT</th>
        <th>EXECUTED AT</th>
        <th>PAIR</th>
        <th>TYPE</th>
        <th>SIDE</th>
        <th>PRICE</th>
        <th>AMOUNT</th>
        <th>TOTAL</th>
        {historyData}
        </table>
        </div>
      </div>
    </div>
  )
}

export default OpenHistory