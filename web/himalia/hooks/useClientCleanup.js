import { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { useLocation } from 'react-router-dom'

import { staleClientRecords } from '@actions'

const useClientCleanup = () => {
  const dispatch = useDispatch()
  const location = useLocation()

  useEffect(() => {
    return () => {
      dispatch(staleClientRecords())
    }
  }, [location.pathname])
}

export default useClientCleanup
