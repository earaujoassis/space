import React from 'react'

import UserStore from '../../stores/users'
import UsersActions from '../../actions/users'

class MyApplications extends React.Component {
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
        return (
            <div className="applications-listing">
                {this.state.loading ? (
                    <p className="text-center">Loading...</p>
                ) : (
                    this._applications()
                )}
            </div>
        )
    }

    _applications() {
        if (this.state.loading) {
            return null
        }

        if (!this.state.clients.length) {
            return (<p className="blank-list">No applications available yet.</p>)
        }

        let applications = []
        for (var i = 0; i < this.state.clients.length; i++) {
            let client = this.state.clients[i]
            applications.push(
                <div className="application-card" key={i}>
                    <p className="title">
                        {client.name}
                        &nbsp;
                        <small>(<a href={client.uri.split('\n')[0]}
                            rel="noopener noreferrer"
                            target="_blank">{client.uri.split('\n')[0].split(/:\/\//)[1]}</a>)</small>
                    </p>
                    <p className="description">{client.description}</p>
                    <ul className="inline-list all-applications-options">
                        <li>
                            <a href="#revoke"
                                onClick={this._revokeAccess.bind(client)}>Revoke access</a>
                        </li>
                    </ul>
                </div>
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

export default MyApplications
