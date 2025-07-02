import React, { useEffect } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

const applications = ({
  fetchClientApplicationsFromUser,
  revokeClientApplicationFromUser,
  loading,
  application,
  clients,
}) => {
  const { user_id, action_token } = application
  let content = null

  useEffect(() => {
    fetchClientApplicationsFromUser(user_id, action_token)
  }, [])

  useEffect(() => {
    if (clients === undefined) {
      fetchClientApplicationsFromUser(user_id, action_token)
    }
  }, [clients])

  if (loading.includes('client') || clients === undefined) {
    content = <SpinningSquare />
  } else if (clients && clients.length) {
    content = (
      <>
        <p>The following applications are associated with your user account.</p>
        <ul className="applications-root__list">
          {clients.map((client, i) => {
            const canonicalUri = client.uri.split('\n')[0]
            const canonicalUriShort = canonicalUri.split(/:\/\//)[1]
            return (
              <li key={i}>
                <div className="applications-root__entry">
                  <h3>
                    {client.name}{' '}
                    <span>
                      (<a href={canonicalUri}>{canonicalUriShort}</a>)
                    </span>
                  </h3>
                  <p>{client.description}</p>
                  <p>
                    <button
                      onClick={e => {
                        e.preventDefault()
                        revokeClientApplicationFromUser(
                          user_id,
                          client.id,
                          action_token
                        )
                      }}
                      className="button-anchor"
                    >
                      Revoke access
                    </button>
                  </p>
                </div>
              </li>
            )
          })}
        </ul>
      </>
    )
  }

  return (
    <>
      <h2>Applications</h2>
      <div className="applications-root">{content}</div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    application: state.root.application,
    clients: state.root.clients,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchClientApplicationsFromUser: (id, token) =>
      dispatch(actions.fetchClientApplicationsFromUser(id, token)),
    revokeClientApplicationFromUser: (userId, clientId, token) =>
      dispatch(
        actions.revokeClientApplicationFromUser(userId, clientId, token)
      ),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(applications)
