import React from 'react'

import './style.css'

const spinningSquare = ({ style }) => (
    <div style={style || {}} className="spinning-square__wrapper">
        <div className="spinning-square" />
    </div>
)

export default spinningSquare
