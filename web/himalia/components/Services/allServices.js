import React, { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import { fetchServices } from '@actions'
import SpinningSquare from '@ui/SpinningSquare'

import Submenu from './submenu'

const allServices = () => {
  const loading = useSelector(state => state.root.loading)
  const services = useSelector(state => state.root.services)

  const dispatch = useDispatch()

  let content = null

  useEffect(() => {
    dispatch(fetchServices())
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

export default allServices
