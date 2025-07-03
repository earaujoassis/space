import React from 'react'
import { Link } from 'react-router-dom'

const submenu = ({ activeAction }) => {
  return (
    <>
      <ul className="submenu__list">
        <li
          className={
            activeAction === 'all-services' ? 'submenu__list-active' : ''
          }
        >
          <Link to="/services">All services</Link>
        </li>
        <li
          className={
            activeAction === 'new-service' ? 'submenu__list-active' : ''
          }
        >
          <Link to="/services/new">New service</Link>
        </li>
      </ul>
    </>
  )
}

export default submenu
