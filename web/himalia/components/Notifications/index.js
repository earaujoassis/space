import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { fetchUserSettings, fetchEmails, patchUserSettings } from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NotificationEmail from './notificationEmail'
import NotificationSettings from './notificationSettings'

const notifications = () => {
  const loading = useSelector(state => state.root.loading)
  const user = useSelector(state => state.root.user)
  const emails = useSelector(state => state.root.emails)
  const settings = useSelector(state => state.root.settings)

  const [pendingFirstRender, setPendingFirstRender] = useState(true)
  const [protectedResource, setProtectedResource] = useState({})

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    dispatch(fetchEmails())
    dispatch(fetchUserSettings())
  }, [])

  useEffect(() => {
    if (!loading.includes('setting') && settings === undefined) {
      dispatch(fetchUserSettings())
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
          patchUserSettings={data => dispatch(patchUserSettings(data))}
        />
        <NotificationSettings
          settings={protectedResource.settings}
          patchUserSettings={data => dispatch(patchUserSettings(data))}
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

export default notifications
