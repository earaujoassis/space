import React, { useEffect } from 'react';
import { connect } from 'react-redux';

import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';
import { Entry } from '@core/components/Forms.jsx';

import { fetchProfile, adminify, requestUserUpdate } from '@europa/actions/users';

import { Adminify } from './Adminify';
import { UpdatePassword } from './UpdatePassword';
import { RecreateRecoveryCode } from './RecreateRecoveryCode';

const profile = ({
  fetchProfile,
  adminify,
  requestUserUpdate,
  user,
  loading = [],
  features = [],
  token,
  userId
}) => {
  useEffect(() => {
    if (userId !== undefined) {
      fetchProfile(token, userId);
    }
  }, [userId]);

  if (loading.includes('users')) {
    return (
      <Row>
        <Columns className="small-offset-1 small-10 end">
          <p className="text-center">Loading...</p>
        </Columns>
      </Row>
    );
  }

  if (user === undefined) {
    return (
      <Row>
        <Columns className="small-offset-1 small-10 end">
          <p className="text-center">Nothing has been loaded...</p>
        </Columns>
      </Row>
    );
  }

  return (
    <Row>
      <Columns className="small-offset-1 small-10 end">
        <Row className="profile">
          <Columns className="small-12">
            <Entry field="Name" value={`${user.first_name} ${user.last_name}`} />
            <Entry field="Username" value={user.username} />
            <Entry field="Email" value={user.email} />
            <Entry field="Timezone" value={user.timezone_identifier} />
            <Adminify adminify={adminify} user={user} features={features} />
          </Columns>
        </Row>
        <Row className="profile-actions">
          <Columns className="small-12">
            <UpdatePassword requestUserUpdate={requestUserUpdate} user={user} />
            <RecreateRecoveryCode requestUserUpdate={requestUserUpdate} user={user} />
          </Columns>
        </Row>
      </Columns>
    </Row>
  );
};

const mapStateToProps = state => {
  return {
    loading: state.internal.loading,
    features: state.internal.data.features,
    token: state.internal.data.action_token,
    userId: state.internal.data.user_id,
    user: state.users.user
  };
};

const mapDispatchToProps = dispatch => {
  return {
    fetchProfile: (token, userId) => dispatch(fetchProfile(token, userId)),
    adminify: (token, data) => dispatch(adminify(token, data)),
    requestUserUpdate: (id) => dispatch(requestUserUpdate(id))
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(profile);
