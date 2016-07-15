import React from 'react';

import Row from '../../core/components/row.jsx';
import Columns from '../../core/components/columns.jsx';
import Applications from './Applications.jsx';
import Profile from './Profile.jsx';
import AccountLog from './AccountLog.jsx';

const SubViews = {
    'Applications': Applications,
    'Profile': Profile,
    'AccountLog': AccountLog
}

export default class Settings extends React.Component {
    constructor() {
        super();
        this.state = { currentSubView: Applications };
        this._setTab = this._setTab.bind(this);
        this._isActive = this._isActive.bind(this);
    }

    componentWillUnmount() {
        this._setTab();
    }

    render() {
        let childComponent = this.state.currentSubView;
        return (
            <Row className="settings-wrapper">
                <Columns className="medium-3 large-2">
                    <ul className="side-nav">
                        <li><a href="/applications" className={this._isActive(Applications)}>Applications</a></li>
                        <li><a href="/profile" className={this._isActive(Profile)}>Profile</a></li>
                        <li><a href="/account-log" className={this._isActive(AccountLog)}>Account Log</a></li>
                        <li class="divider"></li>
                        <li><a href="/settings">Settings</a></li>
                        <li><a href="/signout">Signout</a></li>
                    </ul>
                </Columns>
                <Columns className="medium-9 large-10 settings-content">
                    <div className="breadcrumbs-custom">
                        <ul>
                            <li>Dashboard</li>
                            <li>{childComponent.name}</li>
                        </ul>
                    </div>
                    {React.createElement(childComponent, {})}
                </Columns>
            </Row>
        );
    }

    _isActive(componentClass) {
        return this.state.currentSubView == componentClass ? 'active' : ''
    }

    _setTab() {
        console.log("not implemented")
    }
};
