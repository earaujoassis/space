import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { fetchUserSettings, fetchEmails, patchUserSettings } from '@actions'
import { useProtectedResource } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

import NotificationEmail from './notificationEmail'
import NotificationSettings from './notificationSettings'

const notifications = () => {
  const user = useSelector(state => state.user.data)
  const { data: emails, loading: loadingEmails } = useProtectedResource(
    'emails',
    fetchEmails
  )
  const { data: settings, loading: loadingSettings } = useProtectedResource(
    'settings',
    fetchUserSettings
  )

  const [pendingFirstRender, setPendingFirstRender] = useState(true)

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    if (
      !loadingEmails &&
      !loadingSettings &&
      emails !== undefined &&
      settings !== undefined &&
      user !== undefined
    ) {
      setPendingFirstRender(false)
    }
  }, [loadingEmails, loadingSettings, emails, settings])

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
          System notifications are related to actions performed through the
          application, using your credentials and sessions. Client notifications
          are related to actions performed by client applications, using your
          authorization. Sensitive e-mail messages, like password-less
          authentication through magic-links, are sent to the primary e-mail
          only.
        </p>
        <NotificationEmail
          selectedEmail={
            settings['notifications.system-email-notifications.email-address']
          }
          emails={emailsComplete}
          patchUserSettings={data => dispatch(patchUserSettings(data))}
        />
        <NotificationSettings
          settings={settings}
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
