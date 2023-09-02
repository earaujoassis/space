import React, { Suspense } from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import Layout from '@containers/Layout'

import './style.css'

const app = () => {
    const router = createBrowserRouter([
        {
            path: '/himalia',
            element: <div style={{textAlign: 'center'}}>Hello world!</div>
        }
    ])

    return (
        <React.StrictMode>
            <Layout>
                <Suspense fallback={<p>Pending...</p>}>
                    <RouterProvider router={router} />
                </Suspense>
            </Layout>
        </React.StrictMode>
    )
}

export default app
