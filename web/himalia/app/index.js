import React, { Suspense } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import Layout from '@containers/Layout'
import Personal from '@components/Personal'

import './style.css'

const app = () => {
    const router = createBrowserRouter([
        {
            path: '/himalia',
            element: <Layout />,
            children: [
                {
                    path: 'profile',
                    element: <Personal />,
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
