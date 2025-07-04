import { configureStore } from '@reduxjs/toolkit'
import thunk from 'redux-thunk'

import reducers from './reducers'

const store = configureStore({
  reducer: {
    root: reducers,
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(thunk),
  devTools: process.env.NODE_ENV !== 'production',
})

export default store
