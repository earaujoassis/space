import React from 'react'

import './style.css'

const personal = () => {
    return (
        <>
            <h2>Personal information</h2>
            <div className="personal-root">
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="personal__full-name">Full name</label>
                        <input disabled id="personal__full-name" value="Carlos Assis" type="text" />
                    </div>
                    <div className="globals__input-wrapper">
                        <label htmlFor="personal__username">Username</label>
                        <input disabled id="personal__username" value="earaujoassis" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="personal__email">Email</label>
                        <input disabled id="personal__email" value="earaujoassis@example.com" type="text" />
                    </div>
                    <div className="globals__input-wrapper">
                        <label htmlFor="personal__role">Role</label>
                        <input disabled id="personal__role" value="Administrator" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="personal__timezone">Timezone</label>
                        <input disabled id="personal__timezone" value="UTC" type="text" />
                    </div>
                    <div className="globals__input-wrapper"></div>
                </div>
            </div>
        </>
    )
}

export default personal
