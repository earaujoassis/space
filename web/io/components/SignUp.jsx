import React from 'react'

import Row from '@core/components/Row.jsx'
import Columns from '@core/components/Columns.jsx'
import PasswordInput from '@core/components/PasswordInput.jsx'
import { extractDataForm } from '@core/utils/forms'

import { useApp } from '../context/useApp'

const signUp = ({ validationFailed }) => {
    const { actions } = useApp()

    const handleSubmit = (e) => {
        e.preventDefault()
        const attrs = ['first_name', 'last_name', 'action_token', 'username', 'email', 'password']
        actions.createUser(extractDataForm(e.target, attrs))
    }

    return (
        <div className="signup-content">
            <Row>
                <Columns className="small-offset-1 small-5 description">
                    <h2 className="title">Create a new account</h2>
                    <p>
                        Space is an user management microservice. We aim to provide a secure and reliable authentication
                        and authorization system.
                    </p>
                    <p>
                        By clicking &quot;Sign Up&quot;, you agree to our <a href="//quatrolabs.com/terms-of-service">terms
                        of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. We will send you
                        account related emails occasionally.
                    </p>
                </Columns>
                <Columns className="small-5 end">
                    <form
                        className="form-common"
                        action="."
                        method="post"
                        onSubmit={(e) => handleSubmit(e)}
                    >
                        {
                            validationFailed ? (
                                <p className="error-message">Validation failed</p>
                            ) : null
                        }
                        <Row>
                            <Columns className="small-6">
                                <input type="text" name="first_name" placeholder="First Name" required />
                            </Columns>
                            <Columns className="small-6">
                                <input type="text" name="last_name" placeholder="Last Name" required />
                            </Columns>
                        </Row>
                        <input type="hidden" name="action_token" value="" />
                        <input type="text" name="username" placeholder="Username" autoComplete="username" required />
                        <input type="email" name="email" placeholder="Email" inputMode="email" required />
                        <PasswordInput />
                        <button type="submit" className="button expand">Sign Up</button>
                    </form>
                </Columns>
            </Row>
        </div>
    )
}

export default signUp
