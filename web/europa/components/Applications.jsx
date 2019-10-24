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
                        <small>(<a href={client.uri} rel="noopener noreferrer" target="_blank">{client.uri.split(/:\/\//)[1]}</a>)</small>
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

        if (!this.state.clients || !this.state.clients.length) {
            return (<p className="blank-list">No applications available yet.</p>)
        }

        let applications = this.state.clients.map((client, i) => {
            return (
                <div className="application-card" key={i}>
                    <p className="title">
                        {client.name}
                        &nbsp;
                        <small>(<a href={client.uri} rel="noopener noreferrer" target="_blank">{client.uri.split(/:\/\//)[1]}</a>)</small>
                    </p>
                    <p className="description">{client.description}</p>
                    {this.state.editingId === client.id && (
                        <form
                            className="form-common internal"
                            action="."
                            method="post"
                            onSubmit={(e) => {
                                e.preventDefault()
                                const attrs = [ 'canonical_uri', 'redirect_uri' ]
                                ClientsActions.updateClient(client.id, extractDataForm(e.target, attrs)).then(() => {
                                    ClientsActions.fetchClients()
                                    this.setState({editingId: null})
                                })
                            }}>
                            <Row className="new-application">
                                <Columns className="small-5">
                                    <label htmlFor="canonical_uri">Canonical URI</label>
                                    <input type="url"
                                        id="canonical_uri"
                                        name="canonical_uri"
                                        placeholder="Canonical URI"
                                        pattern="https?://.*"
                                        defaultValue={client.uri}
                                        required />
                                </Columns>
                                <Columns className="small-5">
                                    <label htmlFor="redirect_uri">Redirect URI</label>
                                    <input type="url"
                                        id="redirect_uri"
                                        name="redirect_uri"
                                        placeholder="Redirect URI"
                                        pattern="https?://.*"
                                        defaultValue={client.redirect}
                                        required />
                                </Columns>
                                <Columns className="small-2 end">
                                    <button className="button-anchor" type="submit">Save</button>
                                </Columns>
                            </Row>
                        </form>
                    )}
                    <ul className="inline-list all-applications-options">
                        <li>
                            <a href="#edit"
                                onClick={(e) => {
                                    e.preventDefault()
                                    this.setState({editingId: client.id})
                                }}>Edit</a>
                        </li>
                        <li>
                            <a href={`/api/clients/${client.id}/credentials`}
                                title="It regenerates the client's secret for security reasons"
                                rel="noopener noreferrer"
                                target="_blank">Download credentials</a>
                        </li>
                    </ul>
                </div>
            )
        })

        return applications
    }

    _updateFromStore() {
        if (ClientStore.success()) {
            let state = Object.assign({}, ClientStore.getState().payload || {}, {loading: false})
            this.setState(state)
        }
    }
}

const NewApplication = ({ postCreation }) => {
    return (
        <Row className="new-application">
            <Columns className="small-6">
                <h2 className="title">Create a new client application</h2>
                <p className="description">
                    By clicking &quot;Create Application&quot;, you agree to our <a href="//quatrolabs.com/terms-of-service">terms
                    of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. Also, you guarantee that the corresponding
                    client application will adhere to those terms and policites, while handling user data.
                </p>
            </Columns>
            <Columns className="small-6">
                <form
                    className="form-common"
                    action="."
                    method="post"
                    onSubmit={(e) => {
                        e.persist()
                        e.preventDefault()
                        const attrs = [ 'name', 'description', 'canonical_uri', 'redirect_uri' ]
                        ClientsActions.createClient(extractDataForm(e.target, attrs)).then(() => {
                            ClientsActions.fetchClients()
                            if (postCreation) {
                                postCreation()
                            }
                            e.target.reset()
                        })
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
        <div role="main">
            {UserStore.isCurrentUserAdmin() && (
                <Row>
                    <Columns className="small-12">
                        <NewApplication postCreation={() => setOpenAccordion('all')} />
                    </Columns>
                </Row>
            )}
            <Row>
                <Columns className="small-12">
                    <div className="jupiter-accordion applications-divisor">
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
                                            </Row>
                                        </Columns>
                                    </Row>
                                </div>
                            </div>
                        )}
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
                    </div>
                </Columns>
            </Row>
        </div>
    )
}

export default Applications
