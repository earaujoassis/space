import React from 'react'
import { useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import { fetchClients, setClientForEdition } from '@actions'
import { useProtectedResource, useClientCleanup } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import Submenu from './submenu'

const clients = () => {
  useClientCleanup()
  const { data: clients, loading } = useProtectedResource('clients', fetchClients)

  const dispatch = useDispatch()
  const navigate = useNavigate()

  let content = null

  if (loading || clients === undefined) {
    content = <SpinningSquare />
  } else if (clients && clients.length) {
    content = (
      <ul className="clients-root__list">
        {clients.map((client, i) => {
          const canonicalUri = client.uri.split('\n')[0]
          const canonicalUriShort = canonicalUri.split(/:\/\//)[1]
          return (
            <li key={i}>
              <div className="clients-root__entry">
                <h3>
                  {client.name}{' '}
                  <span>
                    (<a href={canonicalUri}>{canonicalUriShort}</a>)
                  </span>
                </h3>
                <p>{client.description}</p>
                <p>
                  <button
                    onClick={() => {
                      dispatch(setClientForEdition(client))
                      navigate('/clients/edit')
                    }}
                    className="button-anchor"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => {
                      dispatch(setClientForEdition(client))
                      navigate('/clients/edit/scopes')
                    }}
                    className="button-anchor"
                  >
                    Select Scopes
                  </button>
                  <a
                    href={`/api/clients/${client.id}/credentials`}
                    title="It regenerates the client's secret for security reasons"
                    rel="noopener noreferrer"
                    className="button-anchor"
                  >
                    Download credentials
                  </a>
                </p>
              </div>
            </li>
          )
        })}
      </ul>
    )
  }

  return (
    <>
      <h2>Clients</h2>
      <Submenu activeAction="all-clients" />
      <div className="clients-root">{content}</div>
    </>
  )
}

export default clients
