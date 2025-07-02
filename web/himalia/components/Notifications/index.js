import React, { useEffect } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NotificationEmail from './notificationEmail'
import NotificationSettings from './notificationSettings'

const notifications = ({
  fetchUserProfile,
  fetchEmails,
  loading,
  application,
  emails,
  user
}) => {
  let content = null

  useEffect(() => {
    fetchUserProfile(application.user_id, application.action_token)
    fetchEmails(application.action_token)
  }, [])

  useEffect(() => {
    if (emails === undefined) {
      fetchEmails(application.action_token)
    }
  }, [emails])

  if (
      loading.includes('email') ||
      loading.includes('user') ||
      emails === undefined ||
      user === undefined
    ) {
      content = <SpinningSquare />
    } else if (emails) {
      const primaryEmail = {
        verified: user.email_verified,
        address: user.email,
        primary: true,
      }
      const emailsComplete = [primaryEmail, ...emails]
      content = (
        <>
          <p>
            System notifications are related to actions performed through
            the application, using your credentials and sessions.
            Client notifications are related to actions performed by
            client applications, using your authorization.
          </p>
          <NotificationEmail emails={emailsComplete} />
          <NotificationSettings />
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
    emails: state.root.emails,
    user: state.root.user,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchUserProfile: (id, token) =>
      dispatch(actions.fetchUserProfile(id, token)),
    fetchEmails: token => dispatch(actions.fetchEmails(token))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(notifications)
