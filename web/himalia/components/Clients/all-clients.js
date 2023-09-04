import React from 'react'
import { Link } from 'react-router-dom'

import Submenu from './submenu'

const allClients = () => {
    return (
        <>
            <h2>Clients</h2>
            <Submenu activeAction="all-clients" />
            <div className="clients-root">
                <ul className="clients-root__list">
                    <li>
                        <div className="clients-root__entry">
                            <h3>Google <span>(<a href="//google.com">google.com</a>)</span></h3>
                            <p>
                                Google LLC is an American multinational technology company focusing on
                                artificial intelligence, online advertising, search engine technology,
                                cloud computing, computer software, quantum computing, e-commerce, and
                                consumer electronics.
                            </p>
                            <p>
                                <Link to="/himalia/clients/edit" className="button-anchor">Edit</Link>
                                <button className="button-anchor">Download credentials</button>
                            </p>
                        </div>
                    </li>
                    <li>
                        <div className="clients-root__entry">
                            <h3>Facebook <span>(<a href="//facebook.com">facebook.com</a>)</span></h3>
                            <p>
                                Facebook is an online social media and social networking service owned by
                                American technology giant Meta Platforms.
                            </p>
                            <p>
                                <button className="button-anchor">Edit</button>
                                <button className="button-anchor">Download credentials</button>
                            </p>
                        </div>
                    </li>
                </ul>
            </div>
        </>
    )
}

export default allClients
