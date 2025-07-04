import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import {
  fetchApplicationSessionsForUser,
  revokeApplicationSessionForUser,
  requestEmailVerification,
} from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import EmailVerification from './emailVerification'
import Sessions from './sessions'

import './style.css'

const personal = () => {
  const loading = useSelector(state => state.root.loading)
  const application = useSelector(state => state.root.application)
  const user = useSelector(state => state.root.user)
  const sessions = useSelector(state => state.root.sessions)

  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    if (!loading.includes('session') && sessions === undefined) {
      dispatch(fetchApplicationSessionsForUser(application.user_id))
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
              dispatch(requestEmailVerification(user.email, user.email))
            }
          />
          <Sessions
            sessions={protectedResource.sessions}
            revokeApplicationSessionForUser={id =>
              dispatch(revokeApplicationSessionForUser(application.user_id, id))
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

export default personal
