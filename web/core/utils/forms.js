const extractDataForm = (form, attributes) => {
  const data = new FormData();

  for (const attr of attributes) {
    data.append(attr, form.querySelectorAll(`[name='${attr}']`)[0].value);
  }

  return data;
};

export { extractDataForm };
