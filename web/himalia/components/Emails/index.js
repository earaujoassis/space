import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { fetchEmails, requestEmailVerification, addEmail } from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NewEmail from './newEmail'
import Emails from './emails'

const personal = () => {
  const loading = useSelector(state => state.root.loading)
  const user = useSelector(state => state.root.user)
  const emails = useSelector(state => state.root.emails)

  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})
  const [requestedVerification, setRequestedVerification] = useState([])

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    dispatch(fetchEmails())
  }, [])

  useEffect(() => {
    if (!loading.includes('email') && emails === undefined) {
      dispatch(fetchEmails())
    } else {
      setProtectedResource({ ...protectedResource, emails })
    }
  }, [emails])

  useEffect(() => {
    if (
      !loading.includes('email') &&
      emails !== undefined &&
      pendingFirstRender === true
    ) {
      setPendingFirstRender(false)
      setProtectedResource({ ...protectedResource, emails })
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
    const emailsComplete = [primaryEmail, ...protectedResource.emails]
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

export default personal
