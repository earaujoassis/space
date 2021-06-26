import React from 'react';
import { Link } from 'react-router-dom';

import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';

import UserStore from '../stores/users';

const pathToTitle = {
  '/applications': 'Applications',
  '/profile': 'Profile'
};

export default class Settings extends React.Component {
  constructor (props) {
    super(props);
    this._isActive = this._isActive.bind(this);
    this.state = { error: false };
    this._updateFromStore = this._updateFromStore.bind(this);
  }

  componentDidMount () {
    UserStore.addChangeListener(this._updateFromStore);
  }

  componentWillUnmount () {
    UserStore.removeChangeListener(this._updateFromStore);
  }

  render () {
    const { pathname } = this.props.location;

    return (
            <Row className="settings-wrapper">
                <Columns className="small-12 medium-3 large-2">
                    <ul className="side-nav">
                        <li><Link to="/applications" className={this._isActive('/applications')}>Applications</Link></li>
                        <li className="divider">
                            <Link to="/profile" className={this._isActive('/profile')}>Profile</Link>
                        </li>
                        <li><a href="/signout">Sign out</a></li>
                    </ul>
                </Columns>
                <Columns className="small-12 medium-9 large-10 settings-content">
                    <div className="breadcrumbs-custom">
                        <ul>
                            <li>Dashboard</li>
                            <li>{pathToTitle[pathname]}</li>
                        </ul>
                    </div>
                    {this.state.error && (
                        <div className="token-error">Your action token is expired. Please refresh your page.</div>
                    )}
                    {this.props.children}
                </Columns>
            </Row>
    );
  }

  _isActive (path) {
    return this.props.location.pathname === path ? 'active' : '';
  }

  _updateFromStore () {
    this.setState({ error: UserStore.getState().payload.error === 'access_denied' });
  }
}
