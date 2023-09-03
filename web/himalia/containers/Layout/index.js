import React from 'react'
import { Outlet } from 'react-router-dom'

import Menu from '@components/Menu'

import './style.css'

const layout = () => {
    return (
        <div role="main" className="layout-root">
            <div className="layout-root__menu-container">
                <Menu />
            </div>
            <div className="layout-root__corpus-container">
                <Outlet />
            </div>
        </div>
    )
}

export default layout
