import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import Submenu from './submenu'

const allClients = ({ fetchClients, setClientForEdition, loading, application, clients }) => {
    const navigate = useNavigate()
    let content = null

    useEffect(() => {
        fetchClients(application.action_token)
    }, [])

    if (loading.includes('client') || clients === undefined) {
        content = (<SpinningSquare />)
    } else if (clients && clients.length) {
        content = (
            <ul className="clients-root__list">
                {clients.map((client, i) => {
                    const canonicalUri = client.uri.split('\n')[0]
                    const canonicalUriShort = canonicalUri.split(/:\/\//)[1]
                    return (
                        <li key={i}>
                            <div className="clients-root__entry">
                                <h3>{client.name} <span>(<a href={canonicalUri}>{canonicalUriShort}</a>)</span></h3>
                                <p>{client.description}</p>
                                <p>
                                    <button
                                        onClick={() => {
                                            setClientForEdition(client)
                                            navigate('/clients/edit')
                                        }}
                                        className="button-anchor">Edit</button>
                                    <a
                                        href={`/api/clients/${client.id}/credentials`}
                                        title="It regenerates the client's secret for security reasons"
                                        rel="noopener noreferrer"
                                        className="button-anchor">Download credentials</a>
                                </p>
                            </div>
                        </li>
                    )
                })}
            </ul>
        )
    }

    return (
        <>
            <h2>Clients</h2>
            <Submenu activeAction="all-clients" />
            <div className="clients-root">
                {content}
            </div>
        </>
    )
}

const mapStateToProps = state => {
    return {
        loading: state.root.loading,
        application: state.root.application,
        clients: state.root.clients
    }
}

const mapDispatchToProps = dispatch => {
    return {
        fetchClients: (token) => dispatch(actions.fetchClients(token)),
        setClientForEdition: (client) => dispatch(actions.setClientForEdition(client))
    }
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(allClients)
