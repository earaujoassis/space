const extractDataForm = (form, attributes) => {
  let data = new FormData()

  for (let attr of attributes) {
    data.append(attr, form.querySelectorAll(`[name='${attr}']`)[0].value)
  }

  return data
}

export { extractDataForm }
