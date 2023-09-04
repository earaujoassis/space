import React from 'react'

import './style.css'

const applications = () => {
    return (
        <>
            <h2>Applications</h2>
            <div className="applications-root">
                <p>The following applications are associated with your user account.</p>
                <ul className="applications-root__list">
                    <li>
                        <div className="applications-root__entry">
                            <h3>Google <span>(<a href="//google.com">google.com</a>)</span></h3>
                            <p>
                                Google LLC is an American multinational technology company focusing on
                                artificial intelligence, online advertising, search engine technology,
                                cloud computing, computer software, quantum computing, e-commerce, and
                                consumer electronics.
                            </p>
                            <p><button className="button-anchor">Revoke access</button></p>
                        </div>
                    </li>
                    <li>
                        <div className="applications-root__entry">
                            <h3>Facebook <span>(<a href="//facebook.com">facebook.com</a>)</span></h3>
                            <p>
                                Facebook is an online social media and social networking service owned by
                                American technology giant Meta Platforms.
                            </p>
                            <p><button className="button-anchor">Revoke access</button></p>
                        </div>
                    </li>
                </ul>
            </div>
        </>
    )
}

export default applications
