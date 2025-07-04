import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import { updateClient } from '@actions'

import Submenu from './submenu'
import ScopesGroup from '@ui/ScopesGroup'

const editScopes = () => {
  const stateSignal = useSelector(state => state.root.stateSignal)
  const clients = useSelector(state => state.root.clients)

  const client = clients && clients.length ? clients[0] : null
  let content = null

  const [formSent, setFormSent] = useState(false)
  const [scopes, setScopes] = useState([])

  const navigate = useNavigate()
  const dispatch = useDispatch()

  useEffect(() => {
    if (!clients || !clients.length || clients.error || !client) {
      navigate('/clients')
    } else if (stateSignal === 'client_record_success' && formSent) {
      navigate('/clients')
    } else if (stateSignal === 'client_record_error' && formSent) {
      setFormSent(false)
    }
  }, [stateSignal])

  useEffect(() => {
    if (client) {
      setScopes(client.scopes.split(' '))
    }
  }, [clients])

  if (client) {
    content = (
      <form
        className="form-common"
        action="."
        method="post"
        onSubmit={e => {
          e.preventDefault()
          const data = new FormData()
          data.append('scopes', scopes.join(' '))
          dispatch(updateClient(client.id, data))
          setFormSent(true)
        }}
      >
        <p>
          By default, all applications are created with &quot;
          <code>public</code>&quot; scope, which can&#39;t introspect user data
          (read user&#39;s full name, email etc.), nor interact with the OIDC
          Provider endpoints.
        </p>
        <div className="globals__siblings">
          <div className="globals__input-wrapper">
            <input
              className="read-only"
              disabled
              id="new-client__name"
              defaultValue={client.name}
              type="text"
            />
          </div>
        </div>
        <ScopesGroup
          initialScopes={scopes}
          onChange={scopes => setScopes(scopes)}
        />
        <div className="globals__siblings globals__form-actions">
          <div className="globals__input-wrapper">
            <input
              disabled={formSent}
              tabIndex="3"
              type="submit"
              className="button submit"
              value="Save client application"
            />
            <button
              tabIndex="4"
              onClick={e => {
                e.preventDefault()
                navigate('/clients')
              }}
              className="submit cancel"
            >
              Cancel
            </button>
          </div>
        </div>
      </form>
    )
  }

  return (
    <>
      <h2>Select client application scopes</h2>
      <Submenu activeAction="edit-scopes" editingScopes />
      <div className="clients-root">{content}</div>
    </>
  )
}

export default editScopes
