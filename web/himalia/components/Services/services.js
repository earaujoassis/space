import React from 'react'

import { fetchServices } from '@actions'
import { useProtectedResource } from '@hooks'

import SpinningSquare from '@ui/SpinningSquare'

import Submenu from './submenu'

const services = () => {
  const { data: services, loading } = useProtectedResource(
    'services',
    fetchServices
  )

  let content = null

  if (loading || services === undefined) {
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

export default services
