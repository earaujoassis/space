const extractDataForm = (form, attributes) => {
  const data = new FormData()

  for (const attr of attributes) {
    data.append(attr, form.querySelectorAll(`[name='${attr}']`)[0].value)
  }

  return data
}

const prependUrlWithHttps = e => {
  const string = e.target.value
  if (!~string.indexOf('http') && string.length) {
    e.target.value = 'https://' + string
  }
}

export { extractDataForm, prependUrlWithHttps }
