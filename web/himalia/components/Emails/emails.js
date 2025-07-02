import React from 'react'

const emailEntryVerified = (verified) => {
  if (verified) {
    return <span className="emails__entry-tag verified">verified</span>
  } else {
    return <span className="emails__entry-tag unverified">unverified</span>
  }
}

const emailEntryPrimary = (primary) => {
  if (primary) {
    return <span className="emails__entry-tag primary">primary</span>
  }

  return null
}

const emails = ({
  requestEmailVerification,
  setRequestedVerification,
  requestedVerification,
  emails,
}) => {
  if (!emails) {
    return null
  }

  return (
    <ul className="emails__list">
      {emails.map((entry) => (
        <li key={entry.id || 'primary-email'}>
          <div className="emails__entry">
            <p>
              <span className="emails__entry-address">{entry.address}</span>
              {emailEntryPrimary(entry.primary)}
              {emailEntryVerified(entry.verified)}
            </p>
            <p className="emails__entry-action">
              {entry.verified ||
              requestedVerification.includes(entry.id) ? null : (
                <button
                  onClick={(e) => {
                    e.preventDefault()
                    requestEmailVerification(entry.address)
                    setRequestedVerification([
                      entry.id,
                      ...requestedVerification,
                    ])
                  }}
                  className="button-anchor"
                >
                  Request e-mail verification
                </button>
              )}
              {requestedVerification.includes(entry.id) ? 'Requested' : null}
            </p>
          </div>
        </li>
      ))}
    </ul>
  )
}

export default emails
