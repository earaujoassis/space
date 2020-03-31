import React, { useState } from 'react'

import SignIn from './SignIn.jsx'
import MagicLink from './MagicLink.jsx'

const PASSWORD_STRATEGY = 'password'
const MAGIC_LINK_STRATEGY = 'magic-link'

const root = () => {
    const [currentStrategy, setSignInStrategy] = useState(PASSWORD_STRATEGY)

    return (
        <div className="ganymede-root">
            {currentStrategy === PASSWORD_STRATEGY ? (
                <SignIn />
            ) : (
                <MagicLink />
            )}
            <div className="ganymede-strategies">
                {currentStrategy === PASSWORD_STRATEGY ? (
                    <button
                        className="button-anchor"
                        onClick={() => setSignInStrategy(MAGIC_LINK_STRATEGY)}>Request magic link</button>
                ) : (
                    <button
                        className="button-anchor"
                        onClick={() => setSignInStrategy(PASSWORD_STRATEGY)}>Use password and TOTP</button>
                )}
            </div>
        </div>
    )
}

export default root
