import React from 'react'

import Row from '../../../core/components/Row.jsx'
import Columns from '../../../core/components/Columns.jsx'

import { extractDataForm } from '../../../core/utils/forms'

import ClientStore from '../../stores/clients'
import ClientsActions from '../../actions/clients'

export default class AllApplications extends React.Component {
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
                        <small>(<a href={client.uri.split('\n')[0]}
                            rel="noopener noreferrer"
                            target="_blank">{client.uri.split('\n')[0].split(/:\/\//)[1]}</a>)</small>
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
                                    <textarea type="url"
                                        id="canonical_uri"
                                        name="canonical_uri"
                                        placeholder="Canonical URI"
                                        pattern="((https?://.*)(\n)?)+"
                                        defaultValue={client.uri}
                                        required></textarea>
                                </Columns>
                                <Columns className="small-5">
                                    <label htmlFor="redirect_uri">Redirect URI</label>
                                    <textarea type="url"
                                        id="redirect_uri"
                                        name="redirect_uri"
                                        placeholder="Redirect URI"
                                        pattern="((https?://.*)(\n)?)+"
                                        defaultValue={client.redirect}
                                        required></textarea>
                                </Columns>
                                <Columns className="small-2 end">
                                    <div className="button-wrapper">
                                        <button className="button-anchor" type="submit">Save</button>
                                    </div>
                                    <div>
                                        <button
                                            onClick={() => this.setState({editingId: null})}
                                            className="button-anchor"
                                            type="button">Cancel</button>
                                    </div>
                                </Columns>
                            </Row>
                            <Row>
                                <Columns className="small-12">
                                    <p className="form-warning">
                                        Use one line for each redirect URI. Each redirect URI must be included as a canonical URI.
                                    </p>
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