import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { connect } from 'react-redux'

import * as actions from '@actions'
import { extractDataForm, prependUrlWithHttps } from '@utils/forms'

import Submenu from './submenu'

const termsLink = (
    <a href="//quatrolabs.com/terms-of-service">terms of service</a>
)
const privacyPolicyLink = (
    <a href="//quatrolabs.com/privacy-policy">privacy policy</a>
)

const newClient = ({ createClient, application, stateSignal }) => {
    const [formSent, setFormSent] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        if (stateSignal === 'client_record_success' && formSent) {
            navigate('/clients')
        } else if (stateSignal === 'client_record_error' && formSent) {
            setFormSent(false)
        }
    }, [stateSignal])

    return (
        <>
            <h2>Create a new client application</h2>
            <Submenu activeAction="new-client" />
            <p>
                By clicking &quot;Create client application&quot;, you agree to our {termsLink} and {privacyPolicyLink}.
                Also, you guarantee that the corresponding client application will adhere to those terms
                and policies, while handling user data.
            </p>
            <div className="clients-root">
                <form className="form-common" action="." method="post" onSubmit={(e) => {
                    e.preventDefault()
                    const attrs = ['name', 'description', 'canonical_uri', 'redirect_uri']
                    const data = extractDataForm(e.target, attrs)
                    createClient(data, application.action_token)
                    setFormSent(true)
                }}>
                    <div className="globals__siblings">
                        <div className="globals__input-wrapper">
                            <label htmlFor="new-client__name">Name</label>
                            <input autoFocus tabIndex="1" required autoComplete="off" id="new-client__name" name="name" type="text" />
                        </div>
                    </div>
                    <div className="globals__siblings">
                        <div className="globals__input-wrapper">
                            <label htmlFor="new-client__description">Description</label>
                            <input tabIndex="2" required autoComplete="off" id="new-client__description" name="description" type="text" />
                        </div>
                    </div>
                    <div className="globals__siblings">
                        <div className="globals__input-wrapper">
                            <label htmlFor="new-client__canonical-uri">Canonical URI</label>
                            <input tabIndex="3" required autoComplete="off" id="new-client__canonical-uri" name="canonical_uri" inputMode="url" type="url" onBlurCapture={(e) => prependUrlWithHttps(e)} />
                        </div>
                    </div>
                    <div className="globals__siblings">
                        <div className="globals__input-wrapper">
                            <label htmlFor="new-client__redirect-uri">Redirect URI</label>
                            <input tabIndex="4" required autoComplete="off" id="new-client__redirect-uri" name="redirect_uri" inputMode="url" type="url" onBlurCapture={(e) => prependUrlWithHttps(e)} />
                        </div>
                    </div>
                    <div className="globals__siblings">
                        <div className="globals__input-wrapper">
                            <input tabIndex="5" type="submit" className="button" value="Create client application" />
                        </div>
                    </div>
                </form>
            </div>
        </>
    )
}

const mapStateToProps = state => {
    return {
        application: state.root.application,
        stateSignal: state.root.stateSignal
    }
}

const mapDispatchToProps = dispatch => {
    return {
        createClient: (data, token) => dispatch(actions.createClient(data, token))
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(newClient)
