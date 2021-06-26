import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';

import AdminContent from '@core/components/AdminContent.jsx';
import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';

import { fetchActiveClients, revokeActiveClient } from '@europa/actions/users';
import { fetchClients, createClient, updateClient } from '@europa/actions/clients';

import AllApplications from './AllApplications';
import NewApplication from './NewApplication';
import UserApplications from './UserApplications';

const ACCORDION_USER_APPLICATIONS = 'user-applications';
const ACCORDION_ALL_APPLICATIONS = 'all-applications';
const ACCORDION_NONE = 'none';

const setAccordionTransition = (current, context) => {
  if (current === context) {
    return ACCORDION_NONE;
  } else {
    return context;
  }
};

const applications = ({
  fetchActiveClients,
  revokeActiveClient,
  fetchClients,
  createClient,
  updateClient,
  user = {},
  userClients = [],
  clients = [],
  loading = [],
  token,
  userId
}) => {
  const [openAccordion, setOpenAccordion] = useState('none');

  useEffect(() => {
    if (userId !== undefined) {
      fetchActiveClients(token, userId);
    }
  }, [userId]);

  useEffect(() => {
    if (user.is_admin) {
      fetchClients(token);
    }
  }, [user.is_admin]);

  return (
    <div role="main">
      <Row>
        <Columns className="small-12">
          <div className="jupiter-accordion">
            <div className={`jupiter-accordion-child ${openAccordion === ACCORDION_USER_APPLICATIONS ? 'open' : ''}`}>
              <h2 className="jupiter-accordion-title">
                <a
                  href="#my"
                  onClick={(e) => {
                    e.preventDefault();
                    setOpenAccordion(setAccordionTransition(openAccordion, ACCORDION_USER_APPLICATIONS));
                  }}
                >
                  My applications
                </a>
              </h2>
              <div className="jupiter-accordion-body">
                <Row>
                  <Columns className="small-offset-1 small-10 end">
                    <Row className="applications">
                      <UserApplications
                        revokeActiveClient={(clientId) => revokeActiveClient(token, userId, clientId)}
                        clients={userClients}
                        loading={loading}
                      />
                    </Row>
                  </Columns>
                </Row>
              </div>
            </div>
            <AdminContent user={user}>
              <div className={`jupiter-accordion-child ${openAccordion === ACCORDION_ALL_APPLICATIONS ? 'open' : ''}`}>
                <h2 className="jupiter-accordion-title">
                  <a
                    href="#all"
                    onClick={(e) => {
                      e.preventDefault();
                      setOpenAccordion(setAccordionTransition(openAccordion, ACCORDION_ALL_APPLICATIONS));
                    }}
                  >
                    All applications
                  </a>
                </h2>
                <div className="jupiter-accordion-body">
                  <Row>
                    <Columns className="small-offset-1 small-10 end">
                      <Row className="applications">
                        <AllApplications
                          updateClient={(id, data, callback) => updateClient(token, id, data, callback)}
                          clients={clients}
                          loading={loading}
                        />
                      </Row>
                    </Columns>
                  </Row>
                </div>
              </div>
            </AdminContent>
          </div>
        </Columns>
      </Row>
      <AdminContent user={user}>
        <div className="applications-divisor">
          <h2 className="applications-divisor-title">
            <span>New application</span>
          </h2>
          <Row>
            <Columns className="small-12">
              <NewApplication
                createClient={(data, callback) => createClient(token, data, callback)}
                postSubmit={() => setOpenAccordion(ACCORDION_ALL_APPLICATIONS)}
              />
            </Columns>
          </Row>
        </div>
      </AdminContent>
    </div>
  );
};

const mapStateToProps = state => {
  return {
    loading: state.internal.loading,
    token: state.internal.data.action_token,
    userId: state.internal.data.user_id,
    user: state.users.user,
    userClients: state.users.clients,
    clients: state.clients.clients
  };
};

const mapDispatchToProps = dispatch => {
  return {
    fetchActiveClients: (token, userId) => dispatch(fetchActiveClients(token, userId)),
    revokeActiveClient: (token, userId, clientId) => dispatch(revokeActiveClient(token, userId, clientId)),
    fetchClients: (token) => dispatch(fetchClients(token)),
    createClient: (token, data, callback) => dispatch(createClient(token, data, callback)),
    updateClient: (token, id, data, callback) => dispatch(updateClient(token, id, data, callback))
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(applications);
