import React, { useEffect, useState } from 'react'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import './style.css'

const scopesGroup = ({ initialScopes, onChange }) => {
    const [selected, setSelect] = useState(initialScopes)

    useEffect(() => {
        setSelect(initialScopes)
    }, [initialScopes])

    const onSelection = (value) => {
        const newSelected = [...selected]
        if (newSelected.includes(value)) {
            const index = selected.indexOf(value)
            newSelected.splice(index, 1)
        } else {
            newSelected.push(value)
        }
        setSelect(newSelected)
        onChange(newSelected)
    }

    const setIcon = (value) => {
        if (selected.includes(value)) {
            return (
                <FontAwesomeIcon
                    className="globals__radio-icon"
                    icon={icon({name: 'square-check', style: 'regular'})} />
            )
        } else {
            return (
                <FontAwesomeIcon
                    className="globals__radio-icon"
                    icon={icon({name: 'square', style: 'regular'})} />
            )
        }
    }

    return (
        <>
            <div className="scopes-group__entry" onClick={() => onSelection('public')}>
                <div className="scopes-group__checkbox">
                    <span className="scopes-group__icon">
                        {setIcon('public')}
                    </span>
                    <span className="scopes-group__label">
                        <code>public</code>
                    </span>
                </div>
                <div className="scopes-group__description">
                    <p>Can perform:</p>
                    <p>Obtain data for user authentication and authorization only</p>
                    <p>OAuth 2.0 only; cannot use OIDC Provider endpoints</p>
                </div>
            </div>
            <div className="scopes-group__entry" onClick={() => onSelection('openid')}>
                <div className="scopes-group__checkbox">
                    <span className="scopes-group__icon">
                        {setIcon('openid')}
                    </span>
                    <span className="scopes-group__label">
                        <code>openid</code>
                    </span>
                </div>
                <div className="scopes-group__description">
                    <p>Can perform:</p>
                    <p>Obtain data for user authentication and authorization only</p>
                    <p>OIDC only, cannot use OAuth 2.0 Provider endpoints</p>
                </div>
            </div>
            <div className="scopes-group__entry" onClick={() => onSelection('profile')}>
                <div className="scopes-group__checkbox">
                    <span className="scopes-group__icon">
                        {setIcon('profile')}
                    </span>
                    <span className="scopes-group__label">
                        <code>profile</code>
                    </span>
                </div>
                <div className="scopes-group__description">
                    <p>Can perform:</p>
                    <p>Obtain data for user authentication and authorization only</p>
                    <p>OIDC only; read user profile data under OIDC Provider endpoints</p>
                </div>
            </div>
            <div className="scopes-group__entry">
                <div className="scopes-group__checkbox">
                    <span className="scopes-group__icon">
                        <FontAwesomeIcon
                            className="globals__radio-icon"
                            icon={icon({name: 'square-minus', style: 'regular'})} />
                    </span>
                    <span className="scopes-group__label">
                        <code>read</code>
                    </span>
                </div>
                <div className="scopes-group__description">
                    <p>Can perform:</p>
                    <p>Obtain data for user authentication and authorization only</p>
                    <p>(Removed) Read user profile data under OAuth 2.0 Provider endpoints</p>
                </div>
            </div>
            <div className="scopes-group__entry">
                <div className="scopes-group__checkbox">
                    <span className="scopes-group__icon">
                        <FontAwesomeIcon
                            className="globals__radio-icon"
                            icon={icon({name: 'square-minus', style: 'regular'})} />
                    </span>
                    <span className="scopes-group__label">
                        <code>write</code>
                    </span>
                </div>
                <div className="scopes-group__description">
                    <p>Can perform:</p>
                    <p>Obtain data for user authentication and authorization only</p>
                    <p>Update user data &mdash; this scope can&#39;t be claimed by other applications</p>
                </div>
            </div>
        </>
    )
}

export default scopesGroup
