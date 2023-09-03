import React from 'react'
import { Link } from 'react-router-dom'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import './style.css'

const menu = () => {
    return (
        <div role="menu" className="menu-root">
            <h2 className="menu-root__header">User and application<br /> management</h2>
            <ul className="menu-root__list">
                <li className="menu-root__list-selected">
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'user'})} />
                    <Link to="/himalia/profile">Personal Info</Link>
                </li>
                <li>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'shield-halved'})} />
                    <a href="/himalia/security">Password & Security</a>
                </li>
                <li>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'desktop'})} />
                    <a href="/himalia/applications">Applications</a>
                </li>
                <li>
                    <FontAwesomeIcon
                        className="menu-root__icon"
                        icon={icon({name: 'network-wired'})} />
                    <a href="/himalia/clients">Clients</a>
                </li>
                <li>
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
