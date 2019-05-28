import React, { useState } from 'react'

import UsersActions from '../actions/users'
import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

import { extractDataForm } from '../../core/utils/forms'

const hideImg = '/public/imgs/eye-open.png'
const displayImg = '/public/imgs/eye-blocked.png'

const signUp = ({ validationFailed }) => {
    const [showPassword, setShowPassword] = useState(true)

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
                        onSubmit={(e) => {
                            e.preventDefault()
                            const attrs = [ 'first_name', 'last_name', 'action_token', 'username', 'email', 'password' ]
                            UsersActions.signUp(extractDataForm(e.target, attrs))
                        }}>
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
                        <input type="text" name="username" placeholder="Username" required />
                        <input type="email" name="email" placeholder="Email" required />
                        <div className="password-visibility">
                            <input type={showPassword ? 'password' : 'text'} name="password" placeholder="Password" required />
                            <button
                                className="visibility-toggle"
                                onClick={(e) => {
                                    e.preventDefault()
                                    setShowPassword(!showPassword)
                                }}>
                                <img
                                    src={showPassword ? displayImg : hideImg}
                                    width="20"
                                    title="Toggle password visibility"
                                    alt="" />
                            </button>
                        </div>
                        <button type="submit" className="button expand">Sign Up</button>
                    </form>
                </Columns>
            </Row>
        </div>
    )
}

export default signUp
