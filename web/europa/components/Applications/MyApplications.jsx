import React, { useEffect, useState } from 'react'

import UserStore from '../../stores/users'
import UsersActions from '../../actions/users'

const myApplications = () => {
    const [storeState, setStoreState] = useState({isLoading: true})

    useEffect(() => {
        const updateLocalStoreState = () => {
            if (UserStore.success()) {
                const state = Object.assign({}, UserStore.getState().payload || {}, {isLoading: false})
                setStoreState(state)
            }
        }

        UserStore.addChangeListener(updateLocalStoreState)
        UsersActions.fetchActiveClients()

        return function cleanup() {
            UserStore.removeChangeListener(updateLocalStoreState)
        }
    }, [])

    const { isLoading, clients, id } = storeState

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

    const applications = clients.map((client, i) =>
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
                        onClick={(e) => {
                            e.preventDefault()
                            UsersActions.revokeActiveClient(id)
                        }}>Revoke access</a>
                </li>
            </ul>
        </div>)


    return <div className="applications-listing">
        {applications}
    </div>

}

export default myApplications
