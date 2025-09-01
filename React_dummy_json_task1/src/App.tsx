import { useEffect, useRef, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import "./components/Aopp.css"
function App() {
  interface User{
     userId:number, 
     id: number,
    title:string,
  completed:string,
  }
  const[users,second] =useState<User[]>([])

  useEffect(()=>{
    fetch('https://jsonplaceholder.typicode.com/todos/')
      .then(response => response.json())
      .then(json => second(json))
  },[])
  
  return(
  <table  className='app'  border={1} >
    <thead>
    <th>id</th>
    <th>userid</th>
    <th>Title</th>
    <th>status</th>
    </thead>
    <tbody>   {users.map((user) => (
            <tr key={user.userId}>
              <td>{user.id}</td>
              <td>{user.userId}</td>
              <td>{user.title}</td>
              <td>{user.completed?"true" :"false"}  </td>
            </tr>
          ))}
      </tbody>
    </table>)
};

export default App
