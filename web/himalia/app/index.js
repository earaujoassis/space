import React, { Suspense } from 'react'
import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom'

import Layout from '@containers/Layout'
import Applications from '@components/Applications'
import { AllClients, EditClient, EditScopes, NewClient } from '@components/Clients'
import Notifications from '@components/Notifications'
import Profile from '@components/Profile'
import Security from '@components/Security'
import { AllServices, NewService } from '@components/Services'

import './style.css'

const app = () => {
    const router = createBrowserRouter([
        {
            path: '/',
            element: <Layout />,
            children: [
                {
                    index: true,
                    element: <Navigate to="/profile" replace />,
                },
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
                    path: 'clients/edit/scopes',
                    element: <EditScopes />,
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
                    element: <Profile />,
                },
                {
                    path: 'security',
                    element: <Security />,
                },
                {
                    path: 'services',
                    element: <AllServices />,
                },
                {
                    path: 'services/new',
                    element: <NewService />,
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
