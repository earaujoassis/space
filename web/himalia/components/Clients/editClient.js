import React, { useState, useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import { updateClient } from '@actions'
import { useClientCleanup } from '@hooks'

import DynamicList from '@ui/DynamicList'

import Submenu from './submenu'

const editClient = () => {
  useClientCleanup()
  const loading = useSelector(state => state.clients.loading)
  const error = useSelector(state => state.clients.error)
  const clients = useSelector(state => state.clients.data)

  const dispatch = useDispatch()
  const navigate = useNavigate()

  const client = clients && clients.length ? clients[0] : null

  const [formSent, setFormSent] = useState(false)
  const [canonicalUri, setCanonicalUri] = useState(
    client ? client.uri.split('\n') : new Array()
  )
  const [redirectUri, setRedirectUri] = useState(
    client ? client.redirect.split('\n') : new Array()
  )

  let content = null

  useEffect(() => {
    if (!clients || !clients.length || !client) {
      navigate('/clients')
    } else if (formSent && !loading && !error) {
      navigate('/clients')
    } else if (formSent && !loading && error) {
      setFormSent(false)
    }
  }, [loading, error])

  useEffect(() => {
    if (client) {
      setCanonicalUri(client.uri.split('\n'))
      setRedirectUri(client.redirect.split('\n'))
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
          data.append('canonical_uri', canonicalUri.join('\n'))
          data.append('redirect_uri', redirectUri.join('\n'))
          dispatch(updateClient(client.id, data))
          setFormSent(true)
        }}
      >
        <div className="globals__siblings">
          <div className="globals__input-wrapper">
            <label htmlFor="new-client__name">Name</label>
            <input
              className="read-only"
              disabled
              id="new-client__name"
              defaultValue={client.name}
              type="text"
            />
          </div>
        </div>
        <div className="globals__siblings">
          <div className="globals__input-wrapper">
            <label htmlFor="new-client__description">Description</label>
            <input
              className="read-only"
              disabled
              id="new-client__description"
              defaultValue={client.description}
              type="text"
            />
          </div>
        </div>
        <DynamicList
          onChange={list => setCanonicalUri(Array.from(list))}
          defaultList={canonicalUri}
          label="Canonical URI"
          labelPlural="Canonical URIs"
          removeTitle="Remove canonical URI"
          id="canonical-uri"
          tabIndex="1"
        />
        <DynamicList
          onChange={list => setRedirectUri(Array.from(list))}
          defaultList={redirectUri}
          label="Redirect URI"
          labelPlural="Redirect URIs"
          removeTitle="Remove redirect URI"
          id="redirect-uri"
          tabIndex="2"
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
      <h2>Edit client application</h2>
      <Submenu activeAction="edit-client" editingClient />
      <div className="clients-root">{content}</div>
    </>
  )
}

export default editClient
