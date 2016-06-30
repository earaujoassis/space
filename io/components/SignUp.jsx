import React from 'react'
import ReactDOM from 'react-dom'

import UsersActions from '../actions/users'
import UserStore from '../stores/users'
import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

export default class SignUp extends React.Component {
    constructor() {
        super()
        this.state = {}
        this._signUp = this._signUp.bind(this)
        this._updateFromStore = this._updateFromStore.bind(this)
    }

    static get propTypes() {
        return {
            afterSignup: React.PropTypes.func
        }
    }

    componentDidMount() {
        UserStore.addChangeListener(this._updateFromStore);
    }

    componentWillUnmount() {
        UserStore.removeChangeListener(this._updateFromStore);
    }

    render() {
        return (
            <div className="signup-content">
                <Row>
                    <Columns className="large-offset-1 large-5 description">
                        <h2 className="title">Create a new account</h2>
                        <p>Space is an user management microservice. We aim to provide a secure and reliable authentication and authorization system.</p>
                        <p>By clicking "Sign Up", you agree to our <a href="//quatrolabs.com/terms-of-service">terms of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. We will send you account related emails occasionally.</p>
                    </Columns>
                    <Columns className="large-5 end">
                        <form action="." method="post" ref="form" onSubmit={this._signUp}>
                            <Row>
                                <Columns className="large-6">
                                    <input type="text" name="firstName" placeholder="First Name" />
                                </Columns>
                                <Columns className="large-6">
                                    <input type="text" name="lastName" placeholder="Last Name" />
                                </Columns>
                            </Row>
                            <input type="hidden" name="actionToken" value="" />
                            <input type="text" name="username" placeholder="Username" />
                            <input type="email" name="email" placeholder="Email" />
                            <input type="password" name="password" placeholder="Password" />
                            <button type="submit" className="button expand" onClick={this._signUp}>Sign Up</button>
                        </form>
                    </Columns>
                </Row>
            </div>
        );
    }

    _signUp(e) {
        e.preventDefault()
        UsersActions.signIn(new FormData(ReactDOM.findDOMNode(this.refs.form)))
    }

    _updateFromStore() {
        if (UserStore.success()) {
            window.location = '/success'
        }
    }
}
