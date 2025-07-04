import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import {
  fetchApplicationSessionsForUser,
  revokeApplicationSessionForUser,
  requestEmailVerification,
} from '@actions'
import { useProtectedResource } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import EmailVerification from './emailVerification'
import Sessions from './sessions'

import './style.css'

const profile = () => {
  const workspace = useSelector(state => state.workspace.data)
  const { data: sessions, loading } = useProtectedResource('sessions', () =>
    fetchApplicationSessionsForUser(workspace.user_id)
  )
  const user = useSelector(state => state.user.data)

  const [pendingFirstRender, setPendingFirstRender] = useState(
    loading || sessions === undefined
  )

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    if (!loading && sessions !== undefined) {
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
            sessions={sessions}
            revokeApplicationSessionForUser={id =>
              dispatch(revokeApplicationSessionForUser(workspace.user_id, id))
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

export default profile
