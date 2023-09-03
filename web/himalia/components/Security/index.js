import React from 'react'

import './style.css'

const security = () => {
    return (
        <>
            <h2>Password &amp; Security</h2>
            <div className="security-root">
                <div className="globals__warning-box">
                    <h3>Update password through a magic link</h3>
                    <p>
                        Update your password through a magic link sent to your e-mail account. It will generate
                        a temporary token so you can securily modify your account password.
                    </p>
                    <p><button className="button-anchor">Request link to update password</button></p>
                </div>
                <div className="globals__warning-box">
                    <h3>Recreate recovery code and secret code generator</h3>
                    <p>
                        Through this request, you will recreate your account recovery code and the secret code generator
                        used in your secondary factor authenticator app. This process is irreversible.
                    </p>
                    <p><button className="button-anchor">Proceed to recreate recovery code and secret code generator</button></p>
                </div>
            </div>
        </>
    )
}

export default security
