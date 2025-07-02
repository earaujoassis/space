import React from 'react'

import Switch from '@ui/Switch'

const notificationSettings = () => {
  return (
    <>
      <div className="globals__scope">
        <h3 className="globals__scope-header">System notifications</h3>
        <div className="globals__scope-corpus">
          <h4 className="globals__scope-subheader">Sign-in</h4>
          <p>Notification for every sign-in using the account credentials.</p>
          <Switch
            id="system-email-notifications__signin"
            name="system-email-notifications__signin"
            label="System Notification - Sign-in"
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
          />
        </div>
      </div>
    </>
  )
}

export default notificationSettings
