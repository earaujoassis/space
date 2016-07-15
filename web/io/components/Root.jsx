import React from 'react'

import UserStore from '../stores/users'
import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

import SignUp from './SignUp.jsx'
import Success from './Success.jsx'

export default class Root extends React.Component {
    constructor() {
        super()
        this.state = UserStore.getState().payload || {}
        this._updateFromStore = this._updateFromStore.bind(this)
    }

    componentDidMount() {
        UserStore.addChangeListener(this._updateFromStore)
    }

    componentWillUnmount() {
        UserStore.removeChangeListener(this._updateFromStore)
    }

    render() {
        if (!!this.state.recover_secret && !!this.state.code_secret_image) {
            return (<Success
                codeSecretImage={this.state.code_secret_image}
                recoverSecret={this.state.recover_secret} />)
        } else {
            return (<SignUp />)
        }
    }

    _updateFromStore() {
        if (UserStore.success()) {
            this.setState(UserStore.getState().payload || {})
        }
    }
}
