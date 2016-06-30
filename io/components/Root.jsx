import React from 'react'

import SignUp from './SignUp.jsx'

export default class Root extends React.Component {
    constructor() {
        super()
        this.state = { user: null }
        this._afterSignup = this._afterSignup.bind(this)
    }

    render() {
        let component
        if (!this.state.user) {
            component = <SignUp afterSignup={this._afterSignup} />
        } else {
            component = null
        }
        return component
    }

    _afterSignup(user) {
        this.setState({user: user})
    }
};
