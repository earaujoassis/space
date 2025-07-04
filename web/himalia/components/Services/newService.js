import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { connect } from 'react-redux'

import * as actions from '@actions'
import { extractDataForm, prependUrlWithHttps } from '@utils/forms'

import Submenu from './submenu'

const newService = ({ createService, stateSignal }) => {
  const [formSent, setFormSent] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    if (stateSignal === 'service_record_success' && formSent) {
      navigate('/services')
    } else if (stateSignal === 'service_record_error' && formSent) {
      setFormSent(false)
    }
  }, [stateSignal])

  return (
    <>
      <h2>Create a new service application</h2>
      <Submenu activeAction="new-service" />
      <div className="services-root">
        <form
          className="form-common"
          action="."
          method="post"
          onSubmit={e => {
            e.preventDefault()
            const attrs = ['name', 'description', 'canonical_uri']
            const data = extractDataForm(e.target, attrs)
            createService(data)
            setFormSent(true)
          }}
        >
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-service__name">Name</label>
              <input
                autoFocus
                tabIndex="1"
                required
                autoComplete="off"
                id="new-service__name"
                name="name"
                type="text"
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-service__description">Description</label>
              <input
                tabIndex="2"
                autoComplete="off"
                id="new-service__description"
                name="description"
                type="text"
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <label htmlFor="new-service__canonical-uri">Canonical URI</label>
              <input
                tabIndex="3"
                required
                autoComplete="off"
                id="new-service__canonical-uri"
                name="canonical_uri"
                inputMode="url"
                type="url"
                onBlurCapture={e => prependUrlWithHttps(e)}
              />
            </div>
          </div>
          <div className="globals__siblings">
            <div className="globals__input-wrapper">
              <input
                tabIndex="4"
                type="submit"
                className="button submit"
                value="Create service application"
              />
            </div>
          </div>
        </form>
      </div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    stateSignal: state.root.stateSignal,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    createService: data => dispatch(actions.createService(data)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(newService)
