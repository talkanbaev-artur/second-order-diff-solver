import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import Calc from './Calc'
import './index.css'

import { Route, createBrowserRouter, RouterProvider, Link } from "react-router-dom"

const notFound = () => {
  return (<div>Not found</div>)
}

const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>
  },  {
    path: "/calc",
    element: <Calc/>,
  },
  { path: "*", element: notFound() }
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
