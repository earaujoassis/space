import React from 'react'

import Switch from '@ui/Switch'

const notificationSettings = ({ patchUserSettings, settings }) => {
  const checked = key => {
    return settings[key] === true
  }

  const changeSettings = key => {
    const isChecked = checked(key)
    const data = new FormData()
    data.append('key', key)
    data.append('value', !isChecked)
    patchUserSettings(data)
  }

  return (
    <>
      <div className="globals__scope">
        <h3 className="globals__scope-header">System notifications</h3>
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">Authentication</h4>
          <p>
            Notification for every authentication using the account credentials.
          </p>
          <Switch
            id="system-email-notifications__authentication"
            name="system-email-notifications__authentication"
            label="System Notification - authentication"
            checked={checked(
              'notifications.system-email-notifications.authentication'
            )}
            onChange={() =>
              changeSettings(
                'notifications.system-email-notifications.authentication'
              )
            }
          />
        </div>
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">
            Client application authorization
          </h4>
          <p>Notification for every new client application authorization.</p>
          <Switch
            id="system-email-notifications__client-authorization"
            name="system-email-notifications__client-authorization"
            label="System Notification - Client authorization"
            checked={checked(
              'notifications.system-email-notifications.client-authorization'
            )}
            onChange={() =>
              changeSettings(
                'notifications.system-email-notifications.client-authorization'
              )
            }
          />
        </div>
      </div>
      <div className="globals__scope">
        <h3 className="globals__scope-header">
          Client application notifications
        </h3>
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">Token introspection</h4>
          <p>
            Notification for every token introspection performed by client
            applications.
          </p>
          <Switch
            id="client-application-email-notifications__token-introspection"
            name="client-application-email-notifications__token-introspection"
            label="Client Application Notification - Introspection"
            checked={checked(
              'notifications.client-application-email-notifications.token-introspection'
            )}
            onChange={() =>
              changeSettings(
                'notifications.client-application-email-notifications.token-introspection'
              )
            }
          />
        </div>
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">Userinfo introspection</h4>
          <p>
            Notification for every userinfo introspection performed by client
            applications.
          </p>
          <Switch
            id="client-application-email-notifications__userinfo-introspection"
            name="client-application-email-notifications__userinfo-introspection"
            label="Client Application Notification - Userinfo Introspection"
            checked={checked(
              'notifications.client-application-email-notifications.userinfo-introspection'
            )}
            onChange={() =>
              changeSettings(
                'notifications.client-application-email-notifications.userinfo-introspection'
              )
            }
          />
        </div>
      </div>
    </>
  )
}

export default notificationSettings
