import React, { useState } from 'react'

import Row from '../../../core/components/Row.jsx'
import Columns from '../../../core/components/Columns.jsx'

import NewApplication from './NewApplication.jsx'
import AllApplications from './AllApplications.jsx'
import MyApplications from './MyApplications.jsx'

import UserStore from '../../stores/users'

const applications = () => {
    const [openAccordion, setOpenAccordion] = useState('none')

    return (
        <div role="main">
            <Row>
                <Columns className="small-12">
                    <div className="jupiter-accordion">
                        <div className={`jupiter-accordion-child ${openAccordion === 'my' ? 'open' : ''}`}>
                            <h2 className="jupiter-accordion-title">
                                <a href="#my" onClick={(e) => {
                                    e.preventDefault()
                                    if (openAccordion === 'my') {
                                        setOpenAccordion('none')
                                    } else {
                                        setOpenAccordion('my')
                                    }
                                }}>
                                    My applications
                                </a>
                            </h2>
                            <div className="jupiter-accordion-body">
                                <Row>
                                    <Columns className="small-offset-1 small-10 end">
                                        <Row className="applications">
                                            <MyApplications />
                                        </Row>
                                    </Columns>
                                </Row>
                            </div>
                        </div>
                        {UserStore.isCurrentUserAdmin() && (
                            <div className={`jupiter-accordion-child ${openAccordion === 'all' ? 'open' : ''}`}>
                                <h2 className="jupiter-accordion-title">
                                    <a href="#all" onClick={(e) => {
                                        e.preventDefault()
                                        if (openAccordion === 'all') {
                                            setOpenAccordion('none')
                                        } else {
                                            setOpenAccordion('all')
                                        }
                                    }}>
                                        All applications
                                    </a>
                                </h2>
                                <div className="jupiter-accordion-body">
                                    <Row>
                                        <Columns className="small-offset-1 small-10 end">
                                            <Row className="applications">
                                                <AllApplications />
                                            </Row>
                                        </Columns>
                                    </Row>
                                </div>
                            </div>
                        )}
                    </div>
                </Columns>
            </Row>
            {UserStore.isCurrentUserAdmin() && (
                <div className="applications-divisor">
                    <h2 className="applications-divisor-title">
                        <span>New application</span>
                    </h2>
                    <Row>
                        <Columns className="small-12">
                            <NewApplication postCreation={() => setOpenAccordion('all')} />
                        </Columns>
                    </Row>
                </div>
            )}
        </div>
    )
}

export default applications
