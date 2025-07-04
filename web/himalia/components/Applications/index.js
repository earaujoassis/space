import React from 'react'
import { useSelector, useDispatch } from 'react-redux'

import {
  fetchClientApplicationsFromUser,
  revokeClientApplicationFromUser,
} from '@actions'
import { useProtectedResource, useClientCleanup } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import './style.css'

const applications = () => {
  useClientCleanup()
  const { user_id } = useSelector(state => state.workspace.data)
  const { data: clients, loading } = useProtectedResource('clients', () =>
    fetchClientApplicationsFromUser(user_id)
  )

  const dispatch = useDispatch()

  let content = null

  if (loading || clients === undefined) {
    content = <SpinningSquare />
  } else if (clients && clients.length) {
    content = (
      <>
        <p>The following applications are associated with your user account.</p>
        <ul className="applications-root__list">
          {clients.map((client, i) => {
            const canonicalUri = client.uri.split('\n')[0]
            const canonicalUriShort = canonicalUri.split(/:\/\//)[1]
            return (
              <li key={i}>
                <div className="applications-root__entry">
                  <h3>
                    {client.name}{' '}
                    <span>
                      (<a href={canonicalUri}>{canonicalUriShort}</a>)
                    </span>
                  </h3>
                  <p>{client.description}</p>
                  <p>
                    <button
                      onClick={e => {
                        e.preventDefault()
                        dispatch(
                          revokeClientApplicationFromUser(user_id, client.id)
                        )
                      }}
                      className="button-anchor"
                    >
                      Revoke access
                    </button>
                  </p>
                </div>
              </li>
            )
          })}
        </ul>
      </>
    )
  }

  return (
    <>
      <h2>Applications</h2>
      <div className="applications-root">{content}</div>
    </>
  )
}

export default applications
