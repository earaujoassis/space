import React, { useState } from 'react'

import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

import { extractDataForm } from '../../core/utils/forms'

import UserStore from '../stores/users'
import UsersActions from '../actions/users'
import ClientStore from '../stores/clients'
import ClientsActions from '../actions/clients'

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
            <Row className="applications-listing">
                <Columns className="small-12">
                    {this.state.loading ? (
                        <p className="text-center">Loading...</p>
                    ) : (
                        this._applications()
                    )}
                </Columns>
            </Row>
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
                <Row key={i}>
                    <Columns className="small-12">
                        <div className="application-card">
                            <p className="title">
                                {client.name}
                                &nbsp;
                                <small>(<a href={client.uri} rel="noopener noreferrer" target="_blank">{client.uri.split(/:\/\//)[1]}</a>)</small>
                            </p>
                            <p className="action">
                                <button className="button"
                                    onClick={this._revokeAccess.bind(client)}>
                                    Revoke Access
                                </button>
                            </p>
                            <p className="description">{client.description}</p>
                        </div>
                    </Columns>
                </Row>
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

class AllApplications extends React.Component {
    constructor() {
        super()
        this.state = {loading: true}
        this._updateFromStore = this._updateFromStore.bind(this)
        this._applications = this._applications.bind(this)
    }

    componentDidMount() {
        ClientStore.addChangeListener(this._updateFromStore)
        ClientsActions.fetchClients()
    }

    componentWillUnmount() {
        ClientStore.removeChangeListener(this._updateFromStore)
    }

    render() {
        return (
            <Row className="applications-listing">
                <Columns className="small-12">
                    {this.state.loading ? (
                        <p className="text-center">Loading...</p>
                    ) : (
                        this._applications()
                    )}
                </Columns>
            </Row>
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
                <Row key={i}>
                    <Columns className="small-12">
                        <div className="application-card">
                            <p className="title">
                                {client.name}
                                &nbsp;
                                <small>(<a href={client.uri} rel="noopener noreferrer" target="_blank">{client.uri.split(/:\/\//)[1]}</a>)</small>
                            </p>
                            <p className="description">{client.description}</p>
                        </div>
                    </Columns>
                </Row>
            )
        }
        return applications
    }

    _updateFromStore() {
        if (ClientStore.success()) {
            let state = Object.assign({}, ClientStore.getState().payload || {}, {loading: false})
            this.setState(state)
        }
    }
}

const NewApplication = () => {
    return (
        <Row className="new-application">
            <Columns className="small-6 description">
                <h2 className="title">Create a new client application</h2>
                <p className="description">
                    By clicking &quot;Create Application&quot;, you agree to our <a href="//quatrolabs.com/terms-of-service">terms
                    of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. Also, you guarantee that the corresponding
                    client application will adhere to those terms and policites, while handling user data.
                </p>
            </Columns>
            <Columns className="small-6">
                <form
                    className="form-sign-up"
                    action="."
                    method="post"
                    onSubmit={(e) => {
                        e.preventDefault()
                        const attrs = [ 'name', 'description', 'canonical_uri', 'redirect_uri' ]
                        ClientsActions.createClient(extractDataForm(e.target, attrs))
                    }}>
                    <input type="hidden" name="action_token" value="" />
                    <input type="text" name="name" placeholder="Name" required />
                    <input type="text" name="description" placeholder="Description" required />
                    <input type="url" name="canonical_uri" placeholder="Canonical URI" pattern="https://.*" required />
                    <input type="url" name="redirect_uri" placeholder="Redirect URI" pattern="https://.*" required />
                    <button type="submit" className="button expand">Create Application</button>
                </form>
            </Columns>
        </Row>
    )
}

const Applications = () => {
    const [ openAccordion, setOpenAccordion ] = useState('my')

    return (
        <div className="jupiter-accordion" role="main">
            <div className={`jupiter-accordion-child ${openAccordion === 'my' ? 'open' : ''}`}>
                <h2 className="jupiter-accordion-title">
                    <a href="#my" onClick={(e) => {
                        e.preventDefault()
                        if (openAccordion === 'my') {
                            setOpenAccordion('none')
                        } else {
                            setOpenAccordion('my')
                        }
                    }}>
                        My applications
                    </a>
                </h2>
                <div className="jupiter-accordion-body">
                    <Row>
                        <Columns className="small-offset-1 small-10 end">
                            <Row className="applications">
                                <MyApplications />
                            </Row>
                        </Columns>
                    </Row>
                </div>
            </div>
            {UserStore.isCurrentUserAdmin() && (
                <div className={`jupiter-accordion-child ${openAccordion === 'all' ? 'open' : ''}`}>
                    <h2 className="jupiter-accordion-title">
                        <a href="#all" onClick={(e) => {
                            e.preventDefault()
                            if (openAccordion === 'all') {
                                setOpenAccordion('none')
                            } else {
                                setOpenAccordion('all')
                            }
                        }}>
                            All applications
                        </a>
                    </h2>
                    <div className="jupiter-accordion-body">
                        <Row>
                            <Columns className="small-offset-1 small-10 end">
                                <Row className="applications">
                                    <AllApplications />
                                    <NewApplication />
                                </Row>
                            </Columns>
                        </Row>
                    </div>
                </div>
            )}
        </div>
    )
}

export default Applications
