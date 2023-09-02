import React from 'react'

import './style.css'

const layout = ({ children }) => {
    return (
        <div role="main" className="layout-root">
            <div className="layout-root__corpus">
                <div className="layout-root__siblings">
                    <div className="layout-root__corpus">
                        {children}
                    </div>
                </div>
            </div>
        </div>
    )
}

export default layout
