import React from 'react'

import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

import UserStore from '../stores/users'
import UsersActions from '../actions/users'

export default class Applications extends React.Component {
    constructor() {
        super()
        this.state = {loading: true}
        this._updateFromStore = this._updateFromStore.bind(this)
        this._applications = this._applications.bind(this)
    }

    componentDidMount() {
        UserStore.addChangeListener(this._updateFromStore)
        UsersActions.fetchActiveClients()
    }

    componentWillUnmount() {
        UserStore.removeChangeListener(this._updateFromStore)
    }

    render() {
        if (this.state.loading) {
            return (
                <Row>
                    <Columns className="small-offset-1 small-10 end">
                        <p className="text-center">Loading...</p>
                    </Columns>
                </Row>
            )
        }
        return (
            <Row>
                <Columns className="small-offset-1 small-10 end">
                    <Row className="applications">
                        {this._applications()}
                    </Row>
                </Columns>
            </Row>
        )
    }

    _applications() {
        if (this.state.loading) {
            return []
        }

        if (!this.state.clients.length) {
            return (<p className="blank-list">No applications available yet.</p>)
        }

        let applications = []
        for (var i = 0; i < this.state.clients.length; i++) {
            let client = this.state.clients[i]
            applications.push(
                <Columns className="small-12" key={i}>
                    <div className="application-card">
                        <p className="title">{client.name} <small>(<a href={client.uri} rel="noopener noreferrer" target="_blank">{client.uri.split(/:\/\//)[1]}</a>)</small></p>
                        <p className="action"><button className="button" onClick={this._revokeAccess.bind(client)}>Revoke Access</button></p>
                        <p className="scope">{client.description}</p>
                        <p className="last-access"><em>Last access:</em> ?</p>
                    </div>
                </Columns>
            )
        }
        return applications
    }

    _revokeAccess(e) {
        e.preventDefault()
        UsersActions.revokeActiveClient(this.id)
    }

    _updateFromStore() {
        if (UserStore.success()) {
            let state = Object.assign({}, UserStore.getState().payload || {}, {loading: false})
            this.setState(state)
        }
    }
}
