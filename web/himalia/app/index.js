import React from 'react'
import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom'

import Layout from '@containers/Layout'
import Applications from '@components/Applications'
import { Clients, EditClient, EditScopes, NewClient } from '@components/Clients'
import Notifications from '@components/Notifications'
import Profile from '@components/Profile'
import Emails from '@components/Emails'
import Security from '@components/Security'
import { Services, NewService } from '@components/Services'

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
          path: 'emails',
          element: <Emails />,
        },
        {
          path: 'applications',
          element: <Applications />,
        },
        {
          path: 'clients',
          element: <Clients />,
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
          element: <Services />,
        },
        {
          path: 'services/new',
          element: <NewService />,
        },
      ],
    },
  ])

  return (
    <React.StrictMode>
      <React.Suspense fallback={<p>Pending...</p>}>
        <RouterProvider router={router} />
      </React.Suspense>
    </React.StrictMode>
  )
}

export default app
