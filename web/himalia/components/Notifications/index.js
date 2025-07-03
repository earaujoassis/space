import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NotificationEmail from './notificationEmail'
import NotificationSettings from './notificationSettings'

const notifications = ({
  fetchUserProfile,
  fetchUserSettings,
  patchUserSettings,
  fetchEmails,
  loading,
  application,
  emails,
  settings,
  user,
}) => {
  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})
  let content = null

  useEffect(() => {
    fetchUserProfile(application.user_id, application.action_token)
    fetchEmails(application.action_token)
    fetchUserSettings(application.action_token)
  }, [])

  useEffect(() => {
    if (emails === undefined) {
      fetchEmails(application.action_token)
    } else {
      setProtectedResource({ ...protectedResource, emails })
    }
  }, [emails])

  useEffect(() => {
    if (settings === undefined) {
      fetchUserSettings(application.action_token)
    } else {
      setProtectedResource({ ...protectedResource, settings })
    }
  }, [settings])

  useEffect(() => {
    if (
      !loading.includes('user') &&
      !loading.includes('email') &&
      !loading.includes('setting') &&
      user !== undefined &&
      emails !== undefined &&
      settings !== undefined &&
      pendingFirstRender === true
    ) {
      setPendingFirstRender(false)
    }
  }, [loading, user, emails, settings])

  if (pendingFirstRender) {
    content = <SpinningSquare />
  } else if (protectedResource.emails) {
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
          patchUserSettings={data =>
            patchUserSettings(application.action_token, data)
          }
        />
        <NotificationSettings
          settings={protectedResource.settings}
          patchUserSettings={data =>
            patchUserSettings(application.action_token, data)
          }
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
    application: state.root.application,
    user: state.root.user,
    emails: state.root.emails,
    settings: state.root.settings,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchUserProfile: (id, token) =>
      dispatch(actions.fetchUserProfile(id, token)),
    fetchEmails: token => dispatch(actions.fetchEmails(token)),
    fetchUserSettings: token => dispatch(actions.fetchUserSettings(token)),
    patchUserSettings: (token, data) =>
      dispatch(actions.patchUserSettings(token, data)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(notifications)
