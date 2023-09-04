import React from 'react'

import RadioGroup from '@ui/RadioGroup'
import DynamicList from '@ui/DynamicList'

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
                    onChange={(v) => { v }}
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
                <DynamicList
                    label="Canonical URI"
                    labelPlural="Canonical URIs"
                    removeTitle="Remove canonical URI"
                    id="canonical-uri" />
                <DynamicList
                    label="Redirect URI"
                    labelPlural="Redirect URIs"
                    removeTitle="Remove redirect URI"
                    id="redirect-uri" />
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
