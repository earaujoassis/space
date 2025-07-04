import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import EmailVerification from './emailVerification'
import Sessions from './sessions'

import './style.css'

const personal = ({
  requestEmailVerification,
  fetchApplicationSessionsForUser,
  revokeApplicationSessionForUser,
  loading,
  application,
  user,
  sessions,
}) => {
  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})
  let content = null

  useEffect(() => {
    if (!loading.includes('session') && sessions === undefined) {
      fetchApplicationSessionsForUser(application.user_id)
    } else {
      setProtectedResource({ ...protectedResource, sessions })
    }
  }, [sessions])

  useEffect(() => {
    if (
      !loading.includes('session') &&
      sessions !== undefined &&
      pendingFirstRender === true
    ) {
      setPendingFirstRender(false)
    }
  }, [loading, sessions])

  if (pendingFirstRender) {
    content = <SpinningSquare />
  } else if (user && !user.error) {
    content = (
      <div className="globals__siblings">
        <div className="globals__children">
          <div className="globals__input-wrapper">
            <label htmlFor="personal__full-name">Full name</label>
            <input
              className="read-only"
              disabled
              id="personal__full-name"
              value={`${user.first_name} ${user.last_name}`}
              type="text"
            />
          </div>
          <div className="globals__input-wrapper">
            <label htmlFor="personal__username">Username</label>
            <input
              className="read-only"
              disabled
              id="personal__username"
              value={user.username}
              type="text"
            />
          </div>
          <div className="globals__input-wrapper">
            <label htmlFor="personal__email">Primary email</label>
            <input
              className="read-only"
              disabled
              id="personal__email"
              value={user.email}
              type="text"
            />
          </div>
          <div className="globals__input-wrapper">
            <label htmlFor="personal__role">Role</label>
            <input
              className="read-only"
              disabled
              id="personal__role"
              value={user.is_admin ? 'Administrator' : 'Member'}
              type="text"
            />
          </div>
          <div className="globals__input-wrapper">
            <label htmlFor="personal__timezone">Timezone</label>
            <input
              className="read-only"
              disabled
              id="personal__timezone"
              value={user.timezone_identifier}
              type="text"
            />
          </div>
        </div>
        <div className="globals__children">
          <EmailVerification
            emailVerified={user.email_verified}
            requestEmailVerification={() =>
              requestEmailVerification(user.email, user.email)
            }
          />
          <Sessions
            sessions={protectedResource.sessions}
            revokeApplicationSessionForUser={id =>
              revokeApplicationSessionForUser(application.user_id, id)
            }
          />
        </div>
      </div>
    )
  }

  return (
    <>
      <h2>Personal information</h2>
      <div className="personal-root">{content}</div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    application: state.root.application,
    user: state.root.user,
    sessions: state.root.sessions,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    requestEmailVerification: (holder, email) =>
      dispatch(actions.requestEmailVerification(holder, email)),
    fetchApplicationSessionsForUser: id =>
      dispatch(actions.fetchApplicationSessionsForUser(id)),
    revokeApplicationSessionForUser: (userId, sessionId) =>
      dispatch(actions.revokeApplicationSessionForUser(userId, sessionId)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(personal)
