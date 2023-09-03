import React, { Suspense } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import Layout from '@containers/Layout'
import Applications from '@components/Applications'
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
