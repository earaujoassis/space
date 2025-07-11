import React, { useState, useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import { createClient } from '@actions'
import { useClientCleanup } from '@hooks'

import { extractDataForm, prependUrlWithHttps } from '@utils/forms'

import Submenu from './submenu'

const termsLink = (
  <a href="//quatrolabs.com/terms-of-service">terms of service</a>
)
const privacyPolicyLink = (
  <a href="//quatrolabs.com/privacy-policy">privacy policy</a>
)

const newClient = () => {
  useClientCleanup()
  const loading = useSelector(state => state.clients.loading)
  const error = useSelector(state => state.clients.error)

  const [formSent, setFormSent] = useState(false)
  const navigate = useNavigate()
  const dispatch = useDispatch()

  useEffect(() => {
    if (formSent && !loading && !error) {
      navigate('/clients')
    } else if (formSent && !loading && error) {
      setFormSent(false)
    }
  }, [loading, error])

  return (
    <>
      <h2>Create a new client application</h2>
      <Submenu activeAction="new-client" />
      <p>
        By clicking &quot;Create client application&quot;, you agree to our{' '}
        {termsLink} and {privacyPolicyLink}. Also, you guarantee that the
        corresponding client application will adhere to those terms and
        policies, while handling user data.
      </p>
      <div className="clients-root">
        <form
          className="form-common"
          action="."
          method="post"
          onSubmit={e => {
            e.preventDefault()
            const attrs = [
              'name',
              'description',
              'canonical_uri',
              'redirect_uri',
            ]
            const data = extractDataForm(e.target, attrs)
            dispatch(createClient(data))
            setFormSent(true)
          }}
        >
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-client__name">Name</label>
              <input
                autoFocus
                tabIndex="1"
                required
                autoComplete="off"
                id="new-client__name"
                name="name"
                type="text"
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-client__description">Description</label>
              <input
                tabIndex="2"
                required
                autoComplete="off"
                id="new-client__description"
                name="description"
                type="text"
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-client__canonical-uri">Canonical URI</label>
              <input
                tabIndex="3"
                required
                autoComplete="off"
                id="new-client__canonical-uri"
                name="canonical_uri"
                inputMode="url"
                type="url"
                onBlurCapture={e => prependUrlWithHttps(e)}
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-client__redirect-uri">Redirect URI</label>
              <input
                tabIndex="4"
                required
                autoComplete="off"
                id="new-client__redirect-uri"
                name="redirect_uri"
                inputMode="url"
                type="url"
                onBlurCapture={e => prependUrlWithHttps(e)}
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <input
                tabIndex="5"
                type="submit"
                className="button submit"
                value="Create client application"
              />
            </div>
          </div>
        </form>
      </div>
    </>
  )
}

export default newClient
