import React, { Suspense } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import Layout from '@containers/Layout'
import Applications from '@components/Applications'
import { AllClients, EditClient, NewClient } from '@components/Clients'
import Notifications from '@components/Notifications'
import Personal from '@components/Personal'
import Security from '@components/Security'

import './style.css'

const app = () => {
    const router = createBrowserRouter([
        {
            path: '/himalia',
            element: <Layout />,
            children: [
                {
                    path: 'applications',
                    element: <Applications />,
                },
                {
                    path: 'clients',
                    element: <AllClients />,
                },
                {
                    path: 'clients/edit',
                    element: <EditClient />,
                },
                {
                    path: 'clients/new',
                    element: <NewClient />,
                },
                {
                    path: 'notifications',
                    element: <Notifications />,
                },
                {
                    path: 'profile',
                    element: <Personal />,
                },
                {
                    path: 'security',
                    element: <Security />,
                },
            ],
        }
    ])

    return (
        <React.StrictMode>
            <Suspense fallback={<p>Pending...</p>}>
                <RouterProvider router={router} />
            </Suspense>
        </React.StrictMode>
    )
}

export default app
