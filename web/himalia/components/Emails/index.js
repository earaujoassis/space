import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NewEmail from './newEmail'
import Emails from './emails'

const personal = ({
  fetchEmails,
  requestEmailVerification,
  addEmail,
  loading,
  emails,
  user,
}) => {
  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})
  const [requestedVerification, setRequestedVerification] = useState([])
  let content = null

  useEffect(() => {
    fetchEmails()
  }, [])

  useEffect(() => {
    if (!loading.includes('email') && emails === undefined) {
      fetchEmails()
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
            requestEmailVerification(user.email, email)
          }
        />
        <NewEmail addEmail={data => addEmail(data)} />
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

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    emails: state.root.emails,
    user: state.root.user,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchEmails: () => dispatch(actions.fetchEmails()),
    requestEmailVerification: (holder, email) =>
      dispatch(actions.requestEmailVerification(holder, email)),
    addEmail: data => dispatch(actions.addEmail(data)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(personal)
