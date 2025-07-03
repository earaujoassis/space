import React from 'react'
import { useResolvedPath, Link } from 'react-router-dom'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import './style.css'

const selectedItemClass = 'menu-root__list-selected'

const menu = ({ isUserAdmin }) => {
  const { pathname } = useResolvedPath()

  const clientMenuItem = isUserAdmin ? (
    <li className={pathname.startsWith('/clients') ? selectedItemClass : ''}>
      <FontAwesomeIcon
        className="menu-root__icon"
        icon={icon({ name: 'network-wired' })}
      />
      <Link to="/clients">Clients</Link>
    </li>
  ) : null

  return (
    <div role="menu" className="menu-root">
      <h2 className="menu-root__header">
        User and application
        <br /> management
      </h2>
      <ul className="menu-root__list">
        <li className={pathname === '/profile' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'user' })}
          />
          <Link to="/profile">Personal Info</Link>
        </li>
        <li className={pathname === '/emails' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'envelope' })}
          />
          <Link to="/emails">Emails</Link>
        </li>
        <li className={pathname === '/security' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'shield-halved' })}
          />
          <Link to="/security">Password & Security</Link>
        </li>
        {clientMenuItem}
        <li className={pathname === '/applications' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'desktop' })}
          />
          <Link to="/applications">Applications</Link>
        </li>
        <li className={pathname === '/services' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'bell-concierge' })}
          />
          <Link to="/services">Services</Link>
        </li>
        <li className={pathname === '/notifications' ? selectedItemClass : ''}>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'envelope-open-text' })}
          />
          <Link to="/notifications">Notifications</Link>
        </li>
        <li>
          <FontAwesomeIcon
            className="menu-root__icon"
            icon={icon({ name: 'arrow-right-from-bracket' })}
          />
          <a href="/signout">Sign out</a>
        </li>
      </ul>
    </div>
  )
}

export default menu
