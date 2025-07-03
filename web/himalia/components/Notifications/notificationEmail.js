import React from 'react'
import { Link } from 'react-router-dom'

const notificationEmail = ({ patchUserSettings, emails, selectedEmail }) => {
  if (!emails) {
    return null
  }

  const changeNotificationEmail = address => {
    const data = new FormData()
    data.append('key', 'notifications.system-email-notifications.email-address')
    data.append('value', address)
    patchUserSettings(data)
  }

  return (
    <>
      <div className="globals__scope">
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">
            Default notifications email
          </h4>
          <p>
            Emails can added through the <Link to="/emails">Emails</Link>{' '}
            settings. The selected email is used for all notifications and only
            verified emails can be chosen (except for the primary email). By
            default, the notifier uses the account&#39;s primary e-mail for
            notifications.
          </p>
          <label className="max-width">
            Select email
            <select
              value={selectedEmail}
              name="system-email-notifications__email-address"
              onChange={(event) => changeNotificationEmail(event.target.value)}
            >
              {emails.map(entry => (
                <option key={entry.id || 'primary-email'} value={entry.address}>
                  {entry.address}
                </option>
              ))}
            </select>
          </label>
        </div>
      </div>
    </>
  )
}

export default notificationEmail
