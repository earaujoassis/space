import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NotificationEmail from './notificationEmail'
import NotificationSettings from './notificationSettings'

const notifications = ({
  fetchUserSettings,
  patchUserSettings,
  fetchEmails,
  loading,
  emails,
  settings,
  user,
}) => {
  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})
  let content = null

  useEffect(() => {
    fetchEmails()
    fetchUserSettings()
  }, [])

  useEffect(() => {
    if (!loading.includes('email') && emails === undefined) {
      fetchEmails()
    } else {
      setProtectedResource({ ...protectedResource, emails })
    }
  }, [emails])

  useEffect(() => {
    if (!loading.includes('setting') && settings === undefined) {
      fetchUserSettings()
    } else {
      setProtectedResource({ ...protectedResource, settings })
    }
  }, [settings])

  useEffect(() => {
    if (
      !loading.includes('email') &&
      !loading.includes('setting') &&
      emails !== undefined &&
      settings !== undefined &&
      pendingFirstRender === true
    ) {
      setPendingFirstRender(false)
      setProtectedResource({ ...protectedResource, emails, settings })
    }
  }, [loading, emails, settings])

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
          System notifications are related to actions performed through the
          application, using your credentials and sessions. Client notifications
          are related to actions performed by client applications, using your
          authorization. Sensitive e-mail messages, like password-less
          authentication through magic-links, are sent to the primary e-mail
          only.
        </p>
        <NotificationEmail
          selectedEmail={
            protectedResource.settings[
              'notifications.system-email-notifications.email-address'
            ]
          }
          emails={emailsComplete}
          patchUserSettings={data => patchUserSettings(data)}
        />
        <NotificationSettings
          settings={protectedResource.settings}
          patchUserSettings={data => patchUserSettings(data)}
        />
      </>
    )
  }

  return (
    <>
      <h2>Notifications</h2>
      <div className="notifications-root">{content}</div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    user: state.root.user,
    emails: state.root.emails,
    settings: state.root.settings,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchEmails: () => dispatch(actions.fetchEmails()),
    fetchUserSettings: () => dispatch(actions.fetchUserSettings()),
    patchUserSettings: data => dispatch(actions.patchUserSettings(data)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(notifications)
