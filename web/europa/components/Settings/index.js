import React, { useEffect } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';

import { loadData } from '@core/actions/internal';

import { TokenExpirationMessage } from './TokenExpirationMessage';

const pathToTitle = {
  '/applications': 'Applications',
  '/profile': 'Profile'
};

const isActive = (path, currentPath) => path === currentPath ? 'active' : '';

const settings = ({ loadData, children, location, data, tokenExpired }) => {
  useEffect(() => {
    if (data.action_token === undefined) {
      loadData();
    }
  }, [data.action_token]);

  if (data.action_token === undefined) {
    return null;
  }

  return (
    <Row className="settings-wrapper">
      <Columns className="small-12 medium-3 large-2">
        <ul className="side-nav">
          <li>
            <Link to="/applications" className={isActive('/applications', location.pathname)}>Applications</Link>
          </li>
          <li className="divider">
            <Link to="/profile" className={isActive('/profile', location.pathname)}>Profile</Link>
          </li>
          <li>
            <a href="/signout">Sign out</a>
          </li>
        </ul>
      </Columns>
      <Columns className="small-12 medium-9 large-10 settings-content">
        <div className="breadcrumbs-custom">
          <ul>
            <li>Dashboard</li>
            <li>{pathToTitle[location.pathname]}</li>
          </ul>
        </div>
        <TokenExpirationMessage tokenExpired={tokenExpired} />
        {children}
      </Columns>
    </Row>
  );
};

const mapStateToProps = state => {
  return {
    data: state.internal.data,
    tokenExpired: state.internal.tokenExpired
  };
};

const mapDispatchToProps = dispatch => {
  return {
    loadData: () => dispatch(loadData())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(settings);
