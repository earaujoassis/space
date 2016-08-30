import React from 'react'
import { Link } from 'react-router'

import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

export default class Settings extends React.Component {
    constructor() {
        super()
        this._isActive = this._isActive.bind(this)
    }

    render() {
        return (
            <Row className="settings-wrapper">
                <Columns className="medium-3 large-2">
                    <ul className="side-nav">
                        <li><Link to="/" className={this._isActive('Applications')}>Applications</Link></li>
                        <li><Link to="/profile" className={this._isActive('Profile')}>Profile</Link></li>
                        <li><Link to="/log" className={this._isActive('Account log')}>Account log</Link></li>
                        <li className="divider"></li>
                        <li><a href="/signout">Sign out</a></li>
                    </ul>
                </Columns>
                <Columns className="medium-9 large-10 settings-content">
                    <div className="breadcrumbs-custom">
                        <ul>
                            <li>Dashboard</li>
                            <li>{this.props.routes[1].name}</li>
                        </ul>
                    </div>
                    {this.props.children}
                </Columns>
            </Row>
        )
    }

    _isActive(name) {
        return this.props.routes[1].name == name ? 'active' : ''
    }
}
