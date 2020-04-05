import React, { useEffect, useState } from 'react'

import { extractDataForm } from '../../../core/utils/forms'

import ClientStore from '../../stores/clients'
import ClientsActions from '../../actions/clients'

import EditApplication from './EditApplication.jsx'

const allApplications = () => {
    const [storeState, setStoreState] = useState({isLoading: true})
    const [editingId, setEditingId] = useState(null)

    useEffect(() => {
        const updateLocalStoreState = () => {
            if (ClientStore.success()) {
                const state = Object.assign({}, ClientStore.getState().payload || {}, {isLoading: false})
                setStoreState(state)
            }
        }

        ClientStore.addChangeListener(updateLocalStoreState)
        ClientsActions.fetchClients()

        return function cleanup() {
            ClientStore.removeChangeListener(updateLocalStoreState)
        }
    }, [])

    const { isLoading, clients } = storeState

    if (isLoading) {
        return <div className="applications-listing">
            <p className="text-center">Loading...</p>
        </div>
    }

    if (!clients || !clients.length) {
        return <div className="applications-listing">
            <p className="blank-list">No applications available yet.</p>
        </div>
    }

    const applications = clients.map((client, i) => <div className="application-card" key={i}>
        <p className="title">
            {client.name}
            &nbsp;
            <small>(<a href={client.uri.split('\n')[0]}
                rel="noopener noreferrer"
                target="_blank">{client.uri.split('\n')[0].split(/:\/\//)[1]}</a>)</small>
        </p>
        <p className="description">{client.description}</p>
        {editingId === client.id && (
            <EditApplication client={client}
                onCancel={() => setEditingId(null)}
                onSubmit={(target) => {
                    const attrs = ['canonical_uri', 'redirect_uri', 'scopes']
                    ClientsActions.updateClient(client.id, extractDataForm(target, attrs)).then(() => {
                        ClientsActions.fetchClients()
                        setEditingId(null)
                    })
                }} />
        )}
        <ul className="inline-list all-applications-options">
            <li>
                <a href="#edit"
                    onClick={(e) => {
                        e.preventDefault()
                        setEditingId(client.id)
                    }}>Edit</a>
            </li>
            <li>
                <a href={`/api/clients/${client.id}/credentials`}
                    title="It regenerates the client's secret for security reasons"
                    rel="noopener noreferrer"
                    target="_blank">Download credentials</a>
            </li>
        </ul>
    </div>)

    return (
        <div className="applications-listing">
            {applications}
        </div>
    )
}

export default allApplications
