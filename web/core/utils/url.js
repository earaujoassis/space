const getParameterByName = (name, url) => {
  if (!url) {
    url = window.location.href
  }
  // eslint-disable-next-line no-useless-escape
  name = name.replace(/[\[\]]/g, '\\$&')
  var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)')
  var results = regex.exec(url)
  if (!results) return null
  if (!results[2]) return ''
  return results[2]
}

export { getParameterByName }
