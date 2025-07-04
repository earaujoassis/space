import { useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'

const useProtectedResource = (resourceKey, fetchFunction) => {
  const { data, loading, error, stale } = useSelector(state => {
    return state[resourceKey]
  })

  const dispatch = useDispatch()

  useEffect(() => {
    if ((!data || stale) && !loading) {
      dispatch(fetchFunction())
    }
  }, [data, stale, dispatch, fetchFunction])

  return { data, loading, error }
}

export default useProtectedResource
