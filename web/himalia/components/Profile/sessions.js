import React from 'react'

const sessions = ({ revokeApplicationSessionForUser, sessions }) => {
  if (!sessions) {
    return null
  }

  return (
    <>
      <h3 className="globals__subheader">Active sessions</h3>
      <ul className="profile_session__list">
        {sessions.map((entry) => {
          return (
            <li key={entry.id}>
              <div className="profile_session__entry">
                <p>IP: {entry.ip}</p>
                <p>User-Agent: {entry.user_agent}</p>
                <p className="profile_session__entry-subsection">
                  <span>
                    <button
                      onClick={() => revokeApplicationSessionForUser(entry.id)}
                      className="button-anchor"
                    >
                      Revoke session
                    </button>
                  </span>
                  <span>
                    {entry.current ? (
                      <span className="profile_session__current">
                        current session
                      </span>
                    ) : null}
                  </span>
                </p>
              </div>
            </li>
          )
        })}
      </ul>
    </>
  )
}

export default sessions
