import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { fetchEmails, requestEmailVerification, addEmail } from '@actions'
import { useProtectedResource } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NewEmail from './newEmail'
import Emails from './emails'

const emails = () => {
  const user = useSelector(state => state.user.data)
  const { data: emails, loading } = useProtectedResource('emails', fetchEmails)

  const [pendingFirstRender, setPendingFirstRender] = useState(
    loading || emails === undefined
  )
  const [requestedVerification, setRequestedVerification] = useState([])

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    if (!loading && emails !== undefined && user !== undefined) {
      setPendingFirstRender(false)
    }
  }, [loading, emails])

  if (pendingFirstRender) {
    content = <SpinningSquare />
  } else {
    const primaryEmail = {
      verified: user.email_verified,
      address: user.email,
      primary: true,
    }
    const emailsComplete = [primaryEmail, ...emails]
    content = (
      <>
        <p>
          Emails you can use to receive notifications. For authentication you
          can only use your primary e-mail.
        </p>
        <Emails
          emails={emailsComplete}
          requestedVerification={requestedVerification}
          setRequestedVerification={setRequestedVerification}
          requestEmailVerification={email =>
            dispatch(requestEmailVerification(user.email, email))
          }
        />
        <NewEmail addEmail={data => dispatch(addEmail(data))} />
      </>
    )
  }

  return (
    <>
      <h2>Emails</h2>
      <div className="emails-root">{content}</div>
    </>
  )
}

export default emails
