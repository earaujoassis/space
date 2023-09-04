import React from 'react'

import RadioGroup from '@ui/RadioGroup'

import Submenu from './submenu'

const editClient = () => {
    return (
        <>
            <h2>Edit client application</h2>
            <Submenu activeAction="edit-client" editingClient />
            <div className="clients-root">
                <p>
                    By default, all applications are created with &quot;Public&quot; scope, which
                    can&apos;t instrospect user data (read user&apos;s full name, email etc.) If
                    your application needs to read user data, you must set the &quot;Application scope&quot;
                    to &quot;Read&quot;.
                </p>
                <RadioGroup
                    label="Application scope:"
                    leftOption="Public"
                    rightOption="Read" />
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__name">Name</label>
                        <input disabled id="new-client__name" value="" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <label htmlFor="new-client__description">Description</label>
                        <input disabled id="new-client__description" value="" type="text" />
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <h4 className="clients-root__list-header">Canonical URIs</h4>
                        <ul className="clients-root__list-urls">
                            <li>
                                <div className="clients-root__list-entry">
                                    <button
                                        className="clients-root__list-remove-entry"
                                        title="Remove canonical URI">&times;</button>
                                    <span>http://google.com</span>
                                </div>
                            </li>
                            <li>
                                <div className="clients-root__list-entry">
                                    <button
                                        className="clients-root__list-remove-entry"
                                        title="Remove canonical URI">&times;</button>
                                    <span>http://google.org</span>
                                </div>
                            </li>
                        </ul>
                    </div>
                    <div className="globals__input-wrapper clients-root__shared-wrapper">
                        <div className="clients-root__shared-wrapper__field">
                            <label htmlFor="new-client__canonical-uri">Canonical URI</label>
                            <input id="new-client__canonical-uri" value="" inputMode="url" type="text" />
                        </div>
                        <div className="clients-root__shared-wrapper__action">
                            <button className="button-anchor">Add entry</button>
                        </div>
                    </div>
                </div>
                <div className="globals__siblings">
                    <div className="globals__input-wrapper">
                        <h4 className="clients-root__list-header">Redirect URIs</h4>
                        <ul className="clients-root__list-urls">
                            <li>
                                <div className="clients-root__list-entry">
                                    <button
                                        className="clients-root__list-remove-entry"
                                        title="Remove redirect URI">&times;</button>
                                    <span>http://google.com/ouath/callback</span>
                                </div>
                            </li>
                            <li>
                                <div className="clients-root__list-entry">
                                    <button
                                        className="clients-root__list-remove-entry"
                                        title="Remove redirect URI">&times;</button>
                                    <span>http://google.org/ouath/callback</span>
                                </div>
                            </li>
                        </ul>
                    </div>
                    <div className="globals__input-wrapper clients-root__shared-wrapper">
                        <div className="clients-root__shared-wrapper__field">
                            <label htmlFor="new-client__redirect-uri">Redirect URI</label>
                            <input id="new-client__redirect-uri" value="" inputMode="url" type="text" />
                        </div>
                        <div className="clients-root__shared-wrapper__action">
                            <button className="button-anchor">Add entry</button>
                        </div>
                    </div>
                </div>
                <div className="globals__siblings globals__form-actions">
                    <div className="globals__input-wrapper">
                        <input type="submit" className="button" value="Save client application" />
                        <button className="submit cancel">Cancel</button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default editClient
