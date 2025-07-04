import React, { useEffect } from 'react'
import { connect } from 'react-redux'

import * as actions from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import Submenu from './submenu'

const allServices = ({ fetchServices, loading, services }) => {
  let content = null

  useEffect(() => {
    fetchServices()
  }, [])

  if (loading.includes('service') || services === undefined) {
    content = <SpinningSquare />
  } else if (services && services.length) {
    content = (
      <ul className="services-root__list">
        {services.map((service, i) => {
          return (
            <li key={i}>
              <div className="services-root__entry">
                <a href={service.uri} rel="noreferrer" target="_blank">
                  {service.name}
                </a>
              </div>
            </li>
          )
        })}
      </ul>
    )
  }

  return (
    <>
      <h2>Services</h2>
      <Submenu activeAction="all-services" />
      <div className="services-root">{content}</div>
    </>
  )
}

const mapStateToProps = state => {
  return {
    loading: state.root.loading,
    services: state.root.services,
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchServices: () => dispatch(actions.fetchServices()),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(allServices)
