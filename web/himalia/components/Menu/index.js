import React from 'react'
import { useResolvedPath, Link } from 'react-router-dom'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import './style.css'

const menu = ({ isUserAdmin }) => {
    const { pathname } = useResolvedPath()

    const clientMenuItem = isUserAdmin ? (
        <li className={pathname.startsWith('/himalia/clients') ? 'menu-root__list-selected' : ''}>
            <FontAwesomeIcon
                className="menu-root__icon"
                icon={icon({name: 'network-wired'})} />
            <Link to="/himalia/clients">Clients</Link>
        </li>
    ) : null

    return (
        <div role="menu" className="menu-root">
            <h2 className="menu-root__header">User and application<br /> management</h2>
            <ul className="menu-root__list">
                <li className={pathname === '/himalia/profile' ? 'menu-root__list-selected' : ''}>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'user'})} />
                    <Link to="/himalia/profile">Personal Info</Link>
                </li>
                <li className={pathname === '/himalia/security' ? 'menu-root__list-selected' : ''}>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'shield-halved'})} />
                    <Link to="/himalia/security">Password & Security</Link>
                </li>
                <li className={pathname === '/himalia/applications' ? 'menu-root__list-selected' : ''}>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'desktop'})} />
                    <Link to="/himalia/applications">Applications</Link>
                </li>
                {clientMenuItem}
                <li className={pathname === '/himalia/notifications' ? 'menu-root__list-selected' : ''}>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'envelope-open-text'})} />
                    <a href="/himalia/notifications">Notifications</a>
                </li>
                <li>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'arrow-right-from-bracket'})} />
                    <a href="/signout">Sign out</a>
                </li>
            </ul>
        </div>
    )
}

export default menu
