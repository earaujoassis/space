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
  const { user_id } = application
  let content = null

  useEffect(() => {
    if (!loading.includes('client') && clients === undefined) {
      fetchClientApplicationsFromUser(user_id)
    }
  }, [clients])

  useEffect(() => {
    fetchClientApplicationsFromUser(user_id)
  }, [])

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
                        revokeClientApplicationFromUser(user_id, client.id)
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
    fetchClientApplicationsFromUser: id =>
      dispatch(actions.fetchClientApplicationsFromUser(id)),
    revokeClientApplicationFromUser: (userId, clientId) =>
      dispatch(actions.revokeClientApplicationFromUser(userId, clientId)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(applications)
