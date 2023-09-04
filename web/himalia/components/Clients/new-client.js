import React from 'react'

import Submenu from './submenu'

const newClient = () => {
    const termsLink = (
        <a href="//quatrolabs.com/terms-of-service">terms of service</a>
    )
    const privacyPolicyLink = (
        <a href="//quatrolabs.com/privacy-policy">privacy policy</a>
    )

    return (
        <>
            <h2>Create a new client application</h2>
            <Submenu activeAction="new-client" />
            <p>
                By clicking &quot;Create client application&quot;, you agree to our {termsLink} and {privacyPolicyLink}.
                Also, you guarantee that the corresponding client application will adhere to those terms
                and policites, while handling user data.
            </p>
            <div className="clients-root">
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__name">Name</label>
                        <input id="new-client__name" value="" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__description">Description</label>
                        <input id="new-client__description" value="" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__canonical-uri">Canonical URI</label>
                        <input id="new-client__canonical-uri" value="" inputMode="url" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__redirect-uri">Redirect URI</label>
                        <input id="new-client__redirect-uri" value="" inputMode="url" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <input type="submit" className="button" value="Create client application" />
                    </div>
                </div>
            </div>
        </>
    )
}

export default newClient
